/*
 * Maintained by jemo from 2020.2.15 to now
 * Created by jemo on 2020.2.15 11:01:55
 * Db
 */

package db

import (
  "github.com/go-redis/redis/v7"
)

var RedisClient *redis.Client

func Init() {
  RedisClient = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    DB: 1,
  });
}
