package server

import (
	"fmt"
	"net"

	"github.com/riversy/file-storage-server/service"
	"google.golang.org/grpc"
)

func RunGrpcServer() error {

	grpcService := fileService{
		storageFolder:      "",
		maxSize:            0,
		concurrentUploads:  0,
		concurrentRequests: 0,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9000))
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	service.RegisterFilesServiceServer(grpcServer, &grpcService)

	grpcServer.Serve(listener)

	return nil
}
