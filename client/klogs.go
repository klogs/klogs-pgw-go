package klogs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// media type names declaration
const (
	MediaTypeNames_Json           string = "application/json"
	MediaTypeNames_UrlEncodedForm string = "application/x-www-form-urlencoded"
)

func init() {
	fmt.Println()
}

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL

	apiKey                string
	secretKey             string
	defaultRequestHeaders map[string]string

	CardPayment *CardPayment
}

func New(apiKey, secretKey, baseURL string) (*Client, error) {
	client := &Client{apiKey: apiKey, secretKey: secretKey, httpClient: http.DefaultClient}

	client.CardPayment = &CardPayment{Http: client}

	client.defaultRequestHeaders = make(map[string]string)
	client.baseURL, _ = url.Parse(baseURL)

	return client, nil
}

func (cli *Client) Post(ctx context.Context, resourceUri string, body, response interface{}) error {
	req, err := createHttpRequest(ctx, cli, http.MethodPost, resourceUri, &body)

	if err != nil {
		return err
	}

	err = send(ctx, cli, req, &response)

	if err != nil {
		return err
	}

	return nil
}

func (cli *Client) Put(ctx context.Context, resourceUri string, body, response interface{}) error {
	req, err := createHttpRequest(ctx, cli, http.MethodPut, resourceUri, &body)

	if err != nil {
		return err
	}

	err = send(ctx, cli, req, &response)

	if err != nil {
		return err
	}

	return nil
}

func (cli *Client) Get(ctx context.Context, resourceUri string, response interface{}) error {
	req, err := createHttpRequest(ctx, cli, http.MethodGet, resourceUri, nil)

	if err != nil {
		return err
	}

	err = send(ctx, cli, req, &response)

	if err != nil {
		return err
	}

	return nil
}

func (cli *Client) Delete(ctx context.Context, resourceUri string, response interface{}) error {
	req, err := createHttpRequest(ctx, cli, http.MethodDelete, resourceUri, nil)

	if err != nil {
		return err
	}

	err = send(ctx, cli, req, &response)

	if err != nil {
		return err
	}

	return nil
}

func createHttpRequest(ctx context.Context, cli *Client, httpMethod, resourceUri string, body interface{}) (*http.Request, error) {
	uri, err := cli.baseURL.Parse(resourceUri)

	if err != nil {
		return nil, err
	}

	var req *http.Request

	switch httpMethod {
	case http.MethodGet, http.MethodDelete:
		req, err = http.NewRequestWithContext(ctx, httpMethod, uri.String(), nil)

		if err != nil {
			return nil, err
		}
	case http.MethodPost, http.MethodPut:
		buf, err := JsonEncode(body)

		if err != nil {
			return nil, err
		}

		req, err = http.NewRequestWithContext(ctx, httpMethod, uri.String(), buf)

		if err != nil {
			return nil, err
		}
	}

	for k, v := range cli.defaultRequestHeaders {
		req.Header.Add(k, v)
	}

	randomString := UrlFriendlyRandomString(32)
	ticks := strconv.FormatInt(UTCTicks(), 10)

	ciperText := fmt.Sprintf("%s%s%s", cli.apiKey, randomString, ticks)

	signature := HMACSHA256(ciperText, cli.secretKey)

	req.Header.Set("X-Api-Key", cli.apiKey)
	req.Header.Set("X-Klogs-Rnd", randomString)
	req.Header.Set("X-Klogs-Timestamp", ticks)
	req.Header.Set("X-Klogs-Signature", signature)
	req.Header.Set("Content-Type", MediaTypeNames_Json)

	return req, nil
}

func send(ctx context.Context, cli *Client, req *http.Request, responseObj interface{}) error {
	req = req.WithContext(ctx)

	res, err := cli.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if IsSuccessStatusCode(res.StatusCode) {
		if err = json.NewDecoder(res.Body).Decode(&responseObj); err != nil {
			return err
		}
	}

	errObj := &Response{Success: false}

	if err = json.NewDecoder(res.Body).Decode(errObj); err != nil {
		return err
	}

	return errors.New(errObj.Error.Summary)
}
