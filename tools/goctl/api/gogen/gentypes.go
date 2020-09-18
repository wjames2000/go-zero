package gogen

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"path"
	"strings"
	"text/template"

	"github.com/wjames2000/go-zero/tools/goctl/api/spec"
	apiutil "github.com/wjames2000/go-zero/tools/goctl/api/util"
	"github.com/wjames2000/go-zero/tools/goctl/util"
)

const (
	typesFile     = "types.go"
	typesTemplate = `// DO NOT EDIT, generated by goctl
package types{{if .containsTime}}
import (
	"time"
){{end}}
{{.types}}
`
)

func BuildTypes(types []spec.Type) (string, error) {
	var builder strings.Builder
	first := true
	for _, tp := range types {
		if first {
			first = false
		} else {
			builder.WriteString("\n\n")
		}
		if err := writeType(&builder, tp, types); err != nil {
			return "", apiutil.WrapErr(err, "Type "+tp.Name+" generate error")
		}
	}

	return builder.String(), nil
}

func genTypes(dir string, api *spec.ApiSpec, force bool) error {
	val, err := BuildTypes(api.Types)
	if err != nil {
		return err
	}

	filename := path.Join(dir, typesDir, typesFile)
	if !force {
		if err := util.RemoveOrQuit(filename); err != nil {
			return err
		}
	}

	fp, created, err := apiutil.MaybeCreateFile(dir, typesDir, typesFile)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	t := template.Must(template.New("typesTemplate").Parse(typesTemplate))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, map[string]interface{}{
		"types":        val,
		"containsTime": api.ContainsTime(),
	})
	if err != nil {
		return nil
	}
	formatCode := formatCode(buffer.String())
	_, err = fp.WriteString(formatCode)
	return err
}

func convertTypeCase(types []spec.Type, t string) (string, error) {
	ts, err := apiutil.DecomposeType(t)
	if err != nil {
		return "", err
	}

	var defTypes []string
	for _, tp := range ts {
		for _, typ := range types {
			if typ.Name == tp {
				defTypes = append(defTypes, tp)
			}

			if len(typ.Annotations) > 0 {
				if value, ok := apiutil.GetAnnotationValue(typ.Annotations, "serverReplacer", tp); ok {
					t = strings.ReplaceAll(t, tp, value)
				}
			}
		}
	}

	for _, tp := range defTypes {
		t = strings.ReplaceAll(t, tp, util.Title(tp))
	}

	return t, nil
}

func writeType(writer io.Writer, tp spec.Type, types []spec.Type) error {
	fmt.Fprintf(writer, "type %s struct {\n", util.Title(tp.Name))
	for _, member := range tp.Members {
		if member.IsInline {
			var found = false
			for _, ty := range types {
				if strings.ToLower(ty.Name) == strings.ToLower(member.Name) {
					found = true
				}
			}
			if !found {
				return errors.New("inline type " + member.Name + " not exist, please correct api file")
			}
			if _, err := fmt.Fprintf(writer, "%s\n", strings.Title(member.Type)); err != nil {
				return err
			} else {
				continue
			}
		}
		tpString, err := convertTypeCase(types, member.Type)
		if err != nil {
			return err
		}
		pm, err := member.GetPropertyName()
		if err != nil {
			return err
		}
		if !strings.Contains(pm, "_") {
			if strings.Title(member.Name) != strings.Title(pm) {
				fmt.Printf("type: %s, property name %s json tag illegal, "+
					"should set json tag as `json:\"%s\"` \n", tp.Name, member.Name, util.Untitle(member.Name))
			}
		}
		if err := writeProperty(writer, member.Name, tpString, member.Tag, member.GetComment(), 1); err != nil {
			return err
		}
	}
	fmt.Fprintf(writer, "}")
	return nil
}
