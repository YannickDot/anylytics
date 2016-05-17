package db

import (
  "github.com/kavu/go-resque" // Import this package
  _ "github.com/kavu/go-resque/godis" // Use godis driver
  "github.com/simonz05/godis/redis" // Redis client from godis package

  "gopkg.in/mgo.v2"
)

func InitRedisQueue() *resque.RedisEnqueuer {
  client := redis.New("tcp:127.0.0.1:6379", 0, "") // Create new Redis client to use for enqueuing
  enqueuer := resque.NewRedisEnqueuer("godis", client) // Create enqueuer instance
  return enqueuer
}

func InitMongoDB() *mgo.Session {
  s, err := mgo.Dial("mongodb://localhost")
  if err != nil {
      panic(err)
  }
  return s
}
