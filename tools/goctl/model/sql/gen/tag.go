package gen

import (
	"github.com/wjames2000/go-zero/tools/goctl/model/sql/template"
	"github.com/wjames2000/go-zero/tools/goctl/util"
)

func genTag(in string) (string, error) {
	if in == "" {
		return in, nil
	}
	output, err := util.With("tag").
		Parse(template.Tag).
		Execute(map[string]interface{}{
			"field": in,
		})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
