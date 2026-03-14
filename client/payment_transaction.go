package klogs

import "context"

type PaymentTransactionClient struct {
	Http *Client
}

func (cli *PaymentTransactionClient) Detail(ctx context.Context, referenceCode string) (*TransactionDetailResponse, error) {
	response := &PaymentTransactionDetail{}

	err := cli.Http.Get(ctx, "api/trx/"+referenceCode, response)

	if err != nil {
		return nil, err
	}

	return &TransactionDetailResponse{Transaction: *response}, nil
}

func (cli *PaymentTransactionClient) Refund(ctx context.Context, model RefundRequest) (*Response, error) {
	response := &Response{}

	err := cli.Http.Post(ctx, "api/trx/refund", model, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (cli *PaymentTransactionClient) Void(ctx context.Context, model VoidRequest) (*Response, error) {
	response := &Response{}

	err := cli.Http.Post(ctx, "api/trx/void", model, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
