package internal

import "encoding/json"

func InterfaceToStruct(v interface{}, a interface{}) error {
	m, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return json.Unmarshal(m, a)
}
