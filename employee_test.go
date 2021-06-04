package neogovapi

import (
	"context"
	"os"
	"testing"
)

//All test require auth. Set ENV Variables or edit here

//All Test Values are defined here
var username = os.Getenv("ngUsername")
var password = os.Getenv("ngPassword")
var startDate = "1/1/2018"
var endDate = "12/31/2021"
var empID = "111111"
var empNumber = 1111
var departmentID = "111111"
var evaluationID = "111111"

func TestQueryEmployees(t *testing.T) {
	ngClient := NewClient(username, password)

	empData, err := ngClient.QueryEmployees(context.Background(), 10, 1)
	if err != nil {
		t.Error(err)
	}

	if len(empData) == 0 {
		t.Error("Error occurred in Query Employees. query data is empty")
	}
}

func TestQueryEmployeesByCreatedDate(t *testing.T) {
	ngClient := NewClient(username, password)

	empData, err := ngClient.QueryEmployeesByCreatedDate(context.Background(), startDate, endDate)
	if err != nil {
		t.Error(err)
	}

	if len(empData) == 0 {
		t.Error("Error occurred in Query Employees by created date. Query data is empty")
	}
}

func TestQueryEmployeesByUpdatedDate(t *testing.T) {
	ngClient := NewClient(username, password)

	empData, err := ngClient.QueryEmployeesByUpdatedDate(context.Background(), startDate, endDate)
	if err != nil {
		t.Error(err)
	}

	if len(empData) == 0 {
		t.Error("Error occurred in Query Employees by updated date. query data is empty")
	}
}

func TestQueryEmployeeByID(t *testing.T) {
	ngClient := NewClient(username, password)

	_, err := ngClient.QueryEmployeeByID(context.Background(), empID)
	if err != nil {
		t.Error(err)
	}
}

func TestQueryEmployeeByNumber(t *testing.T) {
	ngClient := NewClient(username, password)

	_, err := ngClient.QueryEmployeeByNumber(context.Background(), empNumber)
	if err != nil {
		t.Error(err)
	}
}

func TestQueryEmployeeEvaluations(t *testing.T) {
	ngClient := NewClient(username, password)

	evalData, err := ngClient.QueryEmployeeEvaluations(context.Background(), empID)
	if err != nil {
		t.Error(err)
	}

	if len(evalData) == 0 {
		t.Error("Error occurred in Query Employees Evaluations. query data is empty")
	}
}
