package main

import (
	"github.com/kardianos/service"
	"log"
	"os"
	"time"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *program) run() {
	for {
		time.Sleep(60 * time.Second)
		logger.Info("sservice running ")
	}
}
func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "sservice",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	logger, err = s.Logger(nil)
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			logger.Info("installed ok")
			return
		}
		if os.Args[1] == "remove" {
			s.Uninstall()
			logger.Info("remove ok")
			return
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
