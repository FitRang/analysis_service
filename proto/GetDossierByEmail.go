package proto

import context "context"

func (p *ProfileClient) GetDossierByEmail(ctx context.Context, email string) (*GetDossierResponse, error) {
	resp, err := p.client.GetDossierEmail(ctx, &GetByEmailRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
