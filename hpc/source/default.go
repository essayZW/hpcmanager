package source

import (
	"context"
	"encoding/json"
	"errors"
	"os/exec"
	"path"
	"strconv"
	"time"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

// defaultSource 默认的作业调度系统源,主要通过执行脚本命令实现
type defaultSource struct {
	baseDir string
	conn    *db.DB
}

func (source *defaultSource) AddUserToGroup(
	userName, groupName string,
	gid int,
) (map[string]interface{}, error) {
	return source.timeoutExec(
		"php",
		"useradd.php",
		"--user",
		userName,
		"--group",
		groupName,
		"--gid",
		strconv.Itoa(gid),
	)
}

func (source *defaultSource) AddUserWithGroup(
	userName, groupName string,
) (map[string]interface{}, error) {
	return source.timeoutExec(
		"php",
		"useradd_withgroup.php",
		"--user",
		userName,
		"--group",
		groupName,
	)
}

func (source *defaultSource) GetNodeUsageWithDate(
	ctx context.Context,
	startTime, endTime time.Time,
) ([]*HpcNodeUsage, error) {
	return source.selectWithDate(ctx, startTime, endTime)
}

func (source *defaultSource) selectWithDate(
	ctx context.Context,
	startDate, endDate time.Time,
) ([]*HpcNodeUsage, error) {
	// 从jobDW数据库中获取原始的作业调度信息并计算之后返回
	rows, err := source.conn.Query(
		ctx,
		"SELECT `UserName`,`GroupName`,`Queue`, SUM(`WallDurationSeconds`) AS `WallTime`, SUM(`GpusWallTime`) as `GWallTime` "+
			" FROM `account` WHERE `EventDate` >= ? AND `EventDate`<=? GROUP BY `GroupName`, `UserName`, `Queue`",
		startDate,
		endDate,
	)
	if err != nil {
		logger.Warn("selectWithDate error: ", err)
		return nil, errors.New("selectWithDate error")
	}

	infos := make([]*HpcNodeUsage, 0)
	for rows.Next() {
		var info HpcNodeUsage
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("selectWithDate error: ", err)
			return nil, errors.New("selectWithDate error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// exec 执行指定的命令
func (source *defaultSource) exec(
	ctx context.Context,
	executor, file string,
	args ...string,
) (map[string]interface{}, error) {
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

func (source *defaultSource) timeoutExec(
	executor, file string,
	args ...string,
) (map[string]interface{}, error) {
	// 最大的超时时间为4秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(4))
	defer cancel()
	return source.exec(ctx, executor, file, args...)
}

func newSource(options *Options) (HpcSource, error) {
	source := defaultSource{}
	source.baseDir = options.CmdBaseDir
	var err error
	if options.dbConf == nil {
		return nil, errors.New("need db conf")
	}
	source.conn, err = db.NewDBWithConfig(options.dbConf)
	if err != nil {
		return nil, err
	}
	return &source, nil
}
