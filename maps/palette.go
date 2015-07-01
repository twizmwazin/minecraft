package maps

import "image/color"

var palette = color.Palette{
	color.RGBA{0, 0, 0, 0},
	color.RGBA{0, 0, 0, 0},
	color.RGBA{0, 0, 0, 0},
	color.RGBA{0, 0, 0, 0},
	color.RGBA{89, 125, 39, 255},
	color.RGBA{109, 153, 48, 255},
	color.RGBA{127, 178, 56, 255},
	color.RGBA{67, 94, 29, 255},
	color.RGBA{89, 125, 39, 255},
	color.RGBA{109, 153, 48, 255},
	color.RGBA{127, 178, 56, 255},
	color.RGBA{67, 94, 29, 255},
	color.RGBA{117, 117, 117, 255},
	color.RGBA{144, 144, 144, 255},
	color.RGBA{167, 167, 167, 255},
	color.RGBA{88, 88, 88, 255},
	color.RGBA{117, 117, 117, 255},
	color.RGBA{144, 144, 144, 255},
	color.RGBA{167, 167, 167, 255},
	color.RGBA{88, 88, 88, 255},
	color.RGBA{112, 112, 180, 255},
	color.RGBA{138, 138, 220, 255},
	color.RGBA{160, 160, 255, 255},
	color.RGBA{84, 84, 135, 255},
	color.RGBA{117, 117, 117, 255},
	color.RGBA{144, 144, 144, 255},
	color.RGBA{167, 167, 167, 255},
	color.RGBA{88, 88, 88, 255},
	color.RGBA{0, 87, 0, 255},
	color.RGBA{0, 106, 0, 255},
	color.RGBA{0, 124, 0, 255},
	color.RGBA{0, 65, 0, 255},
	color.RGBA{180, 180, 180, 255},
	color.RGBA{220, 220, 220, 255},
	color.RGBA{255, 255, 255, 255},
	color.RGBA{135, 135, 135, 255},
	color.RGBA{115, 118, 129, 255},
	color.RGBA{141, 144, 158, 255},
	color.RGBA{164, 168, 184, 255},
	color.RGBA{86, 88, 97, 255},
	color.RGBA{129, 74, 33, 255},
	color.RGBA{157, 91, 40, 255},
	color.RGBA{183, 106, 47, 255},
	color.RGBA{96, 56, 24, 255},
	color.RGBA{79, 79, 79, 255},
	color.RGBA{96, 96, 96, 255},
	color.RGBA{112, 112, 112, 255},
	color.RGBA{59, 59, 59, 255},
	color.RGBA{45, 45, 180, 255},
	color.RGBA{55, 55, 220, 255},
	color.RGBA{64, 64, 255, 255},
	color.RGBA{33, 33, 135, 255},
	color.RGBA{45, 45, 180, 255},
	color.RGBA{55, 55, 220, 255},
	color.RGBA{64, 64, 255, 255},
	color.RGBA{33, 33, 135, 255},
	color.RGBA{180, 177, 172, 255},
	color.RGBA{220, 217, 211, 255},
	color.RGBA{255, 252, 245, 255},
	color.RGBA{135, 133, 129, 255},
	color.RGBA{152, 89, 36, 255},
	color.RGBA{186, 109, 44, 255},
	color.RGBA{216, 127, 51, 255},
	color.RGBA{114, 67, 27, 255},
	color.RGBA{125, 53, 152, 255},
	color.RGBA{153, 65, 186, 255},
	color.RGBA{178, 76, 216, 255},
	color.RGBA{94, 40, 114, 255},
	color.RGBA{72, 108, 152, 255},
	color.RGBA{88, 132, 186, 255},
	color.RGBA{102, 153, 216, 255},
	color.RGBA{54, 81, 114, 255},
	color.RGBA{161, 161, 36, 255},
	color.RGBA{197, 197, 44, 255},
	color.RGBA{229, 229, 51, 255},
	color.RGBA{121, 121, 27, 255},
	color.RGBA{89, 144, 17, 255},
	color.RGBA{109, 176, 21, 255},
	color.RGBA{127, 204, 25, 255},
	color.RGBA{67, 108, 13, 255},
	color.RGBA{170, 89, 116, 255},
	color.RGBA{208, 109, 142, 255},
	color.RGBA{242, 127, 165, 255},
	color.RGBA{128, 67, 87, 255},
	color.RGBA{53, 53, 53, 255},
	color.RGBA{65, 65, 65, 255},
	color.RGBA{76, 76, 76, 255},
	color.RGBA{40, 40, 40, 255},
	color.RGBA{108, 108, 108, 255},
	color.RGBA{132, 132, 132, 255},
	color.RGBA{153, 153, 153, 255},
	color.RGBA{81, 81, 81, 255},
	color.RGBA{53, 89, 108, 255},
	color.RGBA{65, 109, 132, 255},
	color.RGBA{76, 127, 153, 255},
	color.RGBA{40, 67, 81, 255},
	color.RGBA{89, 44, 125, 255},
	color.RGBA{109, 54, 153, 255},
	color.RGBA{127, 63, 178, 255},
	color.RGBA{67, 33, 94, 255},
	color.RGBA{36, 53, 125, 255},
	color.RGBA{44, 65, 153, 255},
	color.RGBA{51, 76, 178, 255},
	color.RGBA{27, 40, 94, 255},
	color.RGBA{72, 53, 36, 255},
	color.RGBA{88, 65, 44, 255},
	color.RGBA{102, 76, 51, 255},
	color.RGBA{54, 40, 27, 255},
	color.RGBA{72, 89, 36, 255},
	color.RGBA{88, 109, 44, 255},
	color.RGBA{102, 127, 51, 255},
	color.RGBA{54, 67, 27, 255},
	color.RGBA{108, 36, 36, 255},
	color.RGBA{132, 44, 44, 255},
	color.RGBA{153, 51, 51, 255},
	color.RGBA{81, 27, 27, 255},
	color.RGBA{17, 17, 17, 255},
	color.RGBA{21, 21, 21, 255},
	color.RGBA{25, 25, 25, 255},
	color.RGBA{13, 13, 13, 255},
	color.RGBA{176, 168, 54, 255},
	color.RGBA{215, 205, 66, 255},
	color.RGBA{250, 238, 77, 255},
	color.RGBA{132, 126, 40, 255},
	color.RGBA{64, 154, 150, 255},
	color.RGBA{79, 188, 183, 255},
	color.RGBA{92, 219, 213, 255},
	color.RGBA{48, 115, 112, 255},
	color.RGBA{52, 90, 180, 255},
	color.RGBA{63, 110, 220, 255},
	color.RGBA{74, 128, 255, 255},
	color.RGBA{39, 67, 135, 255},
	color.RGBA{0, 153, 40, 255},
	color.RGBA{0, 187, 50, 255},
	color.RGBA{0, 217, 58, 255},
	color.RGBA{0, 114, 30, 255},
	color.RGBA{14, 14, 21, 255},
	color.RGBA{18, 17, 26, 255},
	color.RGBA{21, 20, 31, 255},
	color.RGBA{11, 10, 16, 255},
	color.RGBA{79, 1, 0, 255},
	color.RGBA{96, 1, 0, 255},
	color.RGBA{112, 2, 0, 255},
	color.RGBA{59, 1, 0, 255},
	color.RGBA{88, 59, 33, 255},
	color.RGBA{108, 72, 41, 255},
	color.RGBA{126, 84, 48, 255},
	color.RGBA{66, 44, 25, 255},
}

/*var palette color.Palette

func init() {
	colours := [...]color.RGBA{
		{0, 0, 0, 0},
		{127, 178, 56, 255},
		{127, 178, 56, 255},
		{167, 167, 167, 255},
		{167, 167, 167, 255},
		{160, 160, 255, 255},
		{167, 167, 167, 255},
		{0, 124, 0, 255},
		{255, 255, 255, 255},
		{164, 168, 184, 255},
		{183, 106, 47, 255},
		{112, 112, 112, 255},
		{64, 64, 255, 255},
		{64, 64, 255, 255},
		{255, 252, 245, 255},
		{216, 127, 51, 255},
		{178, 76, 216, 255},
		{102, 153, 216, 255},
		{229, 229, 51, 255},
		{127, 204, 25, 255},
		{242, 127, 165, 255},
		{76, 76, 76, 255},
		{153, 153, 153, 255},
		{76, 127, 153, 255},
		{127, 63, 178, 255},
		{51, 76, 178, 255},
		{102, 76, 51, 255},
		{102, 127, 51, 255},
		{153, 51, 51, 255},
		{25, 25, 25, 255},
		{250, 238, 77, 255},
		{92, 219, 213, 255},
		{74, 128, 255, 255},
		{0, 217, 58, 255},
		{21, 20, 31, 255},
		{112, 2, 0, 255},
		{126, 84, 48, 255},
	}
	palette = make(color.Palette, len(colours)*4)
	for n, colour := range colours {
		palette[n<<2|0] = multiplyColour(colour, 180)
		palette[n<<2|1] = multiplyColour(colour, 220)
		palette[n<<2|2] = multiplyColour(colour, 255)
		palette[n<<2|3] = multiplyColour(colour, 135)
	}
}

func multiplyColour(c color.RGBA, mult uint8) color.RGBA {
	return color.RGBA{
		uint8(uint(c.R) * uint(mult) / 255),
		uint8(uint(c.G) * uint(mult) / 255),
		uint8(uint(c.B) * uint(mult) / 255),
		c.A,
	}
}*/