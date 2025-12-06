package storage

import "encoding/json"

func toJSON(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func fromJSON(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}
