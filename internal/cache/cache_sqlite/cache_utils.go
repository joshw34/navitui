package cache_sqlite

import "encoding/json"

func stringSliceToJSON(s []string) ([]byte, error) {
	return json.Marshal(s)
}

func jsonToStringSlice(b []byte) ([]string, error) {
	if len(b) == 0 {
		return nil, nil
	}
	var s []string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
