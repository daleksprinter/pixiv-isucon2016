package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

func initializeRedis() {
	conn := pool.Get()
	defer conn.Close()
	conn.Do("FLUSHALL")
}

func userCommentKey(userID int) string {
	return fmt.Sprintf("user-comments:%d", userID)
}

func addComment(c *Comment) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("ZADD", userCommentKey(c.UserID), c.CreatedAt.UnixNano(), c.ID)
	return err
}

func countComment(userID int) (int, error) {
	conn := pool.Get()
	defer conn.Close()

	count, err := redis.Int(conn.Do("ZCARD", userCommentKey(userID)))
	if err != nil {
		return 0, err
	}
	return count, nil
}
