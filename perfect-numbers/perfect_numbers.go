package perfect

import "fmt"

// Define the Classification type here.
type Classification int

const ClassificationAbundant Classification = 0
const ClassificationDeficient Classification = 1
const ClassificationPerfect Classification = 2
const ClassificationNegative Classification = -1

var ErrOnlyPositive = fmt.Errorf("cannot process negative number")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return ClassificationNegative, ErrOnlyPositive
	}

	multipliers := GetMultipliers(n)
	sum := GetSum(multipliers)

	if sum < n {
		return ClassificationDeficient, nil
	} else if sum == n {
		return ClassificationPerfect, nil
	} else /* sum > n*/ {
		return ClassificationAbundant, nil
	}
}

func GetSum(multipliers []int64) int64 {
	var sum int64 = 0
	for _, multiplier := range multipliers {
		sum += multiplier
	}

	return sum
}

func GetMultipliers(n int64) []int64 {

	var multipliers []int64
	for i := 1; int64(i) <= n/2; i++ {
		if n%int64(i) == 0 {
			multipliers = append(multipliers, int64(i))
		}
	}

	return multipliers
}
