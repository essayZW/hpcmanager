package source

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/essayZW/hpcmanager/logger"
	"github.com/go-redis/redis/v8"
)

const hashPrefix = "__HPC_SOURCE_DEV"
const redisGroupName = hashPrefix + "_GROUP_NAME__"
const redisUserName = hashPrefix + "_USER_NAME__"
const redisUserGroupName = hashPrefix + "_USER_GROUP_NAME__"

type hpcDev struct {
	redisConn *redis.Client
}

func (dev *hpcDev) createGroup(ctx context.Context, groupName string) (int64, error) {
	// 先创建用户组
	res := dev.redisConn.HLen(ctx, redisGroupName)
	count, err := res.Result()
	if err != nil {
		return 0, err
	}
	logger.Debug(count)
	count++
	dev.redisConn.HSet(ctx, redisGroupName, groupName, count+1000)
	return 1000 + count, nil
}

func (dev *hpcDev) AddUserWithGroup(
	userName string,
	groupName string,
) (map[string]interface{}, error) {
	// 创建用户
	gid, err := dev.createGroup(context.Background(), groupName)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	uid, err := dev.createUser(context.Background(), userName)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	dev.joinGroup(context.Background(), userName, groupName)
	return map[string]interface{}{
		"success": "true",
		"data": map[string]interface{}{
			"gname": groupName,
			"gid":   int(gid),
			"uname": userName,
			"uid":   int(uid),
		},
	}, nil
}

func (dev *hpcDev) createUser(ctx context.Context, username string) (int64, error) {
	res := dev.redisConn.HLen(ctx, redisUserName)
	count, err := res.Result()
	if err != nil {
		return 0, err
	}
	logger.Debug(count)
	count++
	dev.redisConn.HSet(ctx, redisUserName, username, count+1000)
	return count + 1000, nil
}

func (dev *hpcDev) joinGroup(ctx context.Context, userName, groupName string) error {
	dev.redisConn.HSet(ctx, redisUserGroupName, userName, groupName)
	return nil
}

func (dev *hpcDev) AddUserToGroup(
	userName string,
	groupName string,
	gid int,
) (map[string]interface{}, error) {
	uid, err := dev.createUser(context.Background(), userName)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	dev.joinGroup(context.Background(), userName, groupName)
	return map[string]interface{}{
		"success": "true",
		"data": map[string]interface{}{
			"uname": userName,
			"uid":   int(uid),
		},
	}, nil
}

func (dev *hpcDev) GetNodeUsageWithDate(
	ctx context.Context,
	startTime, endTime time.Time,
) ([]*HpcNodeUsage, error) {
	rander := rand.New(rand.NewSource(time.Now().UnixMicro()))
	infos := make([]*HpcNodeUsage, rander.Intn(64))
	logger.Debug("random info len: ", len(infos))
	usernames, err := dev.getAllUsernames(context.Background())
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	if len(usernames) == 0 {
		return nil, errors.New("empty user lists")
	}
	for i := range infos {
		randUserNameIndex := rander.Intn(len(usernames) - 1)
		groupName, err := dev.getUserGroup(context.Background(), usernames[randUserNameIndex])
		if err != nil {
			logger.Warn(err)
			return nil, err
		}
		wallTime := rander.Float64() * float64(rander.Intn(int(endTime.Sub(startTime).Seconds())))
		gwallTime := rander.Float64() * float64(rander.Intn(64))
		infos[i] = &HpcNodeUsage{
			Username:  usernames[randUserNameIndex],
			GroupName: groupName,
			QueueName: "g_" + groupName,
			WallTime:  wallTime,
			GWallTime: gwallTime,
		}
		logger.Debug(infos[i])
	}
	return infos, nil
}

func (dev *hpcDev) getUserGroup(ctx context.Context, username string) (string, error) {
	res := dev.redisConn.HGet(ctx, redisUserGroupName, username)
	group, err := res.Result()
	if err != nil {
		return "", err
	}
	return group, nil
}

func (dev *hpcDev) getAllUsernames(ctx context.Context) ([]string, error) {
	res := dev.redisConn.HKeys(ctx, redisUserName)
	strSlice, err := res.Result()
	if err != nil {
		return nil, err
	}
	return strSlice, nil
}

func newDev(options *Options) HpcSource {
	if options.redisConn == nil {
		logger.Fatal("invalid redis conn")
	}
	return &hpcDev{
		redisConn: options.redisConn,
	}
}
