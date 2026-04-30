package utils

import (
    "math/rand"
    "strings"
    "time"
    "unicode"
    "strconv"
)

// generate 7 char dari nama company
func GenerateAccountCode(name string) string {
    clean := ""

    for _, r := range name {
        if unicode.IsLetter(r) {
            clean += string(r)
        }
    }

    clean = strings.ToUpper(clean)

    if len(clean) > 5 {
        clean = clean[:5]
    }

    const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano())

    suffix := make([]byte, 2)
    for i := range suffix {
        suffix[i] = charset[rand.Intn(len(charset))]
    }

    return clean + string(suffix) // total 7 char
}

// generate 6 char company id
func GenerateCompanyID() string {
	rand.Seed(time.Now().UnixNano())

	min := 100000
	max := 999999

	return strconv.Itoa(rand.Intn(max-min+1) + min)
}