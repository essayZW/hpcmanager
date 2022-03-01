package source

import (
	"context"
	"encoding/json"
	"errors"
	"os/exec"
	"path"
	"time"
)

// defaultSource 默认的作业调度系统源,主要通过执行脚本命令实现
type defaultSource struct {
	baseDir string
}

func (source *defaultSource) AddUserToGroup(userName, groupName string, gid int) (map[string]interface{}, error) {
	return nil, nil
}

func (source *defaultSource) AddUserWithGroup(userName, groupName string) (map[string]interface{}, error) {
	return source.timeoutExec("php", "useradd_withgroup.php", "--user", userName, "--group", groupName)
}

// exec 执行指定的命令
func (source *defaultSource) exec(ctx context.Context, executor, file string, args ...string) (map[string]interface{}, error) {
	cmdArgs := make([]string, len(args)+1)
	cmdArgs = append(cmdArgs, path.Join(source.baseDir))
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.CommandContext(ctx, executor, cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	if err := json.Unmarshal(output, &res); err != nil {
		return nil, errors.New(string(output))
	}
	return res, nil
}

func (source *defaultSource) timeoutExec(executor, file string, args ...string) (map[string]interface{}, error) {
	// 最大的超时时间为4秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(4))
	defer cancel()
	return source.exec(ctx, executor, file, args...)
}

func newSource(options ...Option) HpcSource {
	opts := Options{}
	for _, option := range options {
		option(&opts)
	}
	source := defaultSource{}
	source.baseDir = opts.CmdBaseDir
	return &source
}
