package redis_client

import (
	"context"
	"strconv"
)

func GetValuesByKeyRange(key string, n int) []interface{} {
	keys := make([]string, n)

	for i := 0; i <= n; i++ {
		keys[i] = key + ":" + strconv.Itoa(i)
	}

	// Lấy nhiều key cùng lúc using MGet()
	results, err := rdb.MGet(context.Background(), keys...).Result()

	if err != nil {
		panic(err)
	}

	return results
}
