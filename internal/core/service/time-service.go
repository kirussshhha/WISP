package service

import (
	"time"
)

type TimeService interface {
	GetCurrentTime() (string, error)
	GetGreeting() (string, error)
}

type timeService struct{}

func NewTimeService() TimeService {
	return &timeService{}	
}

func (s *timeService) GetCurrentTime() (string, error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime, nil
}

func (s *timeService) GetGreeting() (string, error) {
	hour := time.Now().Hour()
	
	if hour < 12 {
		return "Доброе утро", nil
	} else if hour < 18 {
		return "Добрый день", nil
	} else {
		return "Добрый вечер", nil
	}
}