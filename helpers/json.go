package helpers

import "encoding/json"

type JSON map[string]interface{}

func (j JSON) String() string {
	data, _ := json.MarshalIndent(j, "", "  ")
	return string(data)
}

func (j JSON) Bytes() []byte {
	data, _ := json.MarshalIndent(j, "", "  ")
	return data
}

func ToJSON(data []byte) JSON {
	var result JSON
	_ = json.Unmarshal(data, &result)
	return result
}

func ToStruct[T any](data []byte) *T {
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil
	}
	return &result
}

func StructToJSON[T any](data *T) JSON {
	bytes, _ := json.MarshalIndent(data, "", "  ")
	return ToJSON(bytes)
}

func Parse[T any](data string) T {
	var result T
	_ = json.Unmarshal([]byte(data), &result)
	return result
}