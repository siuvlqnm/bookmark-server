package utils

import (
	"fmt"
	"time"

	"github.com/siuvlqnm/bookmark/global"
)

const expTime = 180 * time.Second

// const expTime = 3 * 86400 * time.Second

func joinRedisString(str string, i interface{}) string {
	return fmt.Sprintf("%s%s%d%s", str, ":", i, ":")
}

func GetSetValue(str string, i interface{}) (val string, err error) {
	key := joinRedisString(str, i)
	val, err = global.GVA_REDIS.Get(key).Result()
	return
}

func SetSetValue(str string, i interface{}, val interface{}) (err error) {
	key := joinRedisString(str, i)
	err = global.GVA_REDIS.Set(key, val, expTime).Err()
	return
}

func GetHashValue(str string, i interface{}) (val map[string]string, err error) {
	key := joinRedisString(str, i)
	val, err = global.GVA_REDIS.HGetAll(key).Result()
	return
}

func SetHashValue(str string, i interface{}, val map[string]interface{}) (err error) {
	key := joinRedisString(str, i)
	global.GVA_REDIS.Set(key, expTime, expTime)
	err = global.GVA_REDIS.HMSet(key, val).Err()
	return
}
