package grpc

import (
	"contacts/interfaces"
	"contacts/models"
	pb "contacts/protos"
	"context"
	"fmt"
	"time"
)

type ContactServer struct {
	pb.UnimplementedContactServer
	IContact interfaces.IContact
}

func (s *ContactServer) Create(ctx context.Context, in *pb.ContactMessageRequest) (*pb.GeneralResponse, error) {
	contact := new(models.Contact)
	contact.Email = in.Email
	contact.Name = in.Name
	contact.ContactNo = in.ContactNo
	contact.MoreInfo = in.MoreInfo
	contact.Address = in.Address
	contact.Status = "create"
	contact.LastModified = fmt.Sprint(time.Now().Unix())
	if err := contact.Validate(); err != nil {
		return &pb.GeneralResponse{Code: 400, Message: "Bad request", InnerMessage: err.Error()}, nil
	}
	c, err := s.IContact.Create(contact)
	if err != nil {
		return &pb.GeneralResponse{Code: 400, Message: "Bad request", InnerMessage: err.Error()}, nil
	}
	jsonMessage, _ := c.ToJSONString()
	return &pb.GeneralResponse{Code: 201, Message: "Contact successfully created", InnerMessage: jsonMessage}, nil
}

func (s *ContactServer) GetBy(ctx context.Context, in *pb.ContactByIDRequest) (*pb.ContactMessage, error) {
	contact, err := s.IContact.GetBy(in.ID)
	if err != nil {
		return nil, err
	}
	return &pb.ContactMessage{ID: int64(contact.ID), Name: contact.Name, Email: contact.Email, ContactNo: contact.ContactNo, Address: contact.Address, MoreInfo: contact.MoreInfo, Status: contact.Status, LastModified: contact.LastModified}, nil
}

func (s *ContactServer) UpdateBy(ctx context.Context, in *pb.ContactUpdateRequest) (*pb.ContactMessage, error) {

	fmt.Println(in.ID, in.Data)
	contact, err := s.IContact.UpdateBy(in.ID, in.Data.AsMap())
	if err != nil {
		return nil, err
	}
	return &pb.ContactMessage{ID: int64(contact.ID), Name: contact.Name, Email: contact.Email, ContactNo: contact.ContactNo, Address: contact.Address, MoreInfo: contact.MoreInfo, Status: contact.Status, LastModified: contact.LastModified}, nil
}

func (s *ContactServer) DeleteBy(ctx context.Context, in *pb.ContactByIDRequest) (*pb.GeneralResponse, error) {
	result, err := s.IContact.DeleteBy(in.ID)
	if err != nil {
		return &pb.GeneralResponse{Code: 400, Message: "Bad request", InnerMessage: err.Error()}, nil
	}
	return &pb.GeneralResponse{Code: 202, Message: "successfully deleted", InnerMessage: fmt.Sprint(result.(int64))}, nil
}
