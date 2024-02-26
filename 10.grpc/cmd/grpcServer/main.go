package main

import (
	"database/sql"
	"net"

	"github.com/CaiqueRibeiro/grpc/internal/database"
	"github.com/CaiqueRibeiro/grpc/internal/pb"
	"github.com/CaiqueRibeiro/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(categoryDb)

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
