package models

type ValueType string

// ${{ grid:..... }}
const (
	Var    ValueType = "var"
	Secret ValueType = "secret"
	Grid   ValueType = "grid"
)

type Value struct {
	Key       string    `json:"key"`
	ValueType ValueType `json:"type"`
	Value     *string   `json:"value"`
}

func (v Value) Resolved() (string, bool) {
	if v.Value != nil {
		return *v.Value, true
	}
	return "", false
}
