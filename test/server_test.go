package test

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// StartServer starts a server with the test handler behind a route
// appropriate for the path parameters to be decoded.
// The function blocks until the something listening on the server's port.
// The caller should defer the returned function to shutdown the server.
func StartServer() func() {
	r := chi.NewRouter()
	r.Post("/records/{b}", Post)
	s := http.Server{
		Addr:              fmt.Sprintf("localhost:%d", port),
		Handler:           r,
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second,
	}
	d := make(chan struct{})
	go func() {
		l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
		if err != nil {
			panic(err)
		}
		close(d)
		err = s.Serve(l)
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-d
	return func() {
		err := s.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}
}

// port is the port on which the test server is started.
const port = 24842
