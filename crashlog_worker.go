package main

import (
    "fmt"
    "github.com/benmanns/goworker"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

var EventType = "CrashLog"

type (
    CrashEvent struct {
        Id     bson.ObjectId `json:"id" bson:"_id"`
        Type   string        `json:"type" bson:"type"`
        // Timestamp   string    `json:"timestamp" bson:"timestamp"`
    }
)

func getSession() *mgo.Session {
    s, err := mgo.Dial("mongodb://localhost")

    if err != nil {
        panic(err)
    }
    return s
}

func createEvent(session *mgo.Session) {
    e := CrashEvent{}
    e.Id = bson.NewObjectId()
    e.Type = EventType
    // e.Timestamp = time.Now()

    session.DB("Metrics").C(EventType).Insert(e)
}

func myFunc(queue string, args ...interface{}) error {
    session := getSession()
  	defer session.Close()

    session.SetMode(mgo.Monotonic, true)

    createEvent(session)

    fmt.Printf("%s : From %s, %v\n", EventType, queue, args)
    return nil
}

func init() {
    goworker.Register(EventType, myFunc)
}

func main() {
    if err := goworker.Work(); err != nil {
        fmt.Println("Error:", err)
    }
}
