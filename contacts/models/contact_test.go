package models

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestValidateContact(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// cmodel := NewMockIContactModel(ctrl)
	// //cmodel.EXPECT().Validate()

	// cmodel := &Contact{Name: "Demo"}

}
