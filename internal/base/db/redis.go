package db

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"hongcha/go-expire/internal/base/logger"
	"strconv"
	"time"
)

// RedisClient 全局 redis 客户端
var RedisClient *redis.Client

// InitRedis 初始化 redis 客户端
func InitRedis(conn, username, password string) {
	client := redis.NewClient(&redis.Options{
		Addr:     conn,
		Username: username,
		Password: password,
	})
	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		// 其实可以直接 panic 连接不上 redis 会影响服务可用
		panic(err)
	}
	RedisClient = client
}

// GetCacheBytes 从 redis 中获取 Bytes 类型的缓存
func GetCacheBytes(key string) (resp []byte, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	resp, err = RedisClient.Get(ctx, key).Bytes()
	return
}

// SetCacheJson 将 interface 类型的值放入 redis 缓存中
func SetCacheJson(key string, value interface{}, cacheTime time.Duration) (err error) {
	bytes, _ := json.Marshal(value)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.Set(ctx, key, bytes, cacheTime).Err()
	if err != nil {
		logger.Errorf("redis cache failed %s", err.Error())
	}
	return err
}

// GetCacheString 从 redis 中获取 string 类型的缓存
func GetCacheString(key string) (val string, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	resp := RedisClient.Get(ctx, key)
	val, err = resp.Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

// SetCacheString 将 string 类型的值放入 redis 缓存中
func SetCacheString(key, value string, cacheTime time.Duration) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.Set(ctx, key, value, cacheTime).Err()
	if err != nil {
		logger.Errorf("redis cache failed %s", err.Error())
	}
	return err
}

// DelCacheString 从 redis 中删除 string 类型的缓存
func DelCacheString(key string) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.Del(ctx, key).Err()
	if err != nil {
		logger.Errorf("redis cache failed %s", err.Error())
	}
	return err
}

// 默认没有key值时为-1
func GetCacheInt(key string) (int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	resp := RedisClient.Get(ctx, key)
	if resp.Val() == "" {
		return -1, nil
	}
	return strconv.Atoi(resp.Val())
}

// SetCacheString 将 string 类型的值放入 redis 缓存中
func SetCacheInt(key string, value int, cacheTime time.Duration) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.Set(ctx, key, value, cacheTime).Err()
	if err != nil {
		logger.Error("redis cache failed %s", err.Error())
	}
	return err
}
func SetLock(key string, cacheTime time.Duration) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	if ok, _ := RedisClient.SetNX(ctx, key, 1, cacheTime).Result(); !ok {
		return errors.New("redis lock: already locked")
	}
	return nil
}
func UnLock(key string) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	if err := RedisClient.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
func SetCacheMap(key string, vales ...interface{}) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.SAdd(ctx, key, vales).Err()
	if err != nil {
		logger.Error("redis cache failed %s", err.Error())
	}
	return err
}
func DelCacheMap(key string, vales ...interface{}) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.SRem(ctx, key, vales).Err()
	if err != nil {
		logger.Error("redis DelCacheMap failed %s", err.Error())
	}
	return err
}
func CountCacheMap(key string) (count int64, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	count, err = RedisClient.SCard(ctx, key).Result()
	if err != nil {
		logger.Error("redis CountCacheMap failed %s", err.Error())
		return 0, err
	}
	return count, err
}
func GetCacheMap(key string) (rst []string, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	rst, err = RedisClient.SMembers(ctx, key).Result()
	if err != nil {
		logger.Error("redis GetCacheMap failed %s", err.Error())
	}
	return rst, err
}
func RedisExpire(key string, cacheTime time.Duration) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.Expire(ctx, key, cacheTime).Err()
	if err != nil {
		logger.Error("redis GetCacheMap failed %s", err.Error())
	}
	return err
}
func SetCacheList(key string, vales ...interface{}) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	err = RedisClient.LPush(ctx, key, vales).Err()
	if err != nil {
		logger.Error("redis SetCacheList failed %s", err.Error())
	}
	return err
}
func CountCacheList(key string) (count int64, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	count, err = RedisClient.LLen(ctx, key).Result()
	if err != nil {
		logger.Error("redis CountCacheMap failed %s", err.Error())
		return 0, err
	}
	return count, err
}
func GetCacheList(key string, start, end int64) (rst []string, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	rst, err = RedisClient.LRange(ctx, key, start, end).Result()
	if err != nil {
		logger.Error("redis GetCacheMap failed %s", err.Error())
	}
	return rst, err
}
func GetCacheListAll(key string) (rst []string, err error) {
	total, err := CountCacheList(key)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return rst, nil
	}
	return GetCacheList(key, 0, total)
}

// GetCacheKeyList 获取所有的key
func GetCacheKeyList() (keys []string, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	var cursor uint64
	keys, cursor, err = RedisClient.Scan(ctx, cursor, "*", 100).Result()
	if err != nil {
		logger.Error("redis GetCacheKeyList failed %s", err.Error())
		return
	}

	return keys, nil
}
