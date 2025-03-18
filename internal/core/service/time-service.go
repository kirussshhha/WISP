package service

import (
	"time"
)

func (s *Services) GetCurrentTime() (string, error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime, nil
}

func (s *Services) GetGreeting() (string, error) {
	hour := time.Now().Hour()

	if hour < 12 {
		return "Доброе утро", nil
	} else if hour < 18 {
		return "Добрый день", nil
	} else {
		return "Добрый вечер", nil
	}
}

func (s *Services) GetTimeWithFormat(format string) (string, error) {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	currentTime := time.Now().Format(format)
	return currentTime, nil
}

func (s *Services) CalculateTimeDifference(from, to string) (string, error) {
	layout := "2006-01-02 15:04:05"
	t1, err := time.Parse(layout, from)
	if err != nil {
		return "", err
	}
	t2, err := time.Parse(layout, to)
	if err != nil {
		return "", err
	}
	duration := t2.Sub(t1)
	return duration.String(), nil
}
