package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavrylenkoIvan/gonotes"
	"github.com/gavrylenkoIvan/gonotes/pkg/service"
	mock_service "github.com/gavrylenkoIvan/gonotes/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_SignUp(t *testing.T) {
	type mockBehavior = func(*mock_service.MockAuthorization, gonotes.SignUpInput, string)

	testTable := []struct {
		name                 string
		input                gonotes.SignUpInput
		tempCode             string
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			input: gonotes.SignUpInput{
				Email:    "test@gmail.com",
				Password: "secret123456789",
				Username: "test123",
			},
			tempCode:  "WodTB2rJ8SobMgQ1",
			inputBody: `{"email":"test@gmail.com","password":"secret123456789","username":"test123"}`,
			mockBehavior: func(ma *mock_service.MockAuthorization, in gonotes.SignUpInput, tempCode string) {
				ma.EXPECT().CreateUser(in, tempCode).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"ok":true}`,
		},
		{
			name: "Invalid password",
			input: gonotes.SignUpInput{
				Email:    "test@gmail.com",
				Password: "secret123",
				Username: "test123",
			},
			tempCode:  "WodTB2rJ8SobMgQ1",
			inputBody: `{"email":"test@gmail.com","password":"secret123","username":"test123"}`,
			mockBehavior: func(ma *mock_service.MockAuthorization, in gonotes.SignUpInput, tempCode string) {
				ma.EXPECT().CreateUser(in, tempCode).Return(1, nil).AnyTimes()
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"ok":false,"message":"invalid input body"}`,
		},
		{
			name: "Invalid email",
			input: gonotes.SignUpInput{
				Email:    "test",
				Password: "secret12345678",
				Username: "test123",
			},
			tempCode:  "WodTB2rJ8SobMgQ1",
			inputBody: `{"email":"test","password":"secret12345678","username":"test123"}`,
			mockBehavior: func(ma *mock_service.MockAuthorization, in gonotes.SignUpInput, tempCode string) {
				ma.EXPECT().CreateUser(in, tempCode).Return(1, nil).AnyTimes()
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"ok":false,"message":"invalid input body"}`,
		},
		{
			name: "Service returned error",
			input: gonotes.SignUpInput{
				Email:    "test@gmail.com",
				Password: "secret123456789",
				Username: "test123",
			},
			tempCode:  "nrtR245jxOrsovFi",
			inputBody: `{"email":"test@gmail.com","password":"secret123456789","username":"test123"}`,
			mockBehavior: func(ma *mock_service.MockAuthorization, in gonotes.SignUpInput, tempCode string) {
				ma.EXPECT().CreateUser(in, tempCode).Return(0, errors.New("test_error")).AnyTimes()
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"ok":false,"message":"test_error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.input, testCase.tempCode)

			services := &service.Service{Authorization: auth}

			handler := NewHandler(services)

			r := gin.Default()

			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
