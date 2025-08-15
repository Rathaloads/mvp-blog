package db

import (
	"time"

	"mb-server/common/config"
	"mb-server/common/logger"

	"github.com/go-redis/redis"
)

const (
	REDIS_TAG = "redis"
)

type RedisPipeliner struct {
	redisPipeliner redis.Pipeliner
}

var (
	redisClient *redis.Client
)

func connectRedis(c *config.Redis) error {
	//连接服务器
	redisClient = redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       c.DB})

	_, err := redisClient.Ping().Result()
	if err != nil {
		logger.Errorf("init redis err " + err.Error())
		return err
	}
	logger.Infof("init redis successs: %v", c.Host)
	return nil
}

func RedisHSet(key string, failed string, value interface{}) error {
	_, err := redisClient.HSet(key, failed, value).Result()
	if err != nil {
		logger.Errorf("RedisHSet failed [%v][%v] [err=%v]", key, failed, err)
		return err
	}
	return nil
}

func RedisHIncrBy(key string, failed string, value int64) (int64, error) {
	dst, err := redisClient.HIncrBy(key, failed, value).Result()
	if err != nil {
		logger.Errorf("RedisHIncrBy failed [%v][%v] [err=%v]", key, failed, err)
		return 0, err
	}
	return dst, nil
}

func RedisHSetNX(key string, failed string, value interface{}) bool {
	res, err := redisClient.HSetNX(key, failed, value).Result()
	if err != nil {
		logger.Errorf("RedisHSetNX failed [%v][%v] [err=%v]", key, failed, err)
		return false
	}
	return res
}

func RedisHGet(key string, failed string) (string, error) {
	str, err := redisClient.HGet(key, failed).Result()
	if err != nil && err != redis.Nil {
		logger.Errorf("RedisHGet failed [%v][%v] [err=%v]", key, failed, err)
		return str, err
	}
	return str, nil
}

func RedisHGetAll(key string) (map[string]string, error) {
	str, err := redisClient.HGetAll(key).Result()
	if err != nil && err != redis.Nil {
		logger.Errorf("RedisHGet failed [%v] [err=%v]", key, err)
		return nil, err
	}
	return str, nil
}

func RedisHDel(key string, failed string) error {
	_, err := redisClient.HDel(key, failed).Result()
	return err
}

func RedisSAdd(key string, value string) error {
	_, err := redisClient.SAdd(key, value).Result()
	if err != nil {
		logger.Errorf("RedisSAdd failed [%v] [err=%v]", key, err)
		return err
	}
	return nil
}

func RedisSRandMember(key string) (string, error) {
	str, err := redisClient.SRandMember(key).Result()
	if err != nil && err != redis.Nil {
		logger.Errorf("RedisSRandMember failed [%v] [err=%v]", key, err)
		return str, err
	}
	return str, nil
}

func RedisSRem(key string, value string) (int64, error) {
	cnt, err := redisClient.SRem(key, value).Result()
	if err != nil && err != redis.Nil {
		logger.Errorf("RedisSRem failed [%v] [err=%v]", key, err)
		return cnt, err
	}
	return cnt, nil
}

func RedisSIsMember(key string, failed string) (bool, error) {
	b, err := redisClient.SIsMember(key, failed).Result()
	if err != nil {
		logger.Errorf("RedisSIsMember failed [%v] [err=%v]", key, err)
		return false, err
	}
	return b, nil
}

func RedisPipeline() *RedisPipeliner {
	return &RedisPipeliner{
		redisPipeliner: redisClient.TxPipeline(),
	}
}

func (p RedisPipeliner) HSet(tableName string, key string, object interface{}) error {
	_, err := p.redisPipeliner.HSet(tableName, key, object).Result()
	if err != nil {
		logger.Errorf("redis save object failed [%v:%v:%v] [err=%v]", tableName, key, object, err)
		return err
	}
	return nil
}

func (p RedisPipeliner) SAdd(tableName string, value string) error {
	_, err := p.redisPipeliner.SAdd(tableName, value).Result()
	if err != nil {
		logger.Errorf("redis save object failed [%v:%v] [err=%v]", tableName, value, err)
		return err
	}
	return nil
}

func (p RedisPipeliner) Exec() error {
	_, err := p.redisPipeliner.Exec()
	if err != nil {
		logger.Errorf("redis pipeliner exec failed [err=%v]", err)
		return err
	}
	return nil
}

// 设置Redis string 并加入过期时间
func RedisSetEx(tableName string, key string, value string, time_sec time.Duration) error {
	t_key := tableName + ":" + key
	err := redisClient.Set(t_key, value, time_sec).Err()
	if err != nil {
		logger.Errorf("redis setex failed [%v:%v:%v:%v] [err=%v]", tableName, key, value, time_sec, err)
		return err
	}

	return nil
}

// 读限时Redis表
func RedisGetEx(tableName string, key string) (string, error) {
	t_key := tableName + ":" + key
	value, err := redisClient.Get(t_key).Result()
	if err != nil {
		logger.Errorf("redis get failed [%v:%v] [err=%v]", tableName, key, err)
		return "", err
	}

	return value, nil
}

// 当key 不存在时，设置Redis string 并加入过期时间，key 存在返回false
func RedisSetExNx(tableName string, key string, value string, time_sec time.Duration) bool {
	t_key := tableName + ":" + key
	res, err := redisClient.SetNX(t_key, value, time_sec).Result()
	if err != nil {
		logger.Errorf("RedisSetExNx failed [%v:%v] [err=%v]", tableName, key, err)
		return true
	}
	return res
}

// 删除字段
func RedisDel(tableName string, key string) {
	t_key := tableName + ":" + key
	redisClient.Del(t_key)
}
