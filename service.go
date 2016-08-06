package main

import (
	"gopkg.in/redis.v4"
)

type ProphetService interface {
	Consult() (string, error)
	Prepare()
}

type prophetService struct { 
	rclient *redis.Client
}

func (svc *prophetService) Consult() (string, error) {
	v, err := svc.rclient.SRandMember("consult").Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func (svc *prophetService) Prepare() {
//	svc.rclient.SAdd("consult", "Never doubt a prophet.")
	svc.rclient.SAdd("consult", "Is that so.")
	svc.rclient.SAdd("consult", "A prophet is never late.")
}
