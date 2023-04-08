package upsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type TagsService struct {
	client *Client
}

type Tag struct {
	ID   string `json:"id"`
	Type string `json:"tags"`
}

func (ts *TagsService) List() ([]Tag, *Response, error) {
	req, err := ts.client.NewRequest("GET", "tags", nil)
	if err != nil {
		return nil, nil, err
	}

	tags := new([]Tag)
	res, err := ts.client.Do(req, tags)
	if err != nil {
		return nil, nil, err
	}

	return *tags, res, nil
}

type AddTagRequest struct {
	TransactionID string
	Tag           string
}

type AddResponse struct {
}

func (ts *TagsService) Add(input *AddTagRequest) (*Tag, *Response, error) {
	url := fmt.Sprintf("transactions/%s/relationships/tags", input.TransactionID)

	postBody, _ := json.Marshal(map[string]string{
		"type": "tags",
		"id":   input.Tag,
	})
	reqBody := bytes.NewBuffer(postBody)

	req, err := ts.client.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, nil, err
	}

	tag := new(Tag)
	res, err := ts.client.Do(req, tag)
	if err != nil {
		return nil, nil, err
	}

	return tag, res, nil
}
