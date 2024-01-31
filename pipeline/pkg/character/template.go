package character

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/proto"
)

type charData struct {
	Config
	Data *model.AvatarData
}

func (g *Generator) GenerateCharTemplate() error {
	t, err := template.New("chartemplate").Parse(tmpl)
	if err != nil {
		return err
	}
	for _, v := range g.chars {
		dm, ok := g.data[v.Key]
		if !ok {
			log.Printf("No data found for %v; skipping", v.Key)
			continue
		}
		// write the avatar data to []byte
		b, err := proto.Marshal(dm)
		if err != nil {
			log.Printf("error marshalling %v data to proto\n", v.Key)
			return err
		}
		os.WriteFile(fmt.Sprintf("%v/data_gen.pb", v.RelativePath), b, os.ModePerm)

		buff := new(bytes.Buffer)
		d := charData{
			Config: v,
			Data:   dm,
		}
		if d.CharStructName == "" {
			d.CharStructName = "char"
		}
		t.Execute(buff, d)
		src := buff.Bytes()
		dst, err := format.Source(src)
		if err != nil {
			return err
		}
		os.WriteFile(fmt.Sprintf("%v/%v_gen.go", v.RelativePath, v.PackageName), dst, os.ModePerm)
	}

	return nil
}

const tmpl = `package {{.PackageName}}

import (
	_ "embed"
    "github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/proto"
)

//go:embed data_gen.pb
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := proto.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}

func (x *{{.CharStructName}}) Data() *model.AvatarData {
	return base
}
`
