package main

import (
    "fmt"
    "time"
    "github.com/benmanns/goworker"
    "anylytics/db"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

var ABTestEventType = "ABTest"

type (
    ABTestEvent struct {
        Id     bson.ObjectId `json:"id" bson:"_id"`
        Type   string        `json:"type" bson:"type"`
        Variation   string    `json:"variation" bson:"variation"`
        Timestamp   string    `json:"timestamp" bson:"timestamp"`
    }
)

func createABTestEvent(session *mgo.Session, variation string) {
    e := ABTestEvent{}
    e.Id = bson.NewObjectId()
    e.Type = ABTestEventType
    e.Timestamp = time.Now().Format("2006-01-02 15:04:05")
    e.Variation = variation

    session.DB("Metrics").C(ABTestEventType).Insert(e)
}

func processABTestEvent(mongodb *mgo.Session) func(queue string, args ...interface{}) error {

    fn := func (queue string, args ...interface{}) error {
      variation := args[0].(string)
      mongodb.SetMode(mgo.Monotonic, true)
      createABTestEvent(mongodb, variation)
      fmt.Printf("%s : From %s, %v\n", ABTestEventType, queue, args)
      return nil
    }

    return fn
}

func init() {
    mongodb := db.InitMongoDB()
    goworker.Register(ABTestEventType, processABTestEvent(mongodb))
}

func main() {
    if err := goworker.Work(); err != nil {
        fmt.Println("Error:", err)
    }
}
