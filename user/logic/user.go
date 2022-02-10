package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/user/db"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
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
func (u *User) LoginCheck(username, password string) (bool, error) {
	md5Password := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	md5Password = strings.ToUpper(md5Password)
	return u.userDB.LoginQuery(username, md5Password)
}

// CreateToken 为指定用户名创建token
func (u *User) CreateToken(username string) string {
	u.DeleteToken(username)
	token := uuid.New().String()
	token = strings.Replace(token, "-", "", -1)
	u.mutex.Lock()
	u.redisClient.SetEX(context.Background(), TokenPrefix+token, username, u.TokenExpireTime)
	u.redisClient.SetEX(context.Background(), LoginUserTokenPrefix+username, token, u.TokenExpireTime)
	u.mutex.Unlock()
	return token
}

// DeleteToken 删除指定用户名的登录token
func (u *User) DeleteToken(username string) {
	token := u.GetToken(username)
	u.redisClient.Do(context.Background(), "del", TokenPrefix+token)
	u.redisClient.Do(context.Background(), "del", LoginUserTokenPrefix+username)
}

// GetToken 查询用户的token
func (u *User) GetToken(username string) string {
	token, err := u.redisClient.Get(context.Background(), LoginUserTokenPrefix+username).Result()
	if err != nil {
		return ""
	}
	return token
}

// GetUserByToken 通过token查询对应者的信息
func (u *User) GetUserByToken(token string) (*db.User, error) {
	username, err := u.redisClient.Get(context.Background(), TokenPrefix+token).Result()
	if err != nil {
		return nil, err
	}
	info, err := u.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetByUsername 通过用户名查询用户信息
func (u *User) GetByUsername(username string) (*db.User, error) {
	return u.userDB.QueryByUsername(username)
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
