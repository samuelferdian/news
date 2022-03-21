package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "add news",
			route:         "/api/news/add",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "Success insert news",
		},
	}

	app := fiber.New()

	app.Post("/api/news/add", func(c *fiber.Ctx) error {
		return c.SendString("Success insert news")
	})

	for _, test := range tests {
		req, _ := http.NewRequest(
			"POST",
			test.route,
			nil,
		)
		res, err := app.Test(req, 5)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, test.description)
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
