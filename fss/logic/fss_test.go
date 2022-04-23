package logic

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager/logger"
)

func TestFss_StoreFile(t *testing.T) {
	type fields struct {
		fileStorePath string
	}
	type args struct {
		ctx      context.Context
		fileName string
		file     []byte
	}
	fileTypeWhite = append(fileTypeWhite, "txt")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test 1",
			fields: fields{
				fileStorePath: "/tmp/hpc",
			},
			args: args{
				ctx: func() context.Context {
					ctx, _ := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
					return ctx
				}(),
				fileName: "a.txt",
				file:     []byte("test 1"),
			},
			wantErr: false,
		},
		{
			name: "test 2",
			fields: fields{
				fileStorePath: "/tmp/hpc",
			},
			args: args{
				ctx: func() context.Context {
					ctx, _ := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
					return ctx
				}(),
				fileName: "a.php",
				file:     []byte("test 1"),
			},
			wantErr: true,
		},
		{
			name: "test 3",
			fields: fields{
				fileStorePath: "/tmp/hpc",
			},
			args: args{
				ctx: func() context.Context {
					ctx, _ := context.WithTimeout(context.Background(), time.Duration(10)*time.Millisecond)
					return ctx
				}(),
				fileName: "c.txt",
				file: func() []byte {
					res := make([]byte, 0)
					buffer := bytes.NewBuffer(res)
					builder := strings.Builder{}
					for i := 0; i <= 1024; i++ {
						builder.WriteString("x")
					}
					for i := 0; i <= 204800; i++ {
						buffer.WriteString(builder.String())
					}
					return buffer.Bytes()
				}(),
			},
			wantErr: true,
		},
		{
			name: "test 5",
			fields: fields{
				fileStorePath: "/tmp/hpc/xss",
			},
			args: args{
				ctx: func() context.Context {
					ctx, _ := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
					return ctx
				}(),
				fileName: "a.txt",
				file:     []byte("test 1"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this, err := NewFss(tt.fields.fileStorePath)
			if err != nil {
				t.Errorf("error %v", err)
				return
			}
			got, err := this.StoreFile(tt.args.ctx, tt.args.fileName, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fss.StoreFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			logger.Debug(got)
		})
	}
}
