// Package routes gives all the handler functions corresponding to their routes and methods.
package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/json_struct"
	"fizzbuzz.com/v1/models"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

// Data_fizzbuzz contains the list of all the values of the fizzbuzz request.
type Data_fizzbuzz struct {
	Fizzbuzz []string `json:"fizzbuzz"`
}

// Fizzbuzz_errors lists the potential reading errors of the expected values.
type Fizzbuzz_errors struct {
	error_int1_conv  error
	error_int2_conv  error
	error_limit_conv error
}

// Fizzbuzz_values retrieves the values necessary for the generation of the fizzbuzz.
type Fizzbuzz_values struct {
	int1  int
	int2  int
	limit int
	str1  string
	str2  string
}

// check_int checks that the "int_x" values are not corrupted.
func check_int(int_x int, limit int) bool {
	return limit > 0 && (int_x < 0 || int_x > limit)
}

// check_limit checks that the "limit" value is not corrupted.
func check_limit(limit int) bool {
	return limit < 1 || limit > 100
}

// check_str checks that the "str_x" values are not corrupted.
func check_str(str_x string) bool {
	return len(str_x) < 1 || len(str_x) > 30
}

// final_error_generator generates a list of all possible errors concerning user inputs.
func final_error_generator(fizzbuzz_values *Fizzbuzz_values, fizzbuzz_errors *Fizzbuzz_errors) []string {
	errs := []string{}

	// check int1
	if fizzbuzz_errors.error_int1_conv != nil {
		// If there is a conversion error on int1.
		errs = append(errs, "int1 is missing or incorrectly formatted")
	} else if check_int(fizzbuzz_values.int1, fizzbuzz_values.limit) {
		// Else if int1 is invalid.
		errs = append(errs, "int1 must be between 0 and the limit")
	}
	// check int2
	if fizzbuzz_errors.error_int2_conv != nil {
		// If there is a conversion error on int2.
		errs = append(errs, "int2 is missing or incorrectly formatted")
	} else if check_int(fizzbuzz_values.int2, fizzbuzz_values.limit) {
		// Else if int2 is invalid.
		errs = append(errs, "int2 must be between 0 and the limit")
	}
	// check limit
	if fizzbuzz_errors.error_limit_conv != nil {
		// If there is a conversion error on limit.
		errs = append(errs, "limit is missing or incorrectly formatted")
	} else if check_limit(fizzbuzz_values.limit) {
		// Else if limit is invalid.
		errs = append(errs, "limit must be between 1 and 100")
	}
	// check str1
	if len(fizzbuzz_values.str1) < 1 {
		// If str1 is missing.
		errs = append(errs, "str1 is missing")
	} else if check_str(fizzbuzz_values.str1) {
		// Else if str1 is invalid.
		errs = append(errs, "str1 must be between 0 and 30 characters long")
	}
	// check str2
	if len(fizzbuzz_values.str2) < 1 {
		// If str2 is missing.
		errs = append(errs, "str2 is missing")
	} else if check_str(fizzbuzz_values.str2) {
		// Else if str2 is invalid.
		errs = append(errs, "str2 must be between 0 and 30 characters long")
	}

	return errs
}

// fizzbuzz_generator generates the fizzbuzz from user inputs.
func fizzbuzz_generator(fizzbuzz_values *Fizzbuzz_values) []string {
	fizzbuzz_array := make([]string, fizzbuzz_values.limit)

	for i := 0; i < fizzbuzz_values.limit; i++ {
		ii := i + 1
		fizzbuzzed := false

		if ii%fizzbuzz_values.int1 == 0 {
			// If it's a multiple of int1.
			fizzbuzzed = true
			fizzbuzz_array[i] = fizzbuzz_values.str1
		}
		if ii%fizzbuzz_values.int2 == 0 {
			// If it's a multiple of int2.
			if fizzbuzzed {
				// If it's a multiple of int1 and int2.
				fizzbuzz_array[i] += fizzbuzz_values.str2
			} else {
				// If it's only a multiple of int2.
				fizzbuzzed = true
				fizzbuzz_array[i] = fizzbuzz_values.str2
			}
		}
		if !fizzbuzzed {
			// If it's not multiple of int1 or int2.
			fizzbuzz_array[i] = strconv.Itoa(ii)
		}
	}

	return fizzbuzz_array
}

// Fizzbuzz handles the [GET /fizzbuzz] request.
func Fizzbuzz(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")

	fizzbuzz_values := Fizzbuzz_values{}
	fizzbuzz_errors := Fizzbuzz_errors{}

	fizzbuzz_values.int1, fizzbuzz_errors.error_int1_conv = strconv.Atoi(req.URL.Query().Get("int1"))    // get input int1
	fizzbuzz_values.int2, fizzbuzz_errors.error_int2_conv = strconv.Atoi(req.URL.Query().Get("int2"))    // get input int2
	fizzbuzz_values.limit, fizzbuzz_errors.error_limit_conv = strconv.Atoi(req.URL.Query().Get("limit")) // get input limit
	fizzbuzz_values.str1 = req.URL.Query().Get("str1")                                                   // get input str1
	fizzbuzz_values.str2 = req.URL.Query().Get("str2")                                                   // get input str2

	if errs := final_error_generator(&fizzbuzz_values, &fizzbuzz_errors); len(errs) > 0 {
		// If there is at least one error in the user inputs.
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(json_struct.Data_errors{Errors: errs})
		return
	} else {
		current_fizzbuzz_request_statistic := models.Fizzbuzz_request_statistics{}

		if err := database.Postgres.Table("fizzbuzz_request_statistics").Where(map[string]interface{}{"int1": fizzbuzz_values.int1, "int2": fizzbuzz_values.int2, "_limit": fizzbuzz_values.limit, "str1": fizzbuzz_values.str1, "str2": fizzbuzz_values.str2}).First(&current_fizzbuzz_request_statistic).Error; err != nil {
			// If no fizzbuzz request was found in postgres with these inputs.
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// If no fizzbuzz request is known in postgres with these inputs.
				if err := database.Postgres.Create(&models.Fizzbuzz_request_statistics{Int1: fizzbuzz_values.int1, Int2: fizzbuzz_values.int2, Limit: fizzbuzz_values.limit, Str1: fizzbuzz_values.str1, Str2: fizzbuzz_values.str2}).Error; err != nil {
					// If the creation of the new fizzbuzz request in postgres failed.
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
					return
				} else {
					data_fizzbuzz := Data_fizzbuzz{fizzbuzz_generator(&fizzbuzz_values)}
					json.NewEncoder(res).Encode(data_fizzbuzz)
				}
			} else {
				res.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
				return
			}
			return
		} else {
			if err := database.Postgres.Table("fizzbuzz_request_statistics").Where(map[string]interface{}{"int1": fizzbuzz_values.int1, "int2": fizzbuzz_values.int2, "_limit": fizzbuzz_values.limit, "str1": fizzbuzz_values.str1, "str2": fizzbuzz_values.str2}).UpdateColumn("request_number", current_fizzbuzz_request_statistic.Request_number+1).Error; err != nil {
				// If the update of the fizzbuzz request in postgres failed (increment 'request_number' column).
				res.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(res).Encode(json_struct.Data_error{Error: "internal server error"})
				return
			} else {
				data_fizzbuzz := Data_fizzbuzz{fizzbuzz_generator(&fizzbuzz_values)}
				json.NewEncoder(res).Encode(data_fizzbuzz)
			}
		}
	}
}
