package logic

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/user/db"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/mozillazg/go-pinyin"
)

// 对于已经登录的用户来说，会在redis中存储两个值
// 其中一个是TokenPrefix+token,用来快速判断某个token的归属者,其值为用户帐号
// 另外一个是LoginUserTokenPrefix+username,用来判断某个用户是否已经登录以及其登录的token是多少,其值是对应的token值
const (
	// TokenPrefix 用户token前缀
	TokenPrefix string = "__HPCMANAGER_USER_TOKEN__"
	// LoginUserTokenPrefix 已经登录的用户存储的其token值的前缀
	LoginUserTokenPrefix string = "__HPCMANAGER_LOGINED_USER__"
)

// User 用户logic类，主要处理用户相关的逻辑
type User struct {
	userDB      *db.UserDB
	redisClient *redis.Client

	mutex sync.Mutex
	// TokenExpireTime 用户token过期时间
	TokenExpireTime time.Duration
}

// LoginCheck 检查用户名密码是否正确并返回用户ID
func (u *User) LoginCheck(ctx context.Context, username, password string) (bool, error) {
	md5Password := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	md5Password = strings.ToUpper(md5Password)
	return u.userDB.LoginQuery(ctx, username, md5Password)
}

// CreateToken 为指定用户名创建token
func (u *User) CreateToken(ctx context.Context, username string) string {
	// 单点登录
	u.DeleteToken(ctx, username)
	token := uuid.New().String()
	token = strings.Replace(token, "-", "", -1)
	u.mutex.Lock()
	u.redisClient.SetEX(ctx, TokenPrefix+token, username, u.TokenExpireTime)
	u.redisClient.SetEX(ctx, LoginUserTokenPrefix+username, token, u.TokenExpireTime)
	u.mutex.Unlock()
	return token
}

// DeleteToken 删除指定用户名的登录token
func (u *User) DeleteToken(ctx context.Context, username string) {
	token := u.GetToken(ctx, username)
	u.redisClient.Do(ctx, "del", TokenPrefix+token)
	u.redisClient.Do(ctx, "del", LoginUserTokenPrefix+username)
}

// GetToken 查询用户的token
func (u *User) GetToken(ctx context.Context, username string) string {
	token, err := u.redisClient.Get(ctx, LoginUserTokenPrefix+username).Result()
	if err != nil {
		return ""
	}
	return token
}

// GetUserByToken 通过token查询对应者的信息
func (u *User) GetUserByToken(ctx context.Context, token string) (*db.User, error) {
	username, err := u.redisClient.Get(ctx, TokenPrefix+token).Result()
	if err != nil {
		return nil, err
	}
	info, err := u.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetByUsername 通过用户名查询用户信息
func (u *User) GetByUsername(ctx context.Context, username string) (*db.User, error) {
	return u.userDB.QueryByUsername(ctx, username)
}

// AddUser 添加新的用户
func (u *User) AddUser(ctx context.Context, userInfo *db.User) (int, error) {
	if userInfo.Username == "" {
		return 0, errors.New("username can't be empty")
	}
	if userInfo.Name == "" {
		return 0, errors.New("name can't be empty")
	}
	if userInfo.Password == "" {
		return 0, errors.New("password can't be empty")
	}
	if userInfo.PinyinName == "" {
		pinyinDict := pinyin.LazyPinyin(userInfo.Name, pinyin.NewArgs())
		userInfo.PinyinName = strings.Join(pinyinDict, "")
	}
	if userInfo.CreateTime.IsZero() {
		userInfo.CreateTime = time.Now()
	}
	// password进行MD5加密
	md5Password := fmt.Sprintf("%x", md5.Sum([]byte(userInfo.Password)))
	userInfo.Password = strings.ToUpper(md5Password)
	id, err := u.userDB.InsertUser(ctx, userInfo)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetUserInfoByID 通过ID查询用户信息
func (u *User) GetUserInfoByID(ctx context.Context, userid int) (*db.User, error) {
	return u.userDB.QueryByID(ctx, userid)
}

// PaginationUserResult 分页查询用户信息的结果
type PaginationUserResult struct {
	Infos []*db.User
	Count int
}

// PaginationGetUserInfo 分页查询用户信息
func (u *User) PaginationGetUserInfo(ctx context.Context, pageIndex, pageSize, groupID int) (*PaginationUserResult, error) {
	if pageIndex < 1 {
		return nil, errors.New("pageIndex must large than 0")
	}
	if pageSize <= 0 || pageSize > 200 {
		return nil, errors.New("pageSize must large than 0 and less than 200")
	}
	// 先查询总数
	count, err := u.userDB.QueryAllUserCount(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &PaginationUserResult{
			Infos: make([]*db.User, 0),
			Count: 0,
		}, nil
	}
	offset := pageSize * (pageIndex - 1)
	infos, err := u.userDB.PaginationQuery(ctx, offset, pageSize, groupID)
	if err != nil {
		return nil, err
	}
	return &PaginationUserResult{
		Infos: infos,
		Count: count,
	}, nil

}

// ChangeUserGroup 修改用户的组信息
func (u *User) ChangeUserGroup(ctx context.Context, userID int, groupID int) error {
	return u.userDB.UpdateUserGroup(ctx, userID, groupID)
}

// NewUser 创建一个新的userLogic
func NewUser(db *db.UserDB, configConn config.DynamicConfig, redisConn *redis.Client) *User {
	user := &User{
		userDB:          db,
		redisClient:     redisConn,
		TokenExpireTime: time.Duration(24) * time.Hour,
	}
	var expireTime float64
	configConn.Registry("user/TokenExpireTime", &expireTime, func(newV interface{}) {
		user.mutex.Lock()
		defer user.mutex.Unlock()
		user.TokenExpireTime = time.Duration(int(expireTime)) * time.Minute
	})
	return user
}
