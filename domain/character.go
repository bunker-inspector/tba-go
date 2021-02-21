package domain

type Character struct {
	Name  string
	Level uint8
	Exp   uint16

	Strength     uint8
	Constitution uint8
	Dexterity    uint8
	Intelligence uint8
	Wisdom       uint8
	Charisma     uint8
}

func BaseCharacter(name string) Character {
	return Character{
		Name:  name,
		Level: 1,
		Exp:   0,
	}
}
