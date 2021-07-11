package json_struct

type Data_error struct {
	Error string `json:"error"`
}

type Data_errors struct {
	Errors []string `json:"errors"`
}
