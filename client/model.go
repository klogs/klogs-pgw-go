package klogs

type CreditCard struct {
	CardHolderName string `json:"cardHolderName,omitempty"`
	CardNumber     string `json:"cardNumber,omitempty"`
	Cvv            string `json:"cvv,omitempty"`
	ExpireMonth    int    `json:"expireMonth"`
	ExpireYear     int    `json:"expireYear"`
}

type Reward struct {
	Amount    float64 `json:"amount"`
	UseReward bool    `json:"useReward"`
}

type Address struct {
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	City        string `json:"city,omitempty"`
	District    string `json:"district,omitempty"`
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	Number      string `json:"number,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	Company     string `json:"company,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Fax         string `json:"fax,omitempty"`
}

type Product struct {
	Id          string  `json:"id,omitempty"`
	Category    string  `json:"category,omitempty"`
	Quantity    float64 `json:"quantity"`
	Code        string  `json:"code,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
}

type ChargeType string

const (
	ChargeType_DirectSale ChargeType = "directSale"
	ChargeType_Provision  ChargeType = "provision"
)

type CreatePaymentRequest struct {
	Token           string            `json:"token,omitempty"`
	Amount          float64           `json:"amount"`
	Installment     int               `json:"installment"`
	ReferenceCode   string            `json:"referenceCode,omitempty"`
	UseStoredCard   bool              `json:"useStoredCard"`
	Card            *CreditCard       `json:"card"`
	Reward          *Reward           `json:"reward"`
	Invoice         *Address          `json:"invoice"`
	Shipping        *Address          `json:"shipping"`
	Explanation     string            `json:"explanation,omitempty"`
	Use3d           bool              `json:"use3d"`
	AdditionalData  map[string]string `json:"additionalData"`
	Currency        string            `json:"currency,omitempty"`
	Email           string            `json:"email,omitempty"`
	Phone           string            `json:"phone,omitempty"`
	ReturnURL       string            `json:"returnURL,omitempty"`
	ChargeType      ChargeType        `json:"chargeType,omitempty"`
	PaymentSystemId string            `json:"paymentSystemId,omitempty"`
	NationalNumber  string            `json:"nationalNumber,omitempty"`
	Products        *[]Product        `json:"products"`
}

//Response

type Error struct {
	Summary string `json:"summary,omitempty"`
}

type Response struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}

type CardPaymentResponse struct {
	Response
	Behavior string `json:"behavior,omitempty"`
	Link     string `json:"link,omitempty"`
}
