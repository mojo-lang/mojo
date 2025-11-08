package fake

import "github.com/brianvoe/gofakeit/v6"

func PhoneNumber() string {
	prefixes := []string{"130", "131", "132", "133", "134", "135", "136", "137", "138", "139", "145", "155", "175", "176", "185"}
	return prefixes[gofakeit.IntRange(0, len(prefixes)-1)] + gofakeit.DigitN(8)
}
