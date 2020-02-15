package main

import (
	"encoding/json"
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

const islatestkey = "islatest"
const indexpostskey = "indexpostskey"

func setIslatest(b bool) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", islatestkey, b)
	return err
}

func getIslatest() (bool, error) {
	conn := pool.Get()
	defer conn.Close()
	islatest, err := redis.Bool(conn.Do("GET", islatestkey))
	return islatest, err
}

func setIndexPosts(posts []Post) error {
	conn := pool.Get()
	defer conn.Close()

	for _, p := range posts {
		bytes, err := json.Marshal(p)
		if err != nil {
			return err
		}

		_, err = conn.Do("RPUSH", indexpostskey, bytes)
		if err != nil {
			return err
		}
	}

	return nil
}

func getIndexPosts() ([]Post, error) {
	conn := pool.Get()
	defer conn.Close()
	posts := []Post{}

	postsbytes, err := redis.ByteSlices(conn.Do("LRANGE", indexpostskey, 0, -1))

	if err != nil {
		return posts, err
	}

	for _, bytes := range postsbytes {
		post := Post{}
		json.Unmarshal(bytes, &post)
		posts = append(posts, post)
	}

	return posts, nil
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
