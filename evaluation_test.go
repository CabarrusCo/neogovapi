package neogovapi

import (
	"context"
	"testing"
)

func TestQueryEvaluations(t *testing.T) {
	ngClient := NewClient(username, password)

	evals, err := ngClient.QueryEvaluations(context.Background(), 10, 1)
	if err != nil {
		t.Error(err)
	}

	if len(evals) == 0 {
		t.Error("Length of Evaluations returned is 0")
	}
}

func TestQueryEvaluationByID(t *testing.T) {
	ngClient := NewClient(username, password)

	_, err := ngClient.QueryEvaluationByID(context.Background(), evaluationID)
	if err != nil {
		t.Error(err)
	}
}
