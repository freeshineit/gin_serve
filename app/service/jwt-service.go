package service

import (
	"context"
	"gin_serve/config"
	"gin_serve/helper"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type JWTService interface {
	IsInBlacklist(tokenStr string) bool
	JoinBlackList(tokenStr string) error
}

type jwtService struct {
	redisClient *redis.Client
}

func NewJWTService(redisClient *redis.Client) JWTService {
	return &jwtService{
		redisClient: redisClient,
	}
}

// 获取黑名单缓存 key
func (jwtService *jwtService) getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + helper.MD5([]byte(tokenStr))
}

// JoinBlackList token 加入redis黑名单
// 获取当前token的的过期时间 设置redis的过期时间
func (jwtService *jwtService) JoinBlackList(tokenStr string) error {

	claims, valid, err := helper.ValidateTokenAndBackClaims(tokenStr)

	if err != nil || !valid {
		zap.S().Infof("JoinBlackList: %s", err.Error())
		return err
	}

	nowUnix := time.Now().Unix()

	// redis 中黑名单过期时间
	expTimer := time.Duration(claims.ExpiresAt.Unix() - nowUnix)

	key := jwtService.getBlackListKey(tokenStr)

	if expTimer <= 0 {
		return nil
	}

	// fmt.Printf("key: %s %d %d %s\n", key, expTimer, claims.ExpiresAt.Unix(), tokenStr)
	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
	err = config.RedisClient.SetNX(context.Background(), key, nowUnix, expTimer*time.Second).Err()
	return err
}

// IsInBlacklist token 是否在黑名单中
func (jwtService *jwtService) IsInBlacklist(tokenStr string) bool {
	valueStr, err := config.RedisClient.Get(context.Background(), jwtService.getBlackListKey(tokenStr)).Result()
	// fmt.Printf("valueStr: %s\n", valueStr)
	if valueStr == "" || err != nil {
		return false
	}

	return true
}
