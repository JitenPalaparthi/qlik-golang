//go:generate mockgen.exe -source=../interfaces/contact.go -destination=contactDB_mock.go -package=database
package database

import (
	"contacts/models"
	"strconv"

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

func (cdb *ContactDB) UpdateBy(id string, data map[string]interface{}) (*models.Contact, error) {
	contact := new(models.Contact)
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	contact.ID = _id
	tx := cdb.DB.(*gorm.DB).Model(contact).Updates(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	contact, err = cdb.GetBy(id)
	if err != nil {
		return nil, err
	}

	return contact, nil
}
func (cdb *ContactDB) GetBy(id string) (*models.Contact, error) {
	contact := new(models.Contact)
	tx := cdb.DB.(*gorm.DB).First(contact, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}

func (cdb *ContactDB) DeleteBy(id string) (interface{}, error) {
	tx := cdb.DB.(*gorm.DB).Delete(&models.Contact{}, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx.RowsAffected, nil
}

func (cdb *ContactDB) GetAllByStatus(status string) ([]models.Contact, error) {
	contacts := make([]models.Contact, 0)

	tx := cdb.DB.(*gorm.DB).Where(&models.Contact{Status: status}).First(&contacts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contacts, nil
}
