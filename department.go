package neogovapi

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) QueryDepartments(ctx context.Context) ([]DepartmentResponse, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", "https://api.neogov.com/rest/departments", nil)
	if err != nil {
		return nil, err
	}

	dr := []DepartmentResponse{}

	err = c.Send(r, &dr)
	if err != nil {
		return nil, err
	}

	return dr, nil
}

func (c *Client) QueryDepartmentByID(ctx context.Context, id string) (*DepartmentResponse, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.neogov.com/rest/departments/%v", id), nil)
	if err != nil {
		return nil, err
	}

	dr := &DepartmentResponse{}

	err = c.Send(r, dr)
	if err != nil {
		return nil, err
	}

	return dr, nil
}
