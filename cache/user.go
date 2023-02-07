package cache

import "go-blog/global"

func PushToken(id string, token string) error {
	return global.RedisDB.HSet(ctx, RDB_USERID_TOKEN, id, token).Err()
}

func PullToken(id string) (string, error) {
	cmd := global.RedisDB.HGet(ctx, RDB_USERID_TOKEN, id)

	return cmd.Result()
}
