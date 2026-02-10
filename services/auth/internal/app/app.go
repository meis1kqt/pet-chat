package app

import (
	authgrpc "chat-app/services/auth/internal/grpc"
	"chat-app/services/auth/internal/service"
	"chat-app/services/auth/internal/storage/postgres"
	"fmt"
	"net"

	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, gRPCPort int, dsnString string, jwtSecret string, tokenTTL time.Duration) *App {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsnString)

	if err != nil {
		panic(err)
	}

	err = pool.Ping(ctx)

	if err != nil {
		panic(err)
	}

	storage := postgres.NewStorage(pool)

	authService := service.NewAuthService(log, storage, tokenTTL, jwtSecret)

	gRPCSERVER := grpc.NewServer()

	authgrpc.Register(gRPCSERVER, authService)

	log.Info("gRPC server initialized")

	return &App{
		log:        log,
		gRPCServer: gRPCSERVER,
		port:       gRPCPort,
	}

}

func (a *App) Run() error {

	a.log.Info("statring grpc server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		a.log.Error("listen error", "error", err)
		return err
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		a.log.Error("some trouble", "error", err)
		return err
	}

	return nil
}



func (a *App) Stop() {
	a.log.Info("server stoping")

	a.gRPCServer.GracefulStop()
}