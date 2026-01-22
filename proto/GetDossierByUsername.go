package proto

import context "context"

func (p *ProfileClient) GetDossierByUsername(ctx context.Context, username string) (*GetDossierResponse, error) {
	resp, err := p.client.GetDossierByUsername(ctx, &GetByUsernameRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
