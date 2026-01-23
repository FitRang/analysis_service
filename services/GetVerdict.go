package services

import (
	"context"
	"encoding/json"

	geminiapi "github.com/Foxtrot-14/FitRang/analysis-service/gemini-api"
	"github.com/Foxtrot-14/FitRang/analysis-service/graph/model"
)

func (s *Service) GetVerdict(
	ctx context.Context,
	input model.VerdictInput,
) (*model.Verdict, error) {
	resp, err := s.p.GetDossierByUsername(ctx, input.Username)
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
		input.Username,
		string(dossierJSON),
		input.Product,
	)
	if err != nil {
		return nil, err
	}
	clean := extractJSON(raw)
	var wrapper verdictWrapper
	if err := json.Unmarshal([]byte(clean), &wrapper); err != nil {
		return nil, err
	}

	return finalSerialize(&wrapper.Verdict), nil
}
