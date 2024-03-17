package api

import (
	"context"
	"github.com/nobbyphala/Brick/external/http_request"
)

type bankApiClient struct {
	baseUrl     string
	httpRequest http_request.HTTPRequest
}

type BankApiClientOpts struct {
	BaseUrl     string
	HttpRequest http_request.HTTPRequest
}

func NewBankApiClient(opts BankApiClientOpts) *bankApiClient {
	return &bankApiClient{
		baseUrl:     opts.BaseUrl,
		httpRequest: opts.HttpRequest,
	}
}

func (cl bankApiClient) VerifyAccount(ctx context.Context, account VerifyAccountRequest) (VerifyAccountResponse, error) {
	url := cl.baseUrl + "/verify"
	var response VerifyAccountResponse

	err := cl.httpRequest.Post(ctx, url, nil, account, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (cl bankApiClient) TransferMoney(ctx context.Context, transfer TransferRequest) (TransferResponse, error) {
	url := cl.baseUrl + "/transfer"
	var response TransferResponse

	err := cl.httpRequest.Post(ctx, url, nil, transfer, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
