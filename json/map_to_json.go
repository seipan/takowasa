package json

import "encoding/json"

func PaarseMaptoBody(mp map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
