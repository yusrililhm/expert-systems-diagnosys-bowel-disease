package helper

import "encoding/json"

func ToJSON(response any) []byte {
	res, _ := json.Marshal(response)
	return res
}
