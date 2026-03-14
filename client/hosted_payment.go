package klogs

import (
	"context"
)

type HostedPayment struct {
	Http *Client
}

func (cli *HostedPayment) CreatePayment(ctx context.Context, model HostedPaymentRequest) (*HostedPaymentResponse, error) {
	response := &HostedPaymentResponse{}

	err := cli.Http.Post(ctx, "api/payment", model, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
