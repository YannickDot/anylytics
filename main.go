package main

import (
  "anylytics/db"
  "anylytics/server"
)

func main() {
  queue := db.InitRedisQueue()
  server.Init(queue)
}
