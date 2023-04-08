package upsdk

import (
	"fmt"
)

type AccountsService struct {
	client *Client
}

type AccountType string
type OwnershipType string

const (
	TransactionAccount AccountType = "TRANSACTIONAL"
	SaverAccount       AccountType = "SAVER"
)

const (
	JointOwnership      OwnershipType = "JOINT"
	IndividualOwnership OwnershipType = "INDIVIDUAL"
)

type AccountAttributes struct {
	DisplayName string        `json:"displayName"`
	Type        AccountType   `json:"accountType"`
	Ownership   OwnershipType `json:"ownershipType"`
	Balance     Balance       `json:"balance"`
	CreatedAt   string        `json:"createdAt"`
}

type Balance struct {
	CurrencyCode     string `json:"currencyCode"`
	Value            string `json:"value"`
	ValueInBaseUnits int    `json:"valueInBaseUnits"`
}

type Account struct {
	ID         string            `json:"id"`
	Attributes AccountAttributes `json:"attributes"`
}

func (as *AccountsService) List() ([]Account, *Response, error) {
	url := "accounts"

	req, err := as.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	accounts := new([]Account)
	res, err := as.client.Do(req, &ResponseMessage{Body: accounts})
	if err != nil {
		return nil, nil, err
	}

	return *accounts, res, nil
}

type AccountRequest struct {
	ID string
}

func (as *AccountsService) Get(input *AccountRequest) (*Account, *Response, error) {
	url := fmt.Sprintf("accounts/%s", input.ID)

	req, err := as.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	account := new(Account)
	res, err := as.client.Do(req, &ResponseMessage{Body: account})
	if err != nil {
		return nil, nil, err
	}

	return account, res, nil
}
