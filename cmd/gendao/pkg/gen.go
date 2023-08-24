package pkg

import (
	"bytes"
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"reflect"
	"strings"
	"unicode"
)

// DaoGenerator is a generator for dao.
type DaoGenerator struct {
	buf      *bytes.Buffer
	config   *config
	savePath string
}

// field struct.
type field struct {
	FieldName     string
	NormFieldName string
	FieldType     string
}

// config struct.
type config struct {
	PkgName       string
	NormPkgName   string
	ModelName     string
	NormModelName string
	Fields        []field
}

// getFields get fields from struct.
func getFields(m interface{}) []field {
	return getFieldsRecursively(reflect.TypeOf(m))
}

// getFieldsRecursively get fields recursively.
func getFieldsRecursively(value reflect.Type) []field {
	var fields []field

	for i := 0; i < value.NumField(); i++ {
		v := value.Field(i)

		// Deal with *time.Time
		// Ref: func (field Time) Eq(value time.Time) Expr {} in gorm-gen
		if v.Type.Kind() == reflect.Ptr && v.Type.Elem().String() == "time.Time" {
			f := field{
				FieldName:     v.Name,
				NormFieldName: strings.ToLower(v.Name),
				FieldType:     "time.Time",
			}
			fields = append(fields, f)

			continue
		}

		// Get the tag value
		if v.Tag != "" && v.Type.Kind() != reflect.Struct &&
			v.Type.Kind() != reflect.Bool {
			f := field{
				FieldName:     v.Name,
				NormFieldName: strings.ToLower(v.Name),
				FieldType:     v.Type.String(),
			}
			fields = append(fields, f)

			continue
		}

		// Get the embedded struct fields if the tag is empty
		if v.Tag == "" && v.Type.Kind() == reflect.Struct {
			fields = append(fields, getFieldsRecursively(v.Type)...)
		}
	}

	return fields
}

// NewDaoGenerator returns a new DaoGenerator.
func NewDaoGenerator(m interface{}, savePath string) *DaoGenerator {
	pkgPath := reflect.TypeOf(m).PkgPath()
	paths := strings.Split(pkgPath, "/")
	pkgName := paths[len(paths)-1]

	return &DaoGenerator{
		buf: bytes.NewBuffer(nil),
		config: &config{
			PkgName:       pkgName,
			NormPkgName:   strings.ReplaceAll(pkgName, "_", ""),
			ModelName:     reflect.TypeOf(m).Name(),
			NormModelName: strings.ToLower(reflect.TypeOf(m).Name()),
			Fields:        getFields(m),
		},
		savePath: savePath,
	}
}

// Generate init the hook.
func (g *DaoGenerator) Generate() *DaoGenerator {
	if err := hookTemplate.Execute(g.buf, g.config); err != nil {
		panic(err)
	}

	return g
}

// Format format the generated code.
func (g *DaoGenerator) Format() *DaoGenerator {
	formatOut, err := format.Source(g.buf.Bytes())
	if err != nil {
		panic(err)
	}
	g.buf = bytes.NewBuffer(formatOut)

	return g
}

// Flush write the generated code to file.
func (g *DaoGenerator) Flush() {
	filename := fmt.Sprintf("gen_%s_dao.go", camelToUnderscore(g.config.ModelName))
	if err := os.WriteFile(
		fmt.Sprintf("%s/%s", g.savePath, filename),
		g.buf.Bytes(),
		fs.ModePerm); err != nil {
		panic(err)
	}
	fmt.Println("[kvm-agent] gen file ok: ", fmt.Sprintf("%s/%s", g.savePath, filename))
}

// camelToUnderscore convert camel to underscore.
func camelToUnderscore(s string) string {
	var result []rune

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 && (unicode.IsLower(rune(s[i-1])) || (i+1 < len(s) && unicode.IsLower(rune(s[i+1])))) {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}

	return strings.ToLower(string(result))
}
