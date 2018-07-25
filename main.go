package main

import (
	"flag"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"program"
)

var (
	Action             string
	DebugMode          *bool
	logger             service.Logger
	ServiceName        string = "gdaemon"
	ServiceDisplayName string = "golang daemon"
	ServiceDescription string = "One Golang Daemon Process"
)

func PrintDefaultHelp() bool {
	fmt.Println("")
	flag.Usage()
	fmt.Println("")
	return true
}

func CommonAction(s_p *service.Service) bool {
	flag := true
	if Action == "" {
		flag = PrintDefaultHelp()
	}
	s := *s_p
	logger, _ = s.Logger(nil)
	if Action != "" {
		if Action == "install" {
			s.Install()
			logger.Infof("service %s installed ok", ServiceName)
		} else if Action == "uninstall" {
			s.Stop()
			s.Uninstall()
			logger.Infof("service %s removed done", ServiceName)
		} else if Action == "stop" {
			err := s.Stop()
			if err != nil {
				logger.Error(err)
			}
			logger.Infof("service %s stopped ", ServiceName)
		} else if Action == "start" {
			err := s.Run()
			if err != nil {
				logger.Error(err)
			}
		} else {
			flag = PrintDefaultHelp()
		}
	}
	return flag
}

func main() {
	flag.StringVar(&Action, "a", "", "action to do <install,uninstall,start,stop>")
	DebugMode = flag.Bool("d", false, "run as debug mode")
	flag.Parse()
	svcConfig := &service.Config{
		Name:        ServiceName,
		DisplayName: ServiceDisplayName,
		Description: ServiceDescription,
		Arguments:   []string{"-a", "start"},
	}

	prg := &program.PObj{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	program.PService = &s
	CommonAction(&s)
}
