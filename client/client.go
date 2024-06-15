package client

import (
	"context"
	"fmt"

	"github.com/aftab-hussain-93/empl/internal/service"
	fault "github.com/aftab-hussain-93/empl/pkg/err"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	*resty.Client
}

func New() *Client {
	return &Client{resty.New()}
}

func (c *Client) CreateEmployee(ctx context.Context, data map[string]any) (int, any, error) {
	res := &service.Employee{}
	errRes := &fault.HTTPError{}

	restyResp, err := c.Client.R().
		SetContext(ctx).
		SetBody(data).
		SetResult(res).
		SetError(errRes).
		Post("employees")
	if err != nil {
		return 0, nil, err
	}
	if restyResp.IsError() {
		return restyResp.StatusCode(), errRes, nil
	}

	return restyResp.StatusCode(), res, nil
}

type ListEmpRes struct {
	Total     int                 `json:"total"`
	Count     int                 `json:"count"`
	Employees []*service.Employee `json:"employees"`
}

func (c *Client) GetEmployees(ctx context.Context, pageSize, pageNumber int) (int, any, error) {

	res := &ListEmpRes{}
	errRes := &fault.HTTPError{}
	restyResp, err := c.Client.R().SetContext(ctx).
		SetQueryParams(map[string]string{
			"page": fmt.Sprint(pageNumber),
			"size": fmt.Sprint(pageSize),
		}).
		SetResult(res).
		SetError(errRes).
		Get("employees")
	if err != nil {
		return 0, nil, err
	}
	if restyResp.IsError() {
		return restyResp.StatusCode(), errRes, nil
	}
	return restyResp.StatusCode(), res, nil
}
