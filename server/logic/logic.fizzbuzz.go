package logic

import (
	"strconv"

	"fizzbuzz.com/v1/extractors"
)

// Data_fizzbuzz contains the list of all the values of the fizzbuzz request.
type Data_fizzbuzz struct {
	Fizzbuzz []string `json:"fizzbuzz"`
}

// fizzbuzz_generator generates the fizzbuzz from user inputs.
func Fizzbuzz_generator(extracted_fizzbuzz *extractors.Fizzbuzz) (data_fizzbuzz Data_fizzbuzz) {
	data_fizzbuzz.Fizzbuzz = make([]string, extracted_fizzbuzz.Limit)

	for i := 0; i < extracted_fizzbuzz.Limit; i++ {
		ii := i + 1

		if ii%extracted_fizzbuzz.Int1 == 0 {
			if ii%extracted_fizzbuzz.Int2 == 0 {
				// If it's a multiple of int1.
				data_fizzbuzz.Fizzbuzz[i] = extracted_fizzbuzz.Str1 + extracted_fizzbuzz.Str2
			} else {
				// If it's a multiple of int1 and int2.
				data_fizzbuzz.Fizzbuzz[i] = extracted_fizzbuzz.Str1
			}
		} else if ii%extracted_fizzbuzz.Int2 == 0 { // repeating code but more optimized
			// If it's a multiple of int2.
			data_fizzbuzz.Fizzbuzz[i] = extracted_fizzbuzz.Str2
		} else {
			// If it's not multiple of int1 or int2.
			data_fizzbuzz.Fizzbuzz[i] = strconv.Itoa(ii)
		}
	}
	return
}
