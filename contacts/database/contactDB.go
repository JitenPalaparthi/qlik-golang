package database

import (
	"contacts/models"

	"gorm.io/gorm"
)

type ContactDB struct {
	DB interface{}
}

func (cdb *ContactDB) Create(contact *models.Contact) (*models.Contact, error) {
	// Need to check this implementation. What if second time it is called
	if err := cdb.DB.(*gorm.DB).AutoMigrate(&models.Contact{}); err != nil {
		return nil, err
	}
	tx := cdb.DB.(*gorm.DB).Create(contact)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}

func (cdb *ContactDB) UpdateBy(id string) (*models.Contact, error) {
	return nil, nil
}
func (cdb *ContactDB) GetBy(id string) (*models.Contact, error) {
	return nil, nil
}
func (cdb *ContactDB) DeleteBy(id string) (interface{}, error) {
	return nil, nil
}
