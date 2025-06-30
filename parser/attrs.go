package parser

import "encoding/json"

type Attributes struct {
	Group        string
	Confirmation string
}

func (attrs *Attributes) UnmarshalJSON(data []byte) error {
	var raw []any
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for _, f := range raw {
		attrMap, ok := f.(map[string]any)
		if !ok {
			continue
		}

		if group, ok := attrMap["group"]; ok {
			attrs.Group, _ = group.(string)
			continue
		}

		if c, ok := attrMap["confirm"]; ok {
			attrs.Confirmation, _ = c.(string)
		}
	}

	return nil
}
