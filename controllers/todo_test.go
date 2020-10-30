package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllTodos(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	functionResponse := GetAllTodos

	if functionResponse == nil {
		t.Error()
	} else {
		t.Log(c.Get("status"))
	}
}

type JsonCreate struct {
	Title     string
	Body      string
	Completed string
}
type dataTest struct {
	Json    string
	isError bool
}

func TestCreateTodo(t *testing.T) {
	data := dataTest{Json: `{"Title":"Title","Body":"Body","Completed":"true"}`, isError: false}
	var jsonObj JsonCreate

	err := json.Unmarshal([]byte(data.Json), &jsonObj)

	if err != nil {
		if data.isError {
			t.Log("Error")
		} else {
			t.Error("Failure")
		}
	} else {
		if data.isError {
			t.Error("Failure")
		} else {
			t.Log("OK")
		}
	}
}
