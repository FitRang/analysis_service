package services

import "github.com/Foxtrot-14/FitRang/analysis-service/proto"

type Service struct {
	p *proto.ProfileClient
}

func NewService(p *proto.ProfileClient) *Service {
	return &Service{
		p: p,
	}
}
