package database

import (
	"contacts/models"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

// Note: Use gomock only if unit testing cannot be performed on external components
func TestCreateContact(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	icontact := NewMockIContact(ctrl)
	contact := new(models.Contact)
	lmt := fmt.Sprint(time.Now().Unix())
	contact.Name = "Demo"
	contact.Email = "Demo@Gmail.Com"
	contact.ContactNo = "111111111"
	contact.MoreInfo = "Demo purpose"
	contact.Status = "created"
	contact.LastModified = lmt

	econtact := new(models.Contact)
	econtact.ID = 0
	econtact.Name = "Demo"
	econtact.Email = "Demo@Gmail.Com"
	econtact.ContactNo = "111111111"
	econtact.MoreInfo = "Demo purpose"
	econtact.Status = "created"
	econtact.LastModified = lmt
	icontact.EXPECT().Create(contact).Return(econtact, nil).AnyTimes()

}
