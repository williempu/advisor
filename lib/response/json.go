package response

import (
	"encoding/json"
	"net/http"
)

// JSONMessage is a struct to generate json response
type JSONMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JSONResponse is an empty struct to use response-generator function
type JSONResponse struct{}

// HttpStatus writes the result of msg into json
func (j *JSONResponse) HttpStatus(w http.ResponseWriter, code int, msg interface{}) {
	// Set header to json and write code to header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// try to marshal into json using the message struct
	jsonMsg, err := json.Marshal(JSONMessage{
		Code:    code,
		Message: http.StatusText(code),
		Data:    msg, // msg is an interface, meaning it could be anything
	})
	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.Write(jsonMsg)
}
