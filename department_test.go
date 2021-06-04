package neogovapi

import (
	"context"
	"testing"
)

func TestQueryDepartments(t *testing.T) {
	ngClient := NewClient(username, password)

	qdr, err := ngClient.QueryDepartments(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(qdr) == 0 {
		t.Error("Length of departments returned is 0")
	}
}

func TestQueryDepartmentByID(t *testing.T) {
	ngClient := NewClient(username, password)

	_, err := ngClient.QueryDepartmentByID(context.Background(), departmentID)
	if err != nil {
		t.Error(err)
	}
}
