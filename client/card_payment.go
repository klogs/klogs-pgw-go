package klogs

import (
	"context"
	"net/url"
	"strconv"
)

type CardPayment struct {
	Http *Client
}

func (cli *CardPayment) CreatePaymentToken(ctx context.Context) (*PaymentTokenResponse, error) {
	response := &PaymentTokenResponse{}

	err := cli.Http.Get(ctx, "api/cardPayment/token", response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (cli *CardPayment) Pay(ctx context.Context, model CreatePaymentRequest) (*CardPaymentResponse, error) {

	response := &CardPaymentResponse{}

	err := cli.Http.Post(ctx, "api/cardPayment", model, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (cli *CardPayment) ProvisionCommit(ctx context.Context, model ProvisionCommitRequest) (*Response, error) {

	response := &Response{}

	err := cli.Http.Post(ctx, "api/cardPayment/provisionCommit", model, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (cli *CardPayment) CommissionsByBinAsync(ctx context.Context, parameters PaymentOptionQueryParams) (*PaymentTokenResponse, error) {
	response := &PaymentTokenResponse{}

	params := url.Values{}

	params.Add("amount", strconv.FormatFloat(parameters.Amount, 'f', 2, 64))
	params.Add("binNumber", parameters.BinNumber)
	params.Add("currency", parameters.Currency)

	if len(parameters.CardId) > 0 {
		params.Add("cardId", parameters.CardId)
	}

	if len(parameters.ProductCodes) > 0 {
		for _, code := range parameters.ProductCodes {
			params.Add("productCode", code)
		}
	}

	if len(parameters.ProductCategoryCodes) > 0 {
		for _, code := range parameters.ProductCategoryCodes {
			params.Add("productCategoryCode", code)
		}
	}

	var resourceUri = "api/cardPayment/installments?" + params.Encode()

	err := cli.Http.Get(ctx, resourceUri, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
