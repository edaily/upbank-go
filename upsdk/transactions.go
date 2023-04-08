package upsdk

type TransactionsService struct {
	client *Client
}

type TransactionStatus string

const (
	SettledTransaction TransactionStatus = "SETTLED"
)

type TransactionAmount struct {
	CurrencyCode     string `json:"currencyCode"`
	Value            string `json:"value"`
	ValueInBaseUnits int32  `json:"valueInBaseUnits"`
}

type TransactionHoldInfo struct {
	Amount TransactionAmount `json:"amount"`
}

type TransactionCardPurchaseMethod struct {
	Method           string `json:"method"`
	CardNumberSuffix string `json:"cardNumberSuffix"`
}

type TransactionAttributes struct {
	Status             TransactionStatus             `json:"name"`
	RawText            string                        `json:"rawText"`
	Description        string                        `json:"description"`
	IsCategorizable    bool                          `json:"isCategorizable"`
	Amount             TransactionAmount             `json:"amount"`
	CardPurchaseMethod TransactionCardPurchaseMethod `json:"cardPurchaseMethod"`
	SettledAt          string                        `json:"settledAt"`
	CreatedAt          string                        `json:"createdAt"`
}

type Transaction struct {
	ID         string                `json:"id"`
	Attributes TransactionAttributes `json:"attributes"`
}

func (ts *TransactionsService) List() ([]Transaction, *Response, error) {
	req, err := ts.client.NewRequest("GET", "transactions", nil)
	if err != nil {
		return nil, nil, err
	}

	transactions := new([]Transaction)
	res, err := ts.client.Do(req, &ResponseMessage{Body: transactions})
	if err != nil {
		return nil, nil, err
	}

	return *transactions, res, nil
}
