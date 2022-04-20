package logic

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/essayZW/hpcmanager/logger"
	"github.com/google/uuid"
)

var fileTypeWhite []string

func init() {
	fileTypeWhite = make([]string, 0)
	fileTypeWhite = append(fileTypeWhite, "jpg")
	fileTypeWhite = append(fileTypeWhite, "png")
	fileTypeWhite = append(fileTypeWhite, "jpeg")
}

type Fss struct {
	fileStorePath string
}

// StoreFile 存储文件
func (this *Fss) StoreFile(ctx context.Context, fileName string, file []byte) (string, error) {
	fileTypeName, err := this.validFileType(fileName)
	if err != nil {
		return "", fmt.Errorf("invalid file type, only support type: %v", fileTypeWhite)
	}
	newFileName := fmt.Sprintf("%s.%s", this.randomFileName(), fileTypeName)
	var storePath string
	done := make(chan struct{}, 1)
	go func() {
		storePath, err = this.storeFile(newFileName, file)
		if err == nil {
			logger.Info("store new file: ", storePath)
		} else {
			logger.Warn("store file error: ", err)
			err = errors.New("store file error")
		}
		done <- struct{}{}
	}()

	select {
	case <-done:
		return newFileName, err
	case <-ctx.Done():
		logger.Warn("store file error because context canceled")
		return "", errors.New("context canceled")
	}
}

func (this *Fss) storeFile(filePath string, fileData []byte) (string, error) {
	fileSavePath := path.Join(this.fileStorePath, filePath)
	err := os.WriteFile(fileSavePath, fileData, os.FileMode(0666))
	return fileSavePath, err
}

// validFileType 验证文件类型
func (this *Fss) validFileType(fileName string) (string, error) {
	if fileName == "" {
		return "", errors.New("empty filename")
	}
	names := strings.Split(fileName, ".")
	if len(names) == 1 {
		return "", errors.New("no file type")
	}
	typeName := names[len(names)-1]
	for _, name := range fileTypeWhite {
		if name == typeName {
			return name, nil
		}
	}
	return typeName, errors.New("invalid file type")
}

func (this *Fss) randomFileName() string {
	v4 := uuid.New()
	return fmt.Sprintf("%x", md5.Sum([]byte(v4.String())))
}

func NewFss(fileStorePath string) (*Fss, error) {
	// 检查fileStorePath是否存在
	if _, err := os.Stat(fileStorePath); os.IsNotExist(err) {
		logger.Info("fileStorePath not exists, create it")
		err = os.MkdirAll(fileStorePath, os.FileMode(0755))
		if err != nil {
			return nil, err
		}
	}
	return &Fss{
		fileStorePath: fileStorePath,
	}, nil
}
