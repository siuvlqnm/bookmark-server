package utils

import (
	"fmt"
	"time"

	"github.com/siuvlqnm/bookmark/global"
)

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
	err = global.GVA_REDIS.Set(key, val, 3*86400*time.Second).Err()
	return
}
