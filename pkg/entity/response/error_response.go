package response

import (
	"encoding/json"
	"io"
)

type ErrorResponseValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ErrorResponseData []ErrorResponseValue

//Method to encode from golang struct into JSON format
func (er *ErrorResponseValue) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(er)
}

func NewErrorResponseValue(key string, value string) ErrorResponseValue {
	return ErrorResponseValue{
		Key:   key,
		Value: value,
	}
}

func NewErrorResponseData(errorResponseValues ...ErrorResponseValue) ErrorResponseData {
	errors := ErrorResponseData{}

	for _, values := range errorResponseValues {
		errors = append(errors, values)
	}
	return errors
}
