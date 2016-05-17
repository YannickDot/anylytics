package main

import (
    "fmt"
    "time"
    "github.com/benmanns/goworker"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "anylytics/db"
)

var LoginEventType = "Login"

type (
    LoginEvent struct {
        Id     bson.ObjectId `json:"id" bson:"_id"`
        Type   string        `json:"type" bson:"type"`
        Timestamp   string    `json:"timestamp" bson:"timestamp"`
    }
)

func createLoginEvent(session *mgo.Session) {
    e := LoginEvent{}
    e.Id = bson.NewObjectId()
    e.Type = LoginEventType
    e.Timestamp = time.Now().Format("2006-01-02 15:04:05")

    session.DB("Metrics").C(LoginEventType).Insert(e)
}

func processLoginEvent(mongodb *mgo.Session) func(queue string, args ...interface{}) error {

    fn := func (queue string, args ...interface{}) error {
      mongodb.SetMode(mgo.Monotonic, true)
      createLoginEvent(mongodb)
      fmt.Printf("%s : From %s, %v\n", LoginEventType, queue, args)
      return nil
    }

    return fn
}


func init() {
    mongodb := db.InitMongoDB()
    goworker.Register(LoginEventType, processLoginEvent(mongodb))
}


func main() {
    if err := goworker.Work(); err != nil {
        fmt.Println("Error:", err)
    }
}
