package allergies

const FirstNonExistingAllergyIndex uint = 256

/*
eggs (1)
peanuts (2)
shellfish (4)
strawberries (8)
tomatoes (16)
chocolate (32)
pollen (64)
cats (128)
*/
func Allergies(allergies uint) []string {
	allergiesTable := map[uint]string{
		1:   "eggs",
		2:   "peanuts",
		4:   "shellfish",
		8:   "strawberries",
		16:  "tomatoes",
		32:  "chocolate",
		64:  "pollen",
		128: "cats",
	}

	allergyScores := []uint{128, 64, 32, 16, 8, 4, 2, 1}
	var allergicTo []string

	allergies %= FirstNonExistingAllergyIndex
	for _, value := range allergyScores {
		if value > allergies {
			continue
		}

		allergicTo = append(allergicTo, allergiesTable[value])
		allergies -= value
	}

	return allergicTo
}

func AllergicTo(allergies uint, allergen string) bool {
	allergicTo := Allergies(allergies)
	return contains(allergicTo, allergen)
}

func contains(sliceToCheck []string, valueToCheck string) bool {
	for _, value := range sliceToCheck {
		if value == valueToCheck {
			return true
		}
	}

	return false
}
