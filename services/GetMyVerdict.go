package services

import (
	"context"
	"encoding/json"

	geminiapi "github.com/Foxtrot-14/FitRang/analysis-service/gemini-api"
	"github.com/Foxtrot-14/FitRang/analysis-service/graph/model"
)

func (s *Service) GetMyVerdict(
	ctx context.Context,
	input model.MyVerdictInput,
) (*model.Verdict, error) {
	email, err := getEmailFromContext(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.p.GetDossierByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if resp.Dossier == nil {
		return nil, nil
	}

	dossierJSON, err := json.Marshal(resp.Dossier)
	if err != nil {
		return nil, err
	}

	raw, err := geminiapi.SendRequest(
		ctx,
		"for me",
		string(dossierJSON),
		input.Product,
	)
	if err != nil {
		return nil, err
	}

	var wrapper verdictWrapper
	if err := json.Unmarshal([]byte(raw), &wrapper); err != nil {
		return nil, err
	}

	return &wrapper.Verdict, nil
}
