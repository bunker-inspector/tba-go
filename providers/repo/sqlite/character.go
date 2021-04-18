package sqlite

import (
	//"errors"
	"github.com/bunker-inspector/tba/domain"
	"gorm.io/gorm"
	"log"
)

type characterRepo struct {
	DB *gorm.DB
}

func newCharacterRepo(db *gorm.DB) characterRepo {
	return characterRepo{DB: db}
}

func (r characterRepo) GetByUserID(id int) *domain.Character {
	var c domain.Character
	result := r.DB.First(&c, "user_id = ?", id)
	log.Printf("Found: %+v", result)
	if result.Error != nil {
		return nil
	}
	return &c
}

func (r characterRepo) DeleteByUserID(id int) {
	result := r.DB.Where("user_id = ?", id).Delete(&domain.Character{})
	if result.Error != nil {
		log.Printf("Failed to delete character with ID %d.\n", id)
	}
}

func (r characterRepo) Put(c *domain.Character) {
	result := r.DB.Create(&c)
	if result.Error != nil {
		log.Printf("Failed to create character with attributes %+v.\n", *c)
	}
}
