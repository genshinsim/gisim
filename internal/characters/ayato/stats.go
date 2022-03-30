package ayato

var (
	attack = [][][]float64{
		//1
		{
			{
				44.96 / 100,
				48.62 / 100,
				52.28 / 100,
				57.51 / 100,
				61.17 / 100,
				65.35 / 100,
				71.1 / 100,
				76.85 / 100,
				82.6 / 100,
				88.88 / 100,
				95.15 / 100,
				101.43 / 100,
				107.7 / 100,
				113.97 / 100,
				120.25 / 100,
			},
		},
		//2
		{
			{
				47.16 / 100,
				51.0 / 100,
				54.83 / 100,
				60.32 / 100,
				64.16 / 100,
				68.54 / 100,
				74.57 / 100,
				80.61 / 100,
				86.64 / 100,
				93.22 / 100,
				99.8 / 100,
				106.38 / 100,
				112.96 / 100,
				119.54 / 100,
				126.12 / 100,
			},
		},
		//3
		{
			{
				58.61 / 100,
				63.38 / 100,
				68.15 / 100,
				74.97 / 100,
				79.74 / 100,
				85.19 / 100,
				92.69 / 100,
				100.19 / 100,
				107.68 / 100,
				115.86 / 100,
				124.04 / 100,
				132.22 / 100,
				140.4 / 100,
				148.58 / 100,
				156.75 / 100,
			},
		},
		//4
		{
			{
				29.45 / 100,
				31.85 / 100,
				34.24 / 100,
				37.67 / 100,
				40.06 / 100,
				42.8 / 100,
				46.57 / 100,
				50.34 / 100,
				54.1 / 100,
				58.21 / 100,
				62.32 / 100,
				66.43 / 100,
				70.54 / 100,
				74.65 / 100,
				78.76 / 100,
			},
			{
				29.45 / 100,
				31.85 / 100,
				34.24 / 100,
				37.67 / 100,
				40.06 / 100,
				42.8 / 100,
				46.57 / 100,
				50.34 / 100,
				54.1 / 100,
				58.21 / 100,
				62.32 / 100,
				66.43 / 100,
				70.54 / 100,
				74.65 / 100,
				78.76 / 100,
			},
		},
		//5
		{
			{
				75.6 / 100,
				81.76 / 100,
				87.91 / 100,
				96.7 / 100,
				102.86 / 100,
				109.89 / 100,
				119.56 / 100,
				129.23 / 100,
				138.9 / 100,
				149.45 / 100,
				160 / 100,
				170.55 / 100,
				181.1 / 100,
				191.65 / 100,
				202.2 / 100,
			},
		},
	}
	ca = []float64{
		129.53 / 100,
		140.07 / 100,
		150.62 / 100,
		165.68 / 100,
		176.22 / 100,
		188.27 / 100,
		204.84 / 100,
		221.41 / 100,
		237.97 / 100,
		256.05 / 100,
		274.12 / 100,
		292.19 / 100,
		310.27 / 100,
		328.34 / 100,
		346.42 / 100,
	}
	shunsuiken = [][][]float64{
		{{
			52.89 / 100,
			57.2 / 100,
			61.5 / 100,
			67.65 / 100,
			71.96 / 100,
			76.88 / 100,
			83.64 / 100,
			90.41 / 100,
			97.17 / 100,
			104.55 / 100,
			111.93 / 100,
			119.31 / 100,
			126.69 / 100,
		}},
		{{
			58.91 / 100,
			63.71 / 100,
			68.5 / 100,
			75.35 / 100,
			80.15 / 100,
			85.63 / 100,
			93.16 / 100,
			100.7 / 100,
			108.23 / 100,
			116.45 / 100,
			124.67 / 100,
			132.89 / 100,
			141.11 / 100,
		}},
		{{
			64.93 / 100,
			70.22 / 100,
			75.5 / 100,
			83.05 / 100,
			88.34 / 100,
			94.38 / 100,
			102.68 / 100,
			110.99 / 100,
			119.29 / 100,
			128.35 / 100,
			137.41 / 100,
			146.47 / 100,
			155.53 / 100,
		}},
	}
	skill = []float64{
		101.48 / 100,
		109.74 / 100,
		118.0 / 100,
		129.8 / 100,
		138.06 / 100,
		147.5 / 100,
		160.48 / 100,
		173.46 / 100,
		186.44 / 100,
		200.6 / 100,
		214.76 / 100,
		228.92 / 100,
		243.08 / 100,
		257.24 / 100,
		271.4 / 100,
	}
	skillpp = []float64{
		0.44 / 100,
		0.47 / 100,
		0.51 / 100,
		0.56 / 100,
		0.6 / 100,
		0.64 / 100,
		0.69 / 100,
		0.75 / 100,
		0.81 / 100,
		0.87 / 100,
		0.93 / 100,
		0.99 / 100,
		1.05 / 100,
		1.11 / 100,
		1.17 / 100,
	}
	burst = []float64{
		55.02 / 100,
		59.14 / 100,
		63.27 / 100,
		68.77 / 100,
		72.9 / 100,
		77.02 / 100,
		82.53 / 100,
		88.03 / 100,
		93.53 / 100,
		99.03 / 100,
		104.53 / 100,
		110.04 / 100,
		116.91 / 100,
		123.79 / 100,
		130.67 / 100,
	}
	burstasp = []float64{
		101.48 / 100,
		109.74 / 100,
		118 / 100,
		129.8 / 100,
		138.06 / 100,
		147.5 / 100,
		160.48 / 100,
		173.46 / 100,
		186.44 / 100,
		200.6 / 100,
		214.76 / 100,
		228.92 / 100,
		243.08 / 100,
		257.24 / 100,
		271.4 / 100,
	}
)
