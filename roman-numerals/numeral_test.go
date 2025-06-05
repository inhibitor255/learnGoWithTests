package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {

	t.Run("various cases", func(t *testing.T) {
		for _, test := range cases {
			t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
				got, _ := ConvertToRoman(test.Arabic)
				if got != test.Roman {
					t.Errorf("got %q, want %q", got, test.Roman)
				}
			})
		}
	})

	t.Run("with number greater than 3999", func(t *testing.T) {
		_, err := ConvertToRoman(10000)
		want := ErrOutOfRange
		if err == nil {
			t.Fatal("want err but get none")
		}

		if err != want {
			t.Errorf("err %q, want %q", err, want)
		}

	})

}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfRoman_Numerals(t *testing.T) {
	t.Run("Conversion is reversible", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true // skip values out of range
			}
			roman, _ := ConvertToRoman(arabic)
			fromRoman := ConvertToArabic(roman)
			return fromRoman == arabic
		}

		if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
			t.Error("failed checks", err)
		}
	})

	t.Run("Only I, X, and C can be subtractors", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true
			}
			roman, _ := ConvertToRoman(arabic)

			// Look for any subtractive pair (e.g. IV, XL, etc.)
			for i := 0; i < len(roman)-1; i++ {
				curr := roman[i]
				next := roman[i+1]

				if valueOf(curr) < valueOf(next) {
					// This is a subtraction (e.g. I before V, etc.)
					if curr != 'I' && curr != 'X' && curr != 'C' {
						return false // invalid subtractor
					}
				}
			}
			return true
		}

		if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
			t.Error("Invalid subtractor found", err)
		}
	})

	t.Run("Can't have more than 3 of the same symbol in a row", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true
			}
			roman, _ := ConvertToRoman(arabic)

			count := 1
			for i := 1; i < len(roman); i++ {
				if roman[i] == roman[i-1] {
					count++
					if count > 3 {
						return false
					}
				} else {
					count = 1
				}
			}
			return true
		}

		if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
			t.Error("More than 3 consecutive symbols found", err)
		}
	})
}

func BenchmarkConvertToRoman(b *testing.B) {
	for _, test := range cases {
		b.Run(fmt.Sprintf("%d", test.Arabic), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = ConvertToRoman(test.Arabic)
			}
		})
	}
}

func BenchmarkConvertToArabic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range cases {
			b.Run(test.Roman, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = ConvertToArabic(test.Roman)
				}
			})
		}
	}
}

func valueOf(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
