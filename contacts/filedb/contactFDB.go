package filedb

import (
	"contacts/models"
	"os"
)

type ContactFDB struct {
	DB interface{} // assume DB is file name
}

func (cdb *ContactFDB) Create(contact *models.Contact) (*models.Contact, error) {
	fs, err := os.Create(cdb.DB.(string))
	if err != nil {
		return nil, err
	}
	str, err := contact.ToJSONString()
	if err != nil {
		return nil, err
	}
	_, err = fs.WriteString(str)
	if err != nil {
		return nil, err
	}
	//os.WriteFile(cdb.DB.(string), contact.ToByte())
	return contact, nil
}

func (cdb *ContactFDB) UpdateBy(id string) (*models.Contact, error) {
	return nil, nil
}
func (cdb *ContactFDB) GetBy(id string) (*models.Contact, error) {
	return nil, nil
}
func (cdb *ContactFDB) DeleteBy(id string) (interface{}, error) {
	return nil, nil
}
