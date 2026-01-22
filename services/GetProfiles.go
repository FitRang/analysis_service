package services

import (
	"context"

	"github.com/Foxtrot-14/FitRang/analysis-service/graph/model"
)

func (s *Service) GetProfiles(
	ctx context.Context,
) ([]*model.Profile, error) {
	email, err := getEmailFromContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := s.p.GetProfiles(ctx, email)
	if err != nil {
		return nil, err
	}

	var result []*model.Profile
	for _, prof := range resp.Profile {
		profileURL := prof.ProfileUrl
		result = append(result, &model.Profile{
			ID:         prof.Id,
			FullName:   prof.FullName,
			Username:   prof.Username,
			ProfileURL: &profileURL,
			CreatedAt:  prof.CreatedAt,
			UpdatedAt:  prof.UpdatedAt,
		})
	}

	return result, nil
}
