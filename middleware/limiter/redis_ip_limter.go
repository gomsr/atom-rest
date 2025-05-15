package limiter

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RedisLimit struct {
	LimitBase
	localCache sync.Map
	REDIS      *redis.Client
}

func (l *RedisLimit) CheckOrMark(key string, expire int, limit int) error {
	if l.REDIS == nil {
		return errors.New("redis is not init")
	}

	count, err := l.REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if count == 0 {
		pipe := l.REDIS.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, time.Duration(expire)*time.Second)
		_, err = pipe.Exec(context.Background())
		return err
	}

	// 次数
	if times, err := l.REDIS.Get(context.Background(), key).Int(); err != nil {
		return err
	} else {
		if times < limit {
			return l.REDIS.Incr(context.Background(), key).Err()
		}

		if t, err := l.REDIS.PTTL(context.Background(), key).Result(); err != nil {
			return errors.New("请求太过频繁，请稍后再试")
		} else {
			return errors.New("请求太过频繁, 请 " + t.String() + " 秒后尝试")
		}
	}
}

func RedisLimiter(limitTimeIP, limitCountIP int) gin.HandlerFunc {

	limiter := RedisLimit{}
	limiter.Expire = limitTimeIP
	limiter.Limit = limitCountIP

	return limiter.Process(limiter.CheckOrMark, limiter.Expire, limiter.Limit)
}
