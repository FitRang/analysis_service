package services

import (
	"encoding/json"
	"math"
	"strings"

	"github.com/Foxtrot-14/FitRang/analysis-service/graph/model"
	"github.com/Foxtrot-14/FitRang/analysis-service/proto"
)

type Service struct {
	p *proto.ProfileClient
}

func NewService(p *proto.ProfileClient) *Service {
	return &Service{
		p: p,
	}
}

type aiVerdict struct {
	Analysis    string      `json:"analysis"`
	Score       json.Number `json:"score"`
	Suitability string      `json:"suitability"`
}

type verdictWrapper struct {
	Verdict aiVerdict `json:"verdict"`
}

func extractJSON(s string) string {
	s = strings.TrimSpace(s)

	if strings.HasPrefix(s, "```") {
		s = strings.TrimPrefix(s, "```json")
		s = strings.TrimPrefix(s, "```")
		s = strings.TrimSuffix(s, "```")
	}

	return strings.TrimSpace(s)
}

func finalSerialize(verdict *aiVerdict) *model.Verdict {
	scoreFloat, err := verdict.Score.Float64()
	if err != nil {
		scoreFloat = 0
	}

	score := int32(math.Round(scoreFloat))

	if score == min(score, 0) {
		score = 0
	}
	if score > 10 {
		score = 10
	}

	return &model.Verdict{
		Analysis:    verdict.Analysis,
		Score:       score,
		Suitability: verdict.Suitability,
	}
}
