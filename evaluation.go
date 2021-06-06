package neogovapi

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) QueryEvaluations(ctx context.Context, perpage int, pagenumber int) ([]Evaluation, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.neogov.com/rest/evaluations?perpage=%v&pagenumber=%v", perpage, pagenumber), nil)
	if err != nil {
		return nil, err
	}

	e := []Evaluation{}

	err = c.Send(r, &e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (c *Client) QueryEvaluationByID(ctx context.Context, id string) (Evaluation, error) {
	e := Evaluation{}

	r, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.neogov.com/rest/evaluations/%v", id), nil)
	if err != nil {
		return e, err
	}

	err = c.Send(r, &e)
	if err != nil {
		return e, err
	}

	return e, nil
}
