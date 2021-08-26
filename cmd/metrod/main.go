package main

import (
	"fmt"
	"io"
	"os"

	"github.com/acaee/metro/service/api"
	"go.uber.org/zap"
)

func main() {
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}

type Main struct {
	Stdin io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewMain() *Main {
	return &Main{Stdin: os.Stdin,Stdout: os.Stdout, Stderr: os.Stderr}
}




func (m Main) Run(args ...string) error {
	// TODO 这里写main逻辑
	var server Server
	server.config.Api.Host = "localhost"
	server.config.Api.Post = 2000
	server.Logger = zap.NewExample()
	server.Logger.Info("配置");
	server.appendApiService()
	server.Open()
	return nil
}

type Server struct {



	Logger *zap.Logger

	Services []Service
	config Config
}


type Config struct {
	Api api.Config `toml:Api`
	Port int
}

type Service interface {
	WithLogger(log *zap.Logger)
	Open() error
	Close() error
}

func (server *Server) appendSynchronizationService() error {
	return nil
}

func (server *Server) appendApiService() error {

	var service api.Service
	service.Config = server.config.Api
	server.Services = append(server.Services, &service)
	return nil
}


func (server *Server) Open() {
	for _, service := range server.Services {
		service.WithLogger(server.Logger)
		service.Open()
	}
}
