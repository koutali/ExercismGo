package twelve

func Verse(i int) string {
	dayIndices := map[int]string{
		1:  "first",
		2:  "second",
		3:  "third",
		4:  "fourth",
		5:  "fifth",
		6:  "sixth",
		7:  "seventh",
		8:  "eighth",
		9:  "ninth",
		10: "tenth",
		11: "eleventh",
		12: "twelfth",
	}

	const start = "On the"
	const blank = " "
	const secondPart = "day of Christmas my true love gave to me:"
	const dot = "."
	verse := start + blank + dayIndices[i] + blank + secondPart + blank + generateLyricForDay(i) + dot

	return verse
}

func generateLyricForDay(day int) string {
	lyricForDay := map[int]string{
		1:  "a Partridge in a Pear Tree",
		2:  "two Turtle Doves",
		3:  "three French Hens",
		4:  "four Calling Birds",
		5:  "five Gold Rings",
		6:  "six Geese-a-Laying",
		7:  "seven Swans-a-Swimming",
		8:  "eight Maids-a-Milking",
		9:  "nine Ladies Dancing",
		10: "ten Lords-a-Leaping",
		11: "eleven Pipers Piping",
		12: "twelve Drummers Drumming",
	}

	var line string
	for i := day; i > 0; i-- {
		line += lyricForDay[i]
		if i > 1 {
			line += ", "
		}

		if i == 2 {
			line += "and "
		}
	}

	return line
}

func Song() string {
	var song string
	const maxLineOfVerses = 12
	for i := 1; i <= maxLineOfVerses; i++ {
		song += Verse(i)
		if i < maxLineOfVerses {
			song += "\n"
		}
	}
	return song
}
