package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ayeye11/se-thr/infrastructure/config"
	"github.com/gin-gonic/gin"
)

type server struct {
	srv    *http.Server
	config config.ConfigAPP
	errCh  chan error
	signCh chan os.Signal
}

func NewServer(conn *gin.Engine, config config.ConfigAPP) *server {
	errCh := make(chan error, 1)
	signCh := make(chan os.Signal, 1)

	srv := &http.Server{Addr: config.Port, Handler: conn}

	signal.Notify(signCh, syscall.SIGINT, syscall.SIGTERM)
	return &server{srv, config, errCh, signCh}
}

func (s *server) Run() {
	log.Printf("Server listen on %s%s\n", s.config.Host, s.config.Port)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.errCh <- err
	}
}

func (s *server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case err := <-s.errCh:
		return err

	case sign := <-s.signCh:
		fmt.Printf(" => Signal received: %s\n", sign)
	}

	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
