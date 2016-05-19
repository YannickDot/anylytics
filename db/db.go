package db

import (
  "github.com/kavu/go-resque" // Import this package
  _ "github.com/kavu/go-resque/godis" // Use godis driver
  "github.com/simonz05/godis/redis" // Redis client from godis package

  "gopkg.in/mgo.v2"
)

func InitRedisQueue() *resque.RedisEnqueuer {
  LOCAL_URL := "tcp:127.0.0.1:6379"
  NET_URL := "redis://h:p42fot0r6meaf4199koelefk9b9@ec2-54-217-222-237.eu-west-1.compute.amazonaws.com:12429"
  client := redis.New(NET_URL, 0, "") // Create new Redis client to use for enqueuing
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
