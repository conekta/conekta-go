package conekta

import (
	"encoding/json"
	"fmt"
)

// ErrorDetails gives especific information about an error
type ErrorDetails struct {
	DebugMessage string `json:"debug_message,omitempty"`
	Message      string `json:"message,omitempty"`
	Param        string `json:"param,omitempty"`
	Code         string `json:"code,omitempty"`
}

// Error describes Conekta errors
type Error struct {
	ErrorType string         `json:"type,omitempty"`
	LogID     string         `json:"log_id,omitempty"`
	Details   []ErrorDetails `json:"details,omitempty"`
}

func getConectionError() Error {
	return Error{
		ErrorType: "error.requestor.connection_purchaser",
		Details: []ErrorDetails{
			{
				Message: "error.requestor.connection",
			},
		},
	}
}

func getAPIError(b []byte) Error {
	e := Error{}
	json.Unmarshal(b, &e)
	return e
}

func (e Error) Error() string {
	return fmt.Sprintf("%v: %v", e.ErrorType, e.Details)
}
