package rest

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"toDoApp/pkg/core"
)

func TestValidateUpdate(t *testing.T) {
	type test struct {
		body               string
		request            core.PostToDoRequest
		validationFunction func(request core.PostToDoRequest) (valid bool, errors []string)
		wantedSuccess      bool
		wantedErrors       []string
		wantedRequest      core.PostToDoRequest
	}

	tests := []test{
		{
			body:               "{\"UserName\": \"Dirk\", \"ToDos\": [\"Do something\", \"Do something else\"]}",
			request:            core.PostToDoRequest{},
			validationFunction: core.ValidatePostToDoRequest,
			wantedSuccess:      true,
			wantedErrors:       []string{},
			wantedRequest:      core.PostToDoRequest{UserName: "Dirk", ToDos: []string{"Do something", "Do something else"}}},
		{
			body:               "{\"UserName\": \"Dirk\"}",
			request:            core.PostToDoRequest{},
			validationFunction: core.ValidatePostToDoRequest,
			wantedSuccess:      false,
			wantedErrors:       []string{"At least one todo must be provided"},
			wantedRequest:      core.PostToDoRequest{UserName: "Dirk"}},
		{
			body:               "{\"UserName\": \"\"}",
			request:            core.PostToDoRequest{},
			validationFunction: core.ValidatePostToDoRequest,
			wantedSuccess:      false,
			wantedErrors:       []string{"A user name must be provided", "At least one todo must be provided"},
			wantedRequest:      core.PostToDoRequest{UserName: ""}},
	}

	for _, tc := range tests {
		reader := strings.NewReader(tc.body)
		httpRequest := httptest.NewRequest(http.MethodGet, "localhost:808/todo", reader)
		success, errors, request := convertToRequestType(tc.request, httpRequest, tc.validationFunction)

		if success != tc.wantedSuccess {
			t.Errorf("actual %t expected %t are not equal", success, tc.wantedSuccess)
		}

		if !reflect.DeepEqual(errors, tc.wantedErrors) {
			t.Errorf("actual %q expected %q are not equal", errors, tc.wantedErrors)
		}

		if !reflect.DeepEqual(request, tc.wantedRequest) {
			t.Errorf("actual %q expected %q are not equal", request, tc.wantedRequest)
		}
	}

}
