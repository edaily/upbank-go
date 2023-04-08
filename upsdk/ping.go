package upsdk

import (
	"fmt"
)

type UtilsService struct {
	client *Client
}

type Meta struct {
	ID          string `json:"id"`
	StatusEmoji string `json:"statusEmoji"`
}

type PingResponse struct {
	Meta Meta `json:"data"`
}

func (us *UtilsService) Ping() (*Meta, *Response, error) {
	req, err := us.client.NewRequest("GET", "util/ping", nil)
	if err != nil {
		return nil, nil, err
	}

	meta := new(Meta)
	res, err := us.client.Do(req, meta)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println(res)

	return meta, res, nil
}
