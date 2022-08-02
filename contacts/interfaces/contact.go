package interfaces

import (
	"contacts/models"
)

type IContact interface {
	Create(contact *models.Contact) (*models.Contact, error)
	UpdateBy(id string) (*models.Contact, error)
	GetBy(id string) (*models.Contact, error)
	DeleteBy(id string) (interface{}, error)
}

// 1- What is your approach about Unit Testing?
// 2- gomock or testify
// 3- How do you store you passwords?
// 4- kubernetes secretes are not actually secrets . They are just base64 encoded
// 5- Does your application work with out kubernetes?
// 6- go-client/kube-client --> to access configmaps/secrets
