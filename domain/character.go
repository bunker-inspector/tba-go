package domain

const (
	CharacterAbilityMin     = 8
	CharacterAbilityInitMax = 14
)

type Character struct {
	Name  string
	Level uint8
	Exp   uint16

	UserID int

	Str uint8
	Con uint8
	Dex uint8
	Int uint8
	Wis uint8
	Cha uint8
}

func BaseCharacter(userid int, name string) Character {
	return Character{
		Name:  name,
		Level: 1,
		Exp:   0,

		UserID: userid,

		Str: 10,
		Con: 10,
		Dex: 10,
		Int: 10,
		Wis: 10,
		Cha: 10,
	}
}
