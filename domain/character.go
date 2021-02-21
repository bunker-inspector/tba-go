package domain

const (
	ABILITY_MIN = 8
	ABILITY_INIT_MAX = 14
)

type Character struct {
	Name  string
	Level uint8
	Exp   uint16

	Str uint8
	Con uint8
	Dex uint8
	Int uint8
	Wis uint8
	Cha uint8
}

func BaseCharacter(name string) Character {
	return Character{
		Name:  name,
		Level: 1,
		Exp:   0,

		Str: 10,
		Con: 10,
		Dex: 10,
		Int: 10,
		Wis: 10,
		Cha: 10,
	}
}
