package models

type Contact struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	ContactNo    string `json:"contactNo"`
	MoreInfo     string `json:"moreInfo"`
	Status       string `json:"status"`
	LastModified string `json:"lastModified"`
}

// if Name is name then it is considered as unexported field

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
