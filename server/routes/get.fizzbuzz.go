package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fizzbuzz.com/v1/json_struct"
	"fizzbuzz.com/v1/middleware"
)

type Data_fizzbuzz struct {
	Fizzbuzz []string `json:"fizzbuzz"`
}

type Fizzbuzz_errors struct {
	error_int1_conv  error
	error_int2_conv  error
	error_limit_conv error
}

type Fizzbuzz_values struct {
	int1  int
	int2  int
	limit int
	str1  string
	str2  string
}

func check_int(intx int, limit int) bool {
	return limit > 0 && (intx < 0 || intx > limit)
}

func check_limit(limit int) bool {
	return limit < 1 || limit > 100
}

func check_str(strx string) bool {
	return len(strx) < 1 || len(strx) > 30
}

func final_error_generator(fizzbuzz_values *Fizzbuzz_values, fizzbuzz_errors *Fizzbuzz_errors) []string {
	errs := []string{}

	// ------------- int1
	if fizzbuzz_errors.error_int1_conv != nil {
		errs = append(errs, "int1 is missing or incorrectly formatted")
	} else if check_int(fizzbuzz_values.int1, fizzbuzz_values.limit) {
		errs = append(errs, "int1 must be between 0 and the limit")
	}
	// ------------- int2
	if fizzbuzz_errors.error_int2_conv != nil {
		errs = append(errs, "int2 is missing or incorrectly formatted")
	} else if check_int(fizzbuzz_values.int2, fizzbuzz_values.limit) {
		errs = append(errs, "int2 must be between 0 and the limit")
	}
	// ------------- limit
	if fizzbuzz_errors.error_limit_conv != nil {
		errs = append(errs, "limit is missing or incorrectly formatted")
	} else if check_limit(fizzbuzz_values.limit) {
		errs = append(errs, "limit must be between 1 and 100")
	}
	// ------------- str1
	if len(fizzbuzz_values.str1) < 1 {
		errs = append(errs, "str1 is missing")
	} else if check_str(fizzbuzz_values.str1) {
		errs = append(errs, "str1 must be between 0 and 30 characters long")
	}
	// ------------- str2
	if len(fizzbuzz_values.str2) < 1 {
		errs = append(errs, "str2 is missing")
	} else if check_str(fizzbuzz_values.str2) {
		errs = append(errs, "str2 must be between 0 and 30 characters long")
	}

	return errs
}

func fizzbuzz_generator(fizzbuzz_values *Fizzbuzz_values) []string {
	fizzbuzz_array := make([]string, fizzbuzz_values.limit)

	for i := 0; i < fizzbuzz_values.limit; i++ {
		ii := i + 1
		fizzbuzzed := false

		if ii%fizzbuzz_values.int1 == 0 {
			fizzbuzzed = true
			fizzbuzz_array[i] = fizzbuzz_values.str1
		}
		if ii%fizzbuzz_values.int2 == 0 {
			if fizzbuzzed {
				fizzbuzz_array[i] += fizzbuzz_values.str2
			} else {
				fizzbuzzed = true
				fizzbuzz_array[i] = fizzbuzz_values.str2
			}
		}
		if !fizzbuzzed {
			fizzbuzz_array[i] = strconv.Itoa(ii)
		}
	}

	return fizzbuzz_array
}

func Fizzbuzz(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/fizzbuzz" || req.Method != "GET" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	_, authentified := middleware.Middleware_token(res, req)
	if !authentified {
		return
	}

	fizzbuzz_values := Fizzbuzz_values{}
	fizzbuzz_errors := Fizzbuzz_errors{}

	fizzbuzz_values.int1, fizzbuzz_errors.error_int1_conv = strconv.Atoi(req.URL.Query().Get("int1"))
	fizzbuzz_values.int2, fizzbuzz_errors.error_int2_conv = strconv.Atoi(req.URL.Query().Get("int2"))
	fizzbuzz_values.limit, fizzbuzz_errors.error_limit_conv = strconv.Atoi(req.URL.Query().Get("limit"))
	fizzbuzz_values.str1 = req.URL.Query().Get("str1")
	fizzbuzz_values.str2 = req.URL.Query().Get("str2")

	errs := final_error_generator(&fizzbuzz_values, &fizzbuzz_errors)
	if len(errs) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_errors{Errors: errs})
	} else {
		json.NewEncoder(res).Encode(Data_fizzbuzz{fizzbuzz_generator(&fizzbuzz_values)})
	}
}
