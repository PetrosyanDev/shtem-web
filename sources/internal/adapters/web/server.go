// Erik Petrosyan ©
package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	port = 2222
)

type server struct {
	http.Server
}

type WEBServerOpt func(*server) error

func (s *server) Start() error {
	log.Printf("starting WEB server on port: %d with %d CPU cores ...", port, runtime.NumCPU())
	if s.TLSConfig == nil {
		return s.ListenAndServe()
	}
	return s.ListenAndServeTLS("", "")
}

func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.SetKeepAlivesEnabled(false)
	if err := s.Shutdown(ctx); err == nil {
		log.Println("WEB server stopped")
		return nil
	} else {
		return err
	}
}

func NewWEBServer(hdl *gin.Engine, opts ...WEBServerOpt) (*server, error) {
	srv := &server{}
	srv.Addr = fmt.Sprintf(":%d", port)
	srv.Handler = hdl
	return srv, nil
}
