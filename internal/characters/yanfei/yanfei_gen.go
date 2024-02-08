// Code generated by "pipeline"; DO NOT EDIT.
package yanfei

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.5834159851074219,
		0.6271719932556152,
		0.6709280014038086,
		0.7292699813842773,
		0.7730259895324707,
		0.8167819976806641,
		0.8751239776611328,
		0.9334660172462463,
		0.9918069839477539,
		1.0501489639282227,
		1.108489990234375,
		1.1668319702148438,
		1.2397589683532715,
		1.3126859664916992,
		1.385612964630127,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.5212560296058655,
		0.5603500008583069,
		0.5994439721107483,
		0.6515700221061707,
		0.6906639933586121,
		0.7297580242156982,
		0.7818840146064758,
		0.8340100049972534,
		0.8861349821090698,
		0.9382609724998474,
		0.9903860092163086,
		1.042512059211731,
		1.1076689958572388,
		1.1728260517120361,
		1.237982988357544,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.7601280212402344,
		0.8171380162239075,
		0.8741469979286194,
		0.950160026550293,
		1.0071699619293213,
		1.0641789436340332,
		1.1401920318603516,
		1.2162050008773804,
		1.2922179698944092,
		1.3682299852371216,
		1.4442429542541504,
		1.5202560424804688,
		1.615272045135498,
		1.7102880477905273,
		1.8053040504455566,
	}
	// attack: charge = [3 4 5 6 7]
	charge = [][]float64{
		{
			0.9822940230369568,
			1.0411139726638794,
			1.0999339818954468,
			1.1763999462127686,
			1.235219955444336,
			1.2940399646759033,
			1.3705060482025146,
			1.4469720125198364,
			1.5234379768371582,
			1.59990394115448,
			1.6763700246810913,
			1.752835988998413,
			1.8293019533157349,
			1.9057680368423462,
			1.982234001159668,
		},
		{
			1.1556400060653687,
			1.2248400449752808,
			1.2940399646759033,
			1.3839999437332153,
			1.4531999826431274,
			1.5224000215530396,
			1.6123600006103516,
			1.7023199796676636,
			1.7922799587249756,
			1.8822400569915771,
			1.9722000360488892,
			2.062160015106201,
			2.1521201133728027,
			2.242079973220825,
			2.3320400714874268,
		},
		{
			1.3289860486984253,
			1.4085659980773926,
			1.4881459474563599,
			1.591599941253662,
			1.671180009841919,
			1.7507599592208862,
			1.8542139530181885,
			1.9576679468154907,
			2.061121940612793,
			2.1645760536193848,
			2.2680299282073975,
			2.3714840412139893,
			2.474937915802002,
			2.5783920288085938,
			2.6818459033966064,
		},
		{
			1.5023319721221924,
			1.5922919511795044,
			1.682252049446106,
			1.7992000579833984,
			1.8891600370407104,
			1.9791200160980225,
			2.0960679054260254,
			2.2130160331726074,
			2.3299639225006104,
			2.4469120502471924,
			2.5638599395751953,
			2.6808080673217773,
			2.7977559566497803,
			2.9147040843963623,
			3.0316519737243652,
		},
		{
			1.675678014755249,
			1.7760180234909058,
			1.8763580322265625,
			2.0067999362945557,
			2.107140064239502,
			2.207479953765869,
			2.3379220962524414,
			2.4683640003204346,
			2.5988059043884277,
			2.729248046875,
			2.859689950942993,
			2.9901320934295654,
			3.1205739974975586,
			3.2510159015655518,
			3.381458044052124,
		},
	}
	// skill: skill = [0]
	skill = []float64{
		1.6959999799728394,
		1.823199987411499,
		1.9503999948501587,
		2.119999885559082,
		2.2472000122070312,
		2.3743999004364014,
		2.5439999103546143,
		2.713599920272827,
		2.88319993019104,
		3.052799940109253,
		3.222399950027466,
		3.3919999599456787,
		3.6040000915527344,
		3.815999984741211,
		4.0279998779296875,
	}
	// burst: burst = [0]
	burst = []float64{
		1.8240000009536743,
		1.960800051689148,
		2.097599983215332,
		2.2799999713897705,
		2.416800022125244,
		2.5536000728607178,
		2.7360000610351562,
		2.9184000492095947,
		3.100800037384033,
		3.2832000255584717,
		3.46560001373291,
		3.6480000019073486,
		3.875999927520752,
		4.104000091552734,
		4.331999778747559,
	}
	// burst: burstBonus = [1]
	burstBonus = []float64{
		0.33399999141693115,
		0.3540000021457672,
		0.37400001287460327,
		0.4000000059604645,
		0.41999998688697815,
		0.4399999976158142,
		0.4659999907016754,
		0.492000013589859,
		0.5180000066757202,
		0.5440000295639038,
		0.5699999928474426,
		0.5960000157356262,
		0.621999979019165,
		0.6480000019073486,
		0.6740000247955322,
	}
)
