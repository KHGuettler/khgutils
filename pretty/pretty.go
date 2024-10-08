package pretty

import (
	"encoding/json"
	"fmt"
)

func Print(v any) error {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}

func Fmt(v any) string {
	b, _ := json.MarshalIndent(v, "", " ")

	return string(b)
}
