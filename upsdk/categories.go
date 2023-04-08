package upsdk

type CategoriesService struct {
	client *Client
}

type CategoryAttributes struct {
	Name string `json:"name"`
}

type Category struct {
	ID         string             `json:"id"`
	Attributes CategoryAttributes `json:"attributes"`
}

type CategoriesResponse struct {
	Categories []Category      `json:"data"`
	Pagination PaginationLinks `json:"links"`
}

func (cs *CategoriesService) List() ([]Category, *Response, error) {
	req, err := cs.client.NewRequest("GET", "categories", nil)
	if err != nil {
		return nil, nil, err
	}

	categories := new([]Category)
	res, err := cs.client.Do(req, &ResponseMessage{Body: categories})
	if err != nil {
		return nil, nil, err
	}

	return *categories, res, nil
}
