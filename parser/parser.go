package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"os/exec"
)

type Justfile struct {
	Path    string             `json:"-"`
	Recipes map[string]*Recipe `json:"recipes"`
}

func Parse(ctx context.Context, filePath string) (*Justfile, error) {
	cmd := exec.CommandContext(ctx, "just", "-f", filePath, "--dump", "--dump-format", "json")

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("loading justfile failed (%w)", err)
	}

	var f Justfile
	if err := json.Unmarshal(out, &f); err != nil {
		return nil, fmt.Errorf("failed to parse justfile (%w)", err)
	}

	f.Path = filePath

	for _, recipe := range f.Recipes {
		for _, p := range recipe.Parameters {
			p.Doc = recipe.Doc.ParamDoc(p.Name)
		}
	}

	return &f, nil
}

func (f *Justfile) PublicRecipes() iter.Seq2[string, *Recipe] {
	return func(yield func(string, *Recipe) bool) {
		for k, v := range f.Recipes {
			if v.Private {
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}
