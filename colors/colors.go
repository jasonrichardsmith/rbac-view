package colors

import "math/rand"

var (
	palette = []string{
		"e6194b",
		"3cb44b",
		"ffe119",
		"0082c8",
		"f58231",
		"911eb4",
		"46f0f0",
		"f032e6",
		"d2f53c",
		"fabebe",
		"008080",
		"e6beff",
		"aa6e28",
		"fffac8",
		"800000",
		"aaffc3",
		"808000",
		"ffd8b1",
		"000080",
		"808080",
		"FFFFFF",
		"000000",
	}
)

func GetUnique() string {
	if len(palette) == 0 {
		panic("no more colors!")
	}
	key := rand.Intn(len(palette))
	color := palette[key]
	palette = append(palette[:key], palette[key+1:]...)
	return color
}
