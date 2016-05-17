package main

import (
    "fmt"
    "time"
    "github.com/benmanns/goworker"
    "anylytics/db"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

var CrashEventType = "CrashLog"

type (
    CrashEvent struct {
        Id     bson.ObjectId `json:"id" bson:"_id"`
        Type   string        `json:"type" bson:"type"`
        Timestamp   string    `json:"timestamp" bson:"timestamp"`
    }
)

func createCrashEvent(session *mgo.Session) {
    e := CrashEvent{}
    e.Id = bson.NewObjectId()
    e.Type = CrashEventType
    e.Timestamp = time.Now().Format("2006-01-02 15:04:05")

    session.DB("Metrics").C(CrashEventType).Insert(e)
}

func processCrashEvent(mongodb *mgo.Session) func(queue string, args ...interface{}) error {

    fn := func (queue string, args ...interface{}) error {
      mongodb.SetMode(mgo.Monotonic, true)
      createCrashEvent(mongodb)
      fmt.Printf("%s : From %s, %v\n", CrashEventType, queue, args)
      return nil
    }

    return fn
}

func init() {
    mongodb := db.InitMongoDB()
    goworker.Register(CrashEventType, processCrashEvent(mongodb))
}

func main() {
    if err := goworker.Work(); err != nil {
        fmt.Println("Error:", err)
    }
}
