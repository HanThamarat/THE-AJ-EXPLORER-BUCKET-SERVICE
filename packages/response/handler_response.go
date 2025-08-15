package response

import (
	"github.com/gofiber/fiber/v2"
)

type (
	Response struct {
		Status  int    		`json:"status"`
        Message string 		`json:"message"`
        Body    interface{} `json:"body,omitempty"`
	}

	ErrorHandler struct {
		Status 	int 		`json:"status"`
		Message string 		`json:"message"`
		Error 	string 		`json:"error"`
	}
)

func SendResponseHandler(c *fiber.Ctx, status int, message string, body interface{}) error {
	response := Response{
		Status: status,
		Message: message,
		Body: body,
	}

	return c.Status(status).JSON(response);
}

func SendErrorHandler(c *fiber.Ctx, status int, message string, err string) error {
	error := ErrorHandler{
		Status: status,
		Message: message,
		Error: err,
	}

	return c.Status(status).JSON(error);
}