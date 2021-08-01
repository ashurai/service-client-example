package form3

import (
	"context"
	"fmt"
)

const (
	orgAccBasePath = "organisation/accounts"
)

type AccService service

func (as *AccService) GetByID(ctx context.Context, id string) (*AccountData, *Response, error) {
	cl := as.form3Client
	url := fmt.Sprintf("%s/%s", orgAccBasePath, id)
	req, err := cl.GET(url, nil)
	if err != nil {
		return nil, nil, nil
	}

	account := &AccountData{}
	res, err := cl.Do(ctx, req, account)
	if err != nil {
		return nil, res, err
	}
	return account, res, nil
}

func (as *AccService) NewAccount(ctx context.Context, acc *AccountData) (*AccountData, *Response, error) {
	cl := as.form3Client

	req, err := cl.POST(orgAccBasePath, acc)
	if err != nil {
		return nil, nil, err
	}

	account := &AccountData{}

	res, err := cl.Do(ctx, req, account)
	if err != nil {
		return nil, res, err
	}

	return account, res, nil
}

func (as *AccService) DeleteByID(ctx context.Context, id string, version int) (*Response, error) {
	cl := as.form3Client

	url := fmt.Sprintf("%s/%s?version=%d", orgAccBasePath, id, version)
	reqs, err := cl.DELETE(url, nil)
	if err != nil {
		return nil, nil
	}

	res, err := cl.Do(ctx, reqs, nil)
	if err != nil {
		return res, err
	}

	return res, nil
}

func CreateAccountsService(client *Client) *AccService {
	if client == nil {
		client = CreateClient(nil)
	}
	return &AccService{
		form3Client: client,
	}
}
