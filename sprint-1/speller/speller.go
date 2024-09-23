//go:build !solution

package speller

import (
	"fmt"
	"strings"
)

var (
	billion  = int64(1_000_000_000)
	million  = int64(1_000_000)
	thousand = int64(1_000)
)

var tens = map[int64]string{
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eightteen",
	19: "nineteen",
}
var tenOnes = map[int64]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var ones = map[int64]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func Spell(n int64) string {

	spelling := strings.Builder{}
	if n < 0 {
		spelling.WriteString("minus")
		n = n * -1
	}
	if n == 0 {
		return "zero"
	}

	if n/billion != 0 {
		if spelling.String() != "" {
			spelling.WriteString(" ")
		}

		spelling.WriteString(spellHundreds(n / billion))
		spelling.WriteString(" billion")
		n = n % 1_000_000_000
	}
	if n/million != 0 {
		if spelling.String() != "" {
			spelling.WriteString(" ")
		}
		spelling.WriteString(spellHundreds(n / million))
		spelling.WriteString(" million")
		n = n % 1_000_000
	}
	if n/thousand != 0 {
		if spelling.String() != "" {
			spelling.WriteString(" ")
		}
		spelling.WriteString(spellHundreds(n / thousand))
		spelling.WriteString(" thousand")
		n = n % 1000
	}

	if n > 0 {
		hundreds := spellHundreds(n)
		if spelling.String() != "" {
			spelling.WriteString(" " + hundreds)
		} else {
			spelling.WriteString(hundreds)
		}
	}

	return spelling.String()
}

func spellHundreds(n int64) string {

	spelling := strings.Builder{}

	if n/100 != 0 {
		spelling.WriteString(ones[n/100])
		spelling.WriteString(" hundred")
		n = n % 100

		if n > 0 {
			spelling.WriteString(" " + spellTens(n))
		}
		return spelling.String()
	}

	if spelling.String() != "" {
		spelling.WriteString(" ")
	}

	spelling.WriteString(spellTens(n))
	return spelling.String()
}

func spellTens(n int64) string {

	if n < 10 {
		return ones[n]
	}
	if n >= 10 && n < 20 {
		return tens[n]
	}

	ans := tenOnes[n/10]
	if n%10 > 0 {
		return fmt.Sprint(ans, "-", ones[n%10])
	}

	return ans
}
