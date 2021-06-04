package neogovapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) QueryEmployees(ctx context.Context, perpage int, pagenumber int) ([]EmployeeQueryResponse, error) {
	qme, err := c.queryMultipleEmployees(ctx, fmt.Sprintf("https://api.neogov.com/rest/employees?perpage=%v&pagenumber=%v", perpage, pagenumber))
	if err != nil {
		return nil, err
	}

	return qme, nil
}

func (c *Client) QueryEmployeesByCreatedDate(ctx context.Context, startDate string, endDate string) ([]EmployeeQueryResponse, error) {
	params := url.Values{}
	params.Add("from", startDate)
	params.Add("to", endDate)

	qme, err := c.queryMultipleEmployees(ctx, fmt.Sprintf("https://api.neogov.com/rest/employees/created?%v", params.Encode()))
	if err != nil {
		return nil, err
	}

	return qme, nil
}

func (c *Client) QueryEmployeesByUpdatedDate(ctx context.Context, startDate string, endDate string) ([]EmployeeQueryResponse, error) {
	params := url.Values{}
	params.Add("from", startDate)
	params.Add("to", endDate)

	qme, err := c.queryMultipleEmployees(ctx, fmt.Sprintf("https://api.neogov.com/rest/employees/updated?%v", params.Encode()))
	if err != nil {
		return nil, err
	}

	return qme, nil
}

func (c *Client) QueryEmployeeByID(ctx context.Context, employeeID string) (*EmployeeQueryResponse, error) {
	qse, err := c.querySingleEmployee(ctx, fmt.Sprintf("https://api.neogov.com/rest/employees/%v", employeeID))
	if err != nil {
		return nil, err
	}

	return qse, nil
}

func (c *Client) QueryEmployeeByNumber(ctx context.Context, employeenumber int) (*EmployeeQueryResponse, error) {
	qse, err := c.querySingleEmployee(ctx, fmt.Sprintf("https://api.neogov.com/rest/employees/employeenumber/%v", employeenumber))
	if err != nil {
		return nil, err
	}

	return qse, nil
}

func (c *Client) QueryEmployeeEvaluations(ctx context.Context, id string) ([]Evaluation, error) {
	evals := []Evaluation{}

	r, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.neogov.com/rest/employees/%v/evaluations", id), nil)
	if err != nil {
		return evals, err
	}

	err = c.Send(r, &evals)
	if err != nil {
		return evals, err
	}

	return evals, nil
}

func (c *Client) queryMultipleEmployees(ctx context.Context, url string) ([]EmployeeQueryResponse, error) {
	eqr := []EmployeeQueryResponse{}

	r, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return eqr, err
	}

	err = c.Send(r, &eqr)
	if err != nil {
		return eqr, err
	}

	return eqr, nil
}

func (c *Client) querySingleEmployee(ctx context.Context, url string) (*EmployeeQueryResponse, error) {
	eqr := &EmployeeQueryResponse{}

	r, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return eqr, err
	}

	err = c.Send(r, eqr)
	if err != nil {
		return eqr, err
	}

	return eqr, nil
}
