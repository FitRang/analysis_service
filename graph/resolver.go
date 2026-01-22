package graph

import "github.com/Foxtrot-14/FitRang/analysis-service/proto"

//go:generate go tool gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	ProfileClient *proto.ProfileClient
}
