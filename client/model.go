package klogs

import "time"

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
	Token             string            `json:"token,omitempty"`
	Amount            float64           `json:"amount"`
	Installment       int               `json:"installment"`
	ReferenceCode     string            `json:"referenceCode,omitempty"`
	UseStoredCard     bool              `json:"useStoredCard"`
	Card              *CreditCard       `json:"card"`
	OwnerKey          string            `json:"ownerKey,omitempty"`
	CardId            string            `json:"cardId,omitempty"`
	SaveCard          bool              `json:"saveCard"`
	Reward            *Reward           `json:"reward"`
	Invoice           *Address          `json:"invoice"`
	Shipping          *Address          `json:"shipping"`
	Explanation       string            `json:"explanation,omitempty"`
	Use3d             bool              `json:"use3d"`
	AdditionalData    map[string]string `json:"additionalData"`
	Currency          string            `json:"currency,omitempty"`
	Email             string            `json:"email,omitempty"`
	Phone             string            `json:"phone,omitempty"`
	ReturnURL         string            `json:"returnURL,omitempty"`
	ChargeType        ChargeType        `json:"chargeType,omitempty"`
	PaymentSystemCode string            `json:"paymentSystemCode,omitempty"`
	NationalNumber    string            `json:"nationalNumber,omitempty"`
	Products          *[]Product        `json:"products"`
}

type ProvisionCommitRequest struct {
	ReferenceCode string  `json:"referenceCode,omitempty"`
	Amount        float64 `json:"amount"`
}

type HostedPaymentRequest struct {
	Amount         float64           `json:"amount"`
	Currency       string            `json:"currency,omitempty"`
	ReferenceCode  string            `json:"referenceCode,omitempty"`
	FullName       string            `json:"fullName,omitempty"`
	NationalNumber string            `json:"nationalNumber,omitempty"`
	Email          string            `json:"email,omitempty"`
	Phone          string            `json:"phone,omitempty"`
	ReturnURL      string            `json:"returnURL,omitempty"`
	Invoice        *Address          `json:"invoice"`
	Shipping       *Address          `json:"shipping"`
	Explanation    string            `json:"explanation,omitempty"`
	AdditionalData map[string]string `json:"additionalData"`
	ChargeType     ChargeType        `json:"chargeType,omitempty"`
	Products       *[]Product        `json:"products"`
	PaymentMethod  string            `json:"paymentMethod,omitempty"`
}

type RefundRequest struct {
	ReferenceCode string  `json:"referenceCode,omitempty"`
	Amount        float64 `json:"amount"`
}

type VoidRequest struct {
	ReferenceCode string `json:"referenceCode,omitempty"`
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

type HostedPaymentResponse struct {
	Response
	PaymentId string `json:"paymentId,omitempty"`
	Link      string `json:"link,omitempty"`
}

type PaymentTokenResponse struct {
	Response
	Token string `json:"token,omitempty"`
}

type PaymentOptionQueryParams struct {
	Amount               float64
	BinNumber            string
	Currency             string
	CardId               string
	ProductCodes         []string
	ProductCategoryCodes []string
}

const (
	TransactionType_Sale            string = "sale"
	TransactionType_Provision       string = "provision"
	TransactionType_Void            string = "void"
	TransactionType_Refund          string = "refund"
	TransactionType_ProvisionCommit string = "provisionCommit"
)

const (
	TransactionStatus_Success string = "success"
	TransactionStatus_Failed  string = "failed"
	TransactionStatus_Pending string = "pending"
	TransactionStatus_Unknown string = "unknown"
)

const (
	PaymentSystemType_VirtualPos         string = "virtualPos"
	PaymentSystemType_PaymentFacilitator string = "paymentFacilitator"
	PaymentSystemType_ShoppingLoan       string = "shoppingLoan"
	PaymentSystemType_DigitalPay         string = "digitalPay"
	PaymentSystemType_Other              string = "other"
)

type PaymentTransaction struct {
	Id                  string          `json:"id,omitempty"`
	Balance             float64         `json:"balance,omitempty"`
	ClientReferenceCode string          `json:"clientReferenceCode,omitempty"`
	CommissionRate      float64         `json:"commissionRate,omitempty"`
	CurrencyCode        string          `json:"currencyCode,omitempty"`
	PaymentProviderId   string          `json:"paymentProviderId,omitempty"`
	RewardAmount        float64         `json:"rewardAmount,omitempty"`
	PlusInstallment     int32           `json:"plusInstallment,omitempty"`
	PaymentDeferral     int32           `json:"paymentDeferral,omitempty"`
	Installment         int32           `json:"installment,omitempty"`
	OrderId             string          `json:"orderId,omitempty"`
	PaymentMethod       string          `json:"paymentMethod,omitempty"`
	Amount              float64         `json:"amount"`
	Type                string          `json:"type,omitempty"`
	Status              string          `json:"status,omitempty"`
	CreatedAtUtc        time.Time       `json:"createdAtUtc,omitempty"`
	PaymentChannel      NameCode        `json:"paymentChannel,omitempty"`
	User                TransactionUser `json:"user,omitempty"`
	PaymentSystem       NameCode        `json:"paymentSystem,omitempty"`
	PaymentItems        []NameValue     `json:"paymentItems,omitempty"`
	Fields              []NameValue     `json:"fields,omitempty"`
}

type PaymentTransactionListItem struct {
	Amount       float64            `json:"amount"`
	Type         string             `json:"type,omitempty"`
	Status       string             `json:"status,omitempty"`
	CreatedAtUtc time.Time          `json:"createdAtUtc,omitempty"`
	Errors       []TransactionError `json:"errors,omitempty"`
}

type PaymentTransactionDetail struct {
	RelatedTransactions []PaymentTransactionListItem `json:"relatedTransactions,omitempty"`
	PaymentProvider     PaymentProvider              `json:"paymentProvider,omitempty"`
}

type PaymentSupportedSystem struct {
	Title   string `json:"title,omitempty"`
	Name    string `json:"name,omitempty"`
	Default bool   `json:"default,omitempty"`
}

type PaymentProvider struct {
	Id                      string                   `json:"id,omitempty"`
	Title                   string                   `json:"title,omitempty"`
	Version                 string                   `json:"version,omitempty"`
	IssuerCode              string                   `json:"issuerCode,omitempty"`
	SupportedSystems        []PaymentSupportedSystem `json:"supportedSystems,omitempty"`
	SystemType              string                   `json:"systemType,omitempty"`
	SupportedCurrencies     []string                 `json:"supportedCurrencies,omitempty"`
	SupportedPaymentMethods []string                 `json:"supportedPaymentMethods,omitempty"`
}

type NameCode struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type NameValue struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TransactionUser struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type TransactionError struct {
	ErrorCode    string    `json:"error,omitempty"`
	CreatedAtUtc time.Time `json:"createdAtUtc,omitempty"`
}

type TransactionDetailResponse struct {
	Response
	Transaction PaymentTransactionDetail `json:"transaction,omitempty"`
}

type PagedList[T any] struct {
	List        []T  `json:"list,omitempty"`
	CurrentPage int  `json:"currentPage,omitempty"`
	PageCount   int  `json:"pageCount,omitempty"`
	TotalCount  int  `json:"totalCount,omitempty"`
	HasNext     bool `json:"hasNext,omitempty"`
	HasPrevious bool `json:"hasPrevious,omitempty"`
}
