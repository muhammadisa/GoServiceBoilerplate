package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/muhammadisa/go-service-boilerplate/api/app/foobar"
	"github.com/muhammadisa/go-service-boilerplate/api/app/foobar/delivery/grpc/foobar_grpc"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	usecase foobar.Usecase
}

// NewFoobarServerGrpc intialize handler
func NewFoobarServerGrpc(gserver *grpc.Server, foobarUsecase foobar.Usecase) {
	foobarServer := &server{
		usecase: foobarUsecase,
	}
	foobar_grpc.RegisterFoobarHandlerServer(gserver, foobarServer)
	reflection.Register(gserver)
}

func (s *server) transformFoobarRPC(fBar *models.Foobar) *foobar_grpc.Foobar {
	if fBar == nil {
		return nil
	}
	UpdatedAt := &google_protobuf.Timestamp{
		Seconds: fBar.UpdatedAt.Unix(),
	}
	CraetedAt := &google_protobuf.Timestamp{
		Seconds: fBar.CreatedAt.Unix(),
	}

	res := &foobar_grpc.Foobar{
		ID:            fBar.ID.String(),
		FoobarContent: fBar.FoobarContent,
		UpdatedAt:     UpdatedAt,
		CreatedAt:     CraetedAt,
	}
	return res
}

func (s *server) transformFoobarData(fBar *foobar_grpc.Foobar) *models.Foobar {
	UpdatedAt := time.Unix(fBar.GetUpdatedAt().GetSeconds(), 0)
	CreatedAt := time.Unix(fBar.GetCreatedAt().GetSeconds(), 0)

	id, err := uuid.FromString(fBar.ID)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	res := &models.Foobar{
		ID:            id,
		FoobarContent: fBar.FoobarContent,
		UpdatedAt:     UpdatedAt,
		CreatedAt:     CreatedAt,
	}
	return res
}

func (s *server) GetFoobar(
	ctx context.Context,
	in *foobar_grpc.SingleRequest,
) (*foobar_grpc.Foobar, error) {

	uuid, err := uuid.FromString(in.ID)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fBar, err := s.usecase.GetByID(uuid)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if fBar == nil {
		return nil, fmt.Errorf("Foobar is Not Found")
	}

	res := s.transformFoobarRPC(fBar)
	return res, nil

}

func (s *server) GetListFoobar(
	ctx context.Context,
	in *foobar_grpc.FetchRequest,
) (*foobar_grpc.ListFoobar, error) {

	_, res, err := s.usecase.Fetch()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var arrFoobar []*foobar_grpc.Foobar
	for _, r := range *res {
		arrFoobar = append(arrFoobar, s.transformFoobarRPC(&r))
	}

	result := &foobar_grpc.ListFoobar{
		Foobars: arrFoobar,
	}

	return result, nil
}

func (s *server) UpdateFoobar(
	ctx context.Context,
	fBar *foobar_grpc.Foobar,
) (*foobar_grpc.Foobar, error) {

	fB := s.transformFoobarData(fBar)

	existed, err := s.usecase.GetByID(fB.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if existed == nil {
		return nil, fmt.Errorf("Foobar is Not Found")
	}

	err = s.usecase.Update(fB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	updatedFb := s.transformFoobarRPC(fB)

	return updatedFb, nil

}

func (s *server) Store(
	ctx context.Context,
	fBar *foobar_grpc.Foobar,
) (*foobar_grpc.Foobar, error) {

	fooBar := s.transformFoobarData(fBar)

	err := s.usecase.Store(fooBar)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res := s.transformFoobarRPC(fooBar)

	return res, nil

}

func (s *server) Delete(
	ctx context.Context,
	in *foobar_grpc.SingleRequest,
) (*foobar_grpc.DeleteResponse, error) {

	id, err := uuid.FromString(in.ID)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	err = s.usecase.Delete(id)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		return nil, err
	}

	res := &foobar_grpc.DeleteResponse{
		Status: "Successfully Deleted",
	}

	return res, nil
}
