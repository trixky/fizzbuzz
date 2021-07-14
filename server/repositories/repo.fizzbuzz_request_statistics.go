package repositories

import (
	"errors"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/models"
	"gorm.io/gorm"
)

func Getall_fizzbuzz_request_statistics() ([]models.Fizzbuzz_request_statistics, error) {
	results := []models.Fizzbuzz_request_statistics{}

	if err := database.Postgres.Order("request_number DESC").Limit(10).Find(&results).Error; err != nil {
		// If postgres failed (to deliver the most popular fizzbuzz requests).
		return results, err
	}
	return results, nil
}

func Create_or_increment_fizzbuzz_request_statistics(int1 int, int2 int, limit int, str1 string, str2 string) (next_request_number int, err error) {

	current_fizzbuzz_request_statistic := models.Fizzbuzz_request_statistics{}

	if err := database.Postgres.Table("fizzbuzz_request_statistics").Where(map[string]interface{}{"int1": int1, "int2": int2, "_limit": limit, "str1": str1, "str2": str2}).First(&current_fizzbuzz_request_statistic).Error; err != nil {
		// If no fizzbuzz request was found in postgres with these inputs.
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// If postgres failed.
			return 0, err
		} else if err := database.Postgres.Create(&models.Fizzbuzz_request_statistics{Int1: int1, Int2: int2, Limit: limit, Str1: str1, Str2: str2}).Error; err != nil {
			// If the creation of the new fizzbuzz request in postgres failed.
			return 0, err
		} else {
			// If the creation of the new fizzbuzz request in a success.
			return 1, nil
		}
	} else if err := database.Postgres.Table("fizzbuzz_request_statistics").Where(map[string]interface{}{"int1": int1, "int2": int2, "_limit": limit, "str1": str1, "str2": str2}).UpdateColumn("request_number", current_fizzbuzz_request_statistic.Request_number+1).Error; err != nil {
		// Else if the update of the fizzbuzz request in postgres failed (increment 'request_number' column).
		return current_fizzbuzz_request_statistic.Request_number, err
	}

	return current_fizzbuzz_request_statistic.Request_number + 1, nil
}
