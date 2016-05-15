package main

import (
  "github.com/kavu/go-resque" // Import this package
  _ "github.com/kavu/go-resque/godis" // Use godis driver
  "github.com/simonz05/godis/redis" // Redis client from godis package
)

func main() {
  var err error

  client := redis.New("tcp:127.0.0.1:6379", 0, "") // Create new Redis client to use for enqueuing
  enqueuer := resque.NewRedisEnqueuer("godis", client) // Create enqueuer instance

  // Enqueue into the "default" queue with passing one parameter to the Demo::Job.perform
  var i = 0
  for i < 20 {
    _, err = enqueuer.Enqueue("resque:queue:login", "Login", i)
    if err != nil {
      panic(err)
    }
    i++
  }

  var j = 0
  for j < 20 {
    _, err = enqueuer.Enqueue("resque:queue:crashlog", "CrashLog", j)
    if err != nil {
      panic(err)
    }
    j++
  }



}
