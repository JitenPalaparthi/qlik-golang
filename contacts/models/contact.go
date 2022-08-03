package models

import "encoding/json"

type Contact struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	ContactNo    string `json:"contactNo" gorm:"column:contactNo"`
	MoreInfo     string `json:"moreInfo" gorm:"column:moreInfo"`
	Status       string `json:"status"`
	LastModified string `json:"lastModified" gorm:"column:lastModified"`
}

// if Name is name then it is considered as unexported field

type IContactModel interface {
	Validate() error
	ToByte() ([]byte, error)
	ToJSONString() (string, error)
}

func (c *Contact) Validate() error {
	if c.Name == "" {
		return ErrInvalidName
	}

	if c.Email == "" {
		return ErrInvalidEmail
	}

	if c.ContactNo == "" {
		return ErrInvalidContactNo
	}

	return nil
}

func (c *Contact) ToByte() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Contact) ToJSONString() (string, error) {
	buf, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
