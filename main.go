package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/didikprabowo/jwt/helpers"
	"strings"
)

const secret = "HIGHLY CONFIDENTIAL SECRET KEY"

type Datas struct {
	Name string
}
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func main() {
	inidata := map[string]interface{}{"sub": "1"}
	object := Header{"HS256", "JWT"}
	var jsonData, _ = json.Marshal(object)
	var payload, _ = json.Marshal(inidata)

	var jsonString = string(jsonData)
	var jsonPayload = string(payload)
	fmt.Println(fmt.Sprintln(GenerateJWT(jsonString, jsonPayload)))
	fmt.Println(fmt.Sprintln(Verif(GenerateJWT(jsonString, jsonPayload))))
}

// genrator
func GenerateJWT(str string, payload string) string {
	headerEncodes := helpers.Base64Encode(str)
	PayloadEncodes := helpers.Base64Encode(payload)
	dataencoded := headerEncodes + "." + PayloadEncodes
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(dataencoded))

	finalSignature := helpers.Base64Encode(string(h.Sum(nil)))
	return dataencoded + "." + finalSignature
}
func Verif(hestol string) bool {
	result := strings.Split(hestol, ".")
	dataEncoded := string(result[0] + "." + result[1])
	signature, _ := helpers.Base64Decode(string(result[2]))

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(dataEncoded))
	calculated := h.Sum(nil)
	return hmac.Equal(calculated, []byte(signature))

}
