package main

import (
	"gopkg.in/redis.v4"
)

type ProphetService interface {
	Consult() (string, error)
}

type prophetService struct { 
	rclient *redis.Client
}

func (svc *prophetService) Consult() (string, error) {
	v, err := svc.rclient.SRandMember("consult").Result()
	if err != nil {
		return "...", nil
	}
	return v, nil
}
