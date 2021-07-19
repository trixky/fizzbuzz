package logic

import (
	"testing"

	"fizzbuzz.com/v1/models"
)

func TestFizzbuzz_request_statistics_generator(t *testing.T) {
	expected_fizzbuzz_request_statistics := []models.Fizzbuzz_request_statistics{{Int1: 1, Int2: 2, Limit: 3, Str1: "fizz", Str2: "buzz"}}

	fizzbuzz_request_statistics := Fizzbuzz_request_statistics_generator([]models.Fizzbuzz_request_statistics{{Int1: 1, Int2: 2, Limit: 3, Str1: "fizz", Str2: "buzz"}})

	for index, element := range expected_fizzbuzz_request_statistics {
		fizzbuzz_request_statistics_requests_element := fizzbuzz_request_statistics.Requests[index]
		if element.Int1 != fizzbuzz_request_statistics_requests_element.Int1 {
			t.Errorf("fizzbuzz; want ...")
		} else if element.Int2 != fizzbuzz_request_statistics_requests_element.Int2 {
			t.Errorf("fizzbuzz; want ...")
		}
	}
}
