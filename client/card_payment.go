package klogs

import (
	"context"
)

type CardPayment struct {
	Http *Client
}

func (cli *CardPayment) Pay(ctx context.Context, model CreatePaymentRequest) (*CardPaymentResponse, error) {

	response := &CardPaymentResponse{}

	err := cli.Http.Post(ctx, "api/cardPayment", model, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
