package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/kavu/go-resque" // Import this package
    _ "github.com/kavu/go-resque/godis" // Use godis driver
    "github.com/simonz05/godis/redis" // Redis client from godis package
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/events", EventsIndex)
    router.HandleFunc("/events/login", LoginHandler)
    router.HandleFunc("/events/crash", CrashHandler)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func EventsIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "EventsIndex!")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    client := redis.New("tcp:127.0.0.1:6379", 0, "") // Create new Redis client to use for enqueuing
    enqueuer := resque.NewRedisEnqueuer("godis", client) // Create enqueuer instance
    var err error

    _, err = enqueuer.Enqueue("resque:queue:login", "Login")
    if err != nil {
      panic(err)
    }

    fmt.Fprintln(w, "LoginHandler!")
}

func CrashHandler(w http.ResponseWriter, r *http.Request) {
    client := redis.New("tcp:127.0.0.1:6379", 0, "") // Create new Redis client to use for enqueuing
    enqueuer := resque.NewRedisEnqueuer("godis", client) // Create enqueuer instance
    var err error

    _, err = enqueuer.Enqueue("resque:queue:crashlog", "CrashLog")
    if err != nil {
      panic(err)
    }

    fmt.Fprintln(w, "CrashHandler!")
}
