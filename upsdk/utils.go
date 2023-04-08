package upsdk

type UtilsService struct {
	client *Client
}

type Ping struct {
	ID          string `json:"id"`
	StatusEmoji string `json:"statusEmoji"`
}

type PingResponse struct {
	Ping Ping `json:"meta"`
}

func (us *UtilsService) Ping() (*Ping, *Response, error) {
	req, err := us.client.NewRequest("GET", "util/ping", nil)
	if err != nil {
		return nil, nil, err
	}

	meta := new(PingResponse)
	res, err := us.client.Do(req, meta)
	if err != nil {
		return nil, res, err
	}

	return &meta.Ping, res, nil
}
