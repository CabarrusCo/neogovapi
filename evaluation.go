package neogovapi

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) QueryEvaulations(ctx context.Context, perpage int, pagenumber int) ([]Evaluation, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.neogov.com/rest/evaluations?perpage=%v&pagenumber=%v", perpage, pagenumber), nil)
	if err != nil {
		return nil, err
	}

	evals := []Evaluation{}

	err = c.Send(r, &evals)
	if err != nil {
		return nil, err
	}

	return evals, nil
}

func (c *Client) QueryEvaluationByID(ctx context.Context, id string) (*Evaluation, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.neogov.com/rest/evaluations/%v", id), nil)
	if err != nil {
		return nil, err
	}

	eval := &Evaluation{}

	err = c.Send(r, eval)
	if err != nil {
		return nil, err
	}

	return eval, nil
}
