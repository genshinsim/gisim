package character

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

type charData struct {
	Config
	Data         *model.AvatarData
	SkillLvlData []skillLvlData
}

type skillLvlData struct {
	VarName string
	Type    string
	Values  []float64
	Index   int
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
		err = writePBToFile(fmt.Sprintf("%v/data_gen.pb", v.RelativePath), dm)
		if err != nil {
			return err
		}

		buff := new(bytes.Buffer)
		d := charData{
			Config: v,
			Data:   dm,
		}
		if d.CharStructName == "" {
			d.CharStructName = "char"
		}
		d.buildSkillData(dm)
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

func writePBToFile(path string, dm *model.AvatarData) error {
	// get rid of unncessary fields before saving
	msg := proto.Clone(dm).(*model.AvatarData)
	msg.SkillDetails.AttackScaling = nil
	msg.SkillDetails.SkillScaling = nil
	msg.SkillDetails.BurstScaling = nil
	// write the avatar data to []byte
	opt := prototext.MarshalOptions{
		Multiline: true,
		Indent:    "   ",
	}
	b, err := opt.Marshal(msg)
	if err != nil {
		log.Printf("error marshalling %v data to proto\n", dm.Key)
		return err
	}
	return os.WriteFile(path, b, os.ModePerm)
}

func (c *charData) buildSkillData(dm *model.AvatarData) {
	atk := c.skillDataByType("attack", dm.SkillDetails.AttackScaling)
	skill := c.skillDataByType("skill", dm.SkillDetails.SkillScaling)
	burst := c.skillDataByType("burst", dm.SkillDetails.BurstScaling)
	c.SkillLvlData = append(c.SkillLvlData, atk...)
	c.SkillLvlData = append(c.SkillLvlData, skill...)
	c.SkillLvlData = append(c.SkillLvlData, burst...)
}

func (c *charData) skillDataByType(t string, data []*model.AvatarSkillExcelIndexData) []skillLvlData {
	// build mapping for replacement var names
	varNameMapping := c.Config.SkillDataMapping[t]

	res := make([]skillLvlData, len(data))
	for i, v := range data {
		res[i] = skillLvlData{
			VarName: fmt.Sprintf("%vIndex%v", t, v.Index),
			Index:   int(v.Index),
			Type:    t,
		}
		// this is already pre-sorted
		for _, ld := range v.LevelData {
			res[i].Values = append(res[i].Values, ld.Value)
		}
		if val, ok := varNameMapping[int(v.Index)]; ok {
			res[i].VarName = val
		}
	}

	return res
}

const tmpl = `// Code generated by "pipeline"; DO NOT EDIT.
package {{.PackageName}}

import (
	_ "embed"

    "github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.pb
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}

func (x *{{.CharStructName}}) Data() *model.AvatarData {
	return base
}

{{if .GenerateSkillData}}
var (
{{range $element := .SkillLvlData}}
	//{{$element.Type}} index {{$element.Index}}
    {{$element.VarName}} = []float64{
	{{range $v := $element.Values -}}
	{{$v}},
	{{end}}
}
{{end}}
)
{{end}}
`
