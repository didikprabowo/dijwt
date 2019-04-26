package helpers

import (
	"encoding/base64"
	"strings"
)

func Base64Encode(text string) string {
	ceking := base64.URLEncoding.EncodeToString([]byte(text))
	r := strings.NewReplacer("+", "", "/", "")
	result := r.Replace(ceking)
	sfine := strings.TrimRight(result, "=")
	return sfine
}

func Base64Decode(text string) (string, error) {
	r := strings.NewReplacer("-_", "+/")
	result := r.Replace(text)
	lens := len(result) % 4
	if lens > 0 {
		result += strings.Repeat("=", 4-1)
	}
	decoded, _ := base64.URLEncoding.DecodeString(result)
	return string(decoded), nil
}
