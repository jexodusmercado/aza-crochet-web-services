package main

import (
    "context"
    "flag"
    "log"
    "os"
    "os/signal"
    "time"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

	"aza-crochet-web-services/api/routes"
	"aza-crochet-web-services/api/middleware"
    
)

func main() {
    var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

    log.Println("starting...")
    log.Println("addr: 127.0.0.1:10000")
    r := mux.NewRouter().StrictSlash(true)

    cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type","Authorization"}),
		handlers.AllowedOrigins([]string{"127.0.0.1:10000"}),
        handlers.AllowedMethods([]string{"POST","GET","DELETE","PUT","OPTIONS"}),
        handlers.MaxAge(3600),
	)

    r.Use(cors)
    r.Use(middleware.LoggingMiddleware)

    routes.HandleRequests(r)

    srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:10000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    c := make(chan os.Signal, 1)

    signal.Notify(c, os.Interrupt)

    <-c

    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()

    srv.Shutdown(ctx)

    log.Println("shutting down..")
    os.Exit(0)
}