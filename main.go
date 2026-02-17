package main

import (
	"context"
	"fmt"

	klogs "github.com/klogs/klogs-pgw-go/client"
)

func main() {
	client, _ := klogs.New(
		"lrM54xgeBRw6kABrmyz5GixNW54Eg9zWt3Orgi35E",
		"G99T1V+bzzfU+X0Zv+xvCB4LwLstYtymL8ybsZjvdLGzl98EuNh3AeYUCA1pAOYa6rxv3Y5HsFvhs2v3ufx+nQ==",
		"https://pgw.klogs.dev",
	)

	resp, err := client.CardPayment.Pay(context.Background(), klogs.CreatePaymentRequest{
		Amount:        15,
		Installment:   1,
		ReferenceCode: "TEST-123123123",
		Card: &klogs.CreditCard{
			CardHolderName: "Nadir Yıldız",
			CardNumber:     "5526080000000006",
			Cvv:            "423",
			ExpireMonth:    4,
			ExpireYear:     2027,
		},
		Explanation: "Test from golang client",
		Use3d:       true,
		Currency:    "TRY",
		Email:       "info@klogs.io",
		Phone:       "5554443322",
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Behavior)
		fmt.Println(resp.Success)
	}
}
