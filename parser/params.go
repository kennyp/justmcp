package parser

type Parameter struct {
	Name    string `json:"name"`
	Doc     string `json:"-"`
	Kind    string `json:"kind"`
	Default string `json:"default"`
}
