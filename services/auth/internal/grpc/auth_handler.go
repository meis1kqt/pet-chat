package grpc

import (
	v1pb "chat-app/protos/gen/go/protos/sso/auth"
	"chat-app/services/auth/internal/grpc"
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context, email string, password string) (token string, userID int64, err error)
	RegisterNewUser(ctx context.Context, email string, password string) (int64, error)
}

type serverApi struct {
	v1pb.UnimplementedAuthServiceServer 
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	v1pb.RegisterAuthServiceServer(gRPC, &serverApi{auth: auth})
}




func (s *serverApi) Login(ctx context.Context, req *v1pb.LoginRequest)(*v1pb.LoginResponse, error) {
	if req.Email == "" {
		return &v1pb.LoginResponse{}, status.Error(codes.InvalidArgument, "email is empty")
	}
	if req.Password == "" {
		return &v1pb.LoginResponse{}, status.Error(codes.InvalidArgument, "password is empty")
	}

	token, userID, err := s.auth.Login(ctx, req.GetEmail(), req.Password)
	
	if err != nil {
		return &v1pb.LoginResponse{}, status.Error(codes.InvalidArgument, "invalid something")
	}

	return &v1pb.LoginResponse{
		Token: token,
		UserId: userID,
	}, nil
}


func (s *serverApi) Register(ctx context.Context, req *v1pb.RegisterRequest) (*v1pb.RegisterResponse, error) {
	if req.Email == "" {
		return &v1pb.RegisterResponse{}, status.Error(codes.InvalidArgument, "email is empty")
	}
	if req.Password == "" {
		return &v1pb.RegisterResponse{}, status.Error(codes.InvalidArgument, "password is empty")
	}

	id, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())

	if err != nil {
		return &v1pb.RegisterResponse{}, status.Error(codes.InvalidArgument, "lolhz 63 grpc")
	}

	return &v1pb.RegisterResponse{
		UserId: id,
	}, nil
}