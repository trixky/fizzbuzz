package logic

import (
	"testing"

	"fizzbuzz.com/v1/extractors"
)

// TestFizzbuzz_generator tests the Fizzbuzz_generator function.
func TestFizzbuzz_generator(t *testing.T) {
	// incomplete... just to try
	expected_fizzbuzz := Data_fizzbuzz{Fizzbuzz: []string{"1",
		"2",
		"fizz",
		"4",
		"buzz",
		"fizz",
		"7",
		"8",
		"fizz",
		"buzz",
		"11",
		"fizz",
		"13",
		"14",
		"fizzbuzz",
		"16",
		"17",
		"fizz",
		"19",
		"buzz",
		"fizz",
		"22",
		"23",
		"fizz",
		"buzz",
		"26",
		"fizz",
		"28",
		"29",
		"fizzbuzz",
		"31",
		"32",
		"fizz",
		"34",
		"buzz",
		"fizz",
		"37",
		"38",
		"fizz",
		"buzz",
		"41",
		"fizz",
		"43",
		"44",
		"fizzbuzz",
		"46",
		"47",
		"fizz",
		"49",
		"buzz",
		"fizz",
		"52",
		"53",
		"fizz",
		"buzz",
		"56",
		"fizz",
		"58",
		"59",
		"fizzbuzz",
		"61",
		"62",
		"fizz",
		"64",
		"buzz",
		"fizz",
		"67",
		"68",
		"fizz",
		"buzz",
		"71",
		"fizz",
		"73",
		"74",
		"fizzbuzz",
		"76",
		"77",
		"fizz",
		"79",
		"buzz",
		"fizz",
		"82",
		"83",
		"fizz",
		"buzz",
		"86",
		"fizz",
		"88",
		"89",
		"fizzbuzz",
		"91",
		"92",
		"fizz",
		"94",
		"buzz",
		"fizz",
		"97",
		"98",
		"fizz",
		"buzz"}}

	fizzbuzz := Fizzbuzz_generator(&extractors.Fizzbuzz{Int1: 3, Int2: 5, Limit: 100, Str1: "fizz", Str2: "buzz"})

	for index, element := range expected_fizzbuzz.Fizzbuzz {
		if element != fizzbuzz.Fizzbuzz[index] {
			t.Errorf("expected: %s, result: %s", element, fizzbuzz.Fizzbuzz[index])
		}
	}
}
