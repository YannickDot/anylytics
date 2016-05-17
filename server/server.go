package server

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/kavu/go-resque" // Import this package
    _ "github.com/kavu/go-resque/godis" // Use godis driver

    "anylytics/app"
)

func Init(queue *resque.RedisEnqueuer) *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", app.HandleIndex)
    router.HandleFunc("/events", app.HandleEventsIndex)
    router.HandleFunc("/events/login", app.HandleLogin(queue))
    router.HandleFunc("/events/crash", app.HandleCrash(queue))
    router.HandleFunc("/events/ab", app.HandleABTest(queue))

    log.Fatal(http.ListenAndServe(":8080", router))
    return router
}
