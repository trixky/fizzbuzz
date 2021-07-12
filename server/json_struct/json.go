// Package json_struct declares all generic structures related to json.
package json_struct

// Data_error contains error description.
type Data_error struct {
	Error string `json:"error"`
}

// Data_errors contains multiple error description.
type Data_errors struct {
	Errors []string `json:"errors"`
}
