package sqlite

import (
	"github.com/bunker-inspector/tba/domain"
	"log"
)

func (r *repo) GetCharacterByUserID(id int) *domain.Character {
	var c domain.Character
	result := r.First(&c, "user_id = ?", id)
	log.Printf("Found: %+v", result)
	if result.Error != nil {
		return nil
	}
	return &c
}

func (r *repo) DeleteCharacterByUserID(id int) {
	result := r.Delete(&domain.Character{UserID: id})
	if result.Error != nil {
		log.Printf("Failed to delete character with ID %d.\n", id)
	}
}

func (r *repo) SaveCharacter(c *domain.Character) {
	result := r.Create(&c)
	if result.Error != nil {
		log.Printf("Failed to create character with attributes %+v.\n", *c)
	}
}
