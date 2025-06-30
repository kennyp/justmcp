package parser

type Recipe struct {
	Name       string       `json:"name"`
	Doc        Doc          `json:"doc"`
	Attributes *Attributes  `json:"attributes"`
	Parameters []*Parameter `json:"parameters"`
	Private    bool         `json:"private"`
}
