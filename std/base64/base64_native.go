package base64

import (
	"encoding/base64"

	. "github.com/candid82/joker/core"
)

func base64DecodeString(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(RT.NewError("Invalid bas64 string: " + err.Error()))
	}
	return string(decoded)
}