package rpc

import (
	"context"
	"fmt"
	"log"
	"time"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/muhammadisa/restful-api-boilerplate/api/foobar"
	"github.com/muhammadisa/restful-api-boilerplate/api/foobar/delivery/rpc/foobar_grpc"
	"github.com/muhammadisa/restful-api-boilerplate/api/models"
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
		ID:            fBar.ID,
		FoobarContent: fBar.FoobarContent,
		UpdatedAt:     UpdatedAt,
		CreatedAt:     CraetedAt,
	}
	return res
}

func (s *server) transformFoobarData(fBar *foobar_grpc.Foobar) *models.Foobar {
	UpdatedAt := time.Unix(fBar.GetUpdatedAt().GetSeconds(), 0)
	CreatedAt := time.Unix(fBar.GetCreatedAt().GetSeconds(), 0)
	res := &models.Foobar{
		ID:            fBar.ID,
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

	id := uint64(0)
	if in != nil {
		id = in.Id
	}

	fBar, err := s.usecase.GetByID(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if fBar == nil {
		return nil, fmt.Errorf("Foobar is nil")
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
	return nil, nil
}

func (s *server) Store(
	ctx context.Context,
	fBar *foobar_grpc.Foobar,
) (*foobar_grpc.Foobar, error) {
	return nil, nil
}

func (s *server) Delete(
	ctx context.Context,
	in *foobar_grpc.SingleRequest,
) (*foobar_grpc.DeleteResponse, error) {
	return nil, nil
}
