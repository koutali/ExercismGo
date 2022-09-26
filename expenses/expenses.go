package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var filtered []Record
	for _, record := range in {
		if predicate(record) {
			filtered = append(filtered, record)
			fmt.Println(record)
		}
	}

	return filtered
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return record.Day >= p.From && record.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	totalCost := 0.0
	byDaysFunc := ByDaysPeriod(p)
	for _, record := range in {
		if byDaysFunc(record) {
			totalCost += record.Amount
		}
	}
	return totalCost
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	totalCost := 0.0
	byCategoryFunc := ByCategory(c)
	filtered := Filter(in, byCategoryFunc)
	if len(filtered) == 0 {
		errorMessage := fmt.Sprintf("%s%s", "unknown category ", c)
		return totalCost, errors.New(errorMessage)

	}
	byPeriodFunc := ByDaysPeriod(p)
	for _, record := range filtered {
		if byPeriodFunc(record) {
			totalCost += record.Amount
		}
	}

	return totalCost, nil
}
