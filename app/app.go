package app

import (
    "fmt"
    "strconv"
    "log"
    "net/http"
    "math/rand"
    "github.com/kavu/go-resque" // Import this package
    _ "github.com/kavu/go-resque/godis" // Use godis driver
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}


func HandleEventsIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "EventsIndex!")
}

func HandleLogin(enqueuer *resque.RedisEnqueuer) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
      _, err := enqueuer.Enqueue("resque:queue:login", "Login")
      if err != nil {
        log.Fatal("Couldn't find Redis", err)
      }

      fmt.Fprintln(w, "LoginHandler!")
    }

    return http.HandlerFunc(fn)
}

func HandleCrash(enqueuer *resque.RedisEnqueuer) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
      _, err := enqueuer.Enqueue("resque:queue:crashlog", "CrashLog")
      if err != nil {
        log.Fatal("Couldn't find Redis", err)
      }

      fmt.Fprintln(w, "CrashHandler!")
    }

    return http.HandlerFunc(fn)
}

func HandleABTest(enqueuer *resque.RedisEnqueuer) http.HandlerFunc {
    fn := func(w http.ResponseWriter, r *http.Request) {
      testType := strconv.Itoa(rand.Intn(100) % 2)
      _, err := enqueuer.Enqueue("resque:queue:abtest", "ABTest", testType)
      if err != nil {
        log.Fatal("Couldn't find Redis", err)
      }

      fmt.Fprintln(w, "ABTestingHandler! Variation #", testType)
    }

    return http.HandlerFunc(fn)
}
