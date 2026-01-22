package proto

import context "context"

func (p *ProfileClient) GetProfiles(ctx context.Context, email string) (*GetProfilesResponse, error) {
	resp, err := p.client.GetProfilesByEmail(ctx, &GetByEmailRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
