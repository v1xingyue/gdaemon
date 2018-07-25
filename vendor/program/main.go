package program

import (
	"github.com/kardianos/service"
	"time"
)

var (
	PService *service.Service
)

type PObj struct {
}

func (p *PObj) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *PObj) run() {
	logger, _ := (*PService).Logger(nil)
	for {
		logger.Infof("Service %s is running !", (*PService))
		time.Sleep(10 * time.Second)
	}
}
func (p *PObj) Stop(s service.Service) error {
	logger, _ := (*PService).Logger(nil)
	logger.Infof("Service %s is stopped!", (*PService))
	return nil
}
