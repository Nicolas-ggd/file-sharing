package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

// Application struct is a wrap of application, which controls everything and have top management role, everything is united arround this struct.
type Application struct {
	debug    *bool
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// parse addr flag to dynamicaly change address value
	addr := flag.String("addr", ":8080", "HTTP network address")

	// parse debug mode to enable debug mode in application
	debug := flag.Bool("debug", false, "Enable debug mode")
	// define logs
	infoLog := log.New(os.Stderr, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Llongfile)

	flag.Parse()

	if *debug {
		infoLog.Println("Enabled debug mode...")
	}

	app := &Application{
		debug:    debug,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := http.Server{
		Addr:         *addr,
		Handler:      app.NewServer(),
		ErrorLog:     errorLog,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Server starting on %s", *addr)

	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatalln(err)
	}
}
