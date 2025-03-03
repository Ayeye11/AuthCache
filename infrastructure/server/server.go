package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ayeye11/se-thr/infrastructure/config"
	"github.com/gin-gonic/gin"
)

type server struct {
	conn   *gin.Engine
	config config.ConfigAPP
	errCh  chan error
	signCh chan os.Signal
}

func NewServer(conn *gin.Engine, config config.ConfigAPP) *server {
	errCh := make(chan error, 1)
	signCh := make(chan os.Signal, 1)

	signal.Notify(signCh, syscall.SIGINT, syscall.SIGTERM)
	return &server{conn, config, errCh, signCh}
}

func (s *server) Run() {
	log.Printf("Server listen on %s%s\n", s.config.Host, s.config.Port)
	s.errCh <- s.conn.Run(s.config.Port)
}

func (s *server) Close() error {
	select {
	case err := <-s.errCh:
		return err

	case sign := <-s.signCh:
		fmt.Printf(" => Signal received: %s\n", sign)
		return nil
	}
}
