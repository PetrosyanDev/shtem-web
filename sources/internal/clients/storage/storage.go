package storageclient

import (
	"context"
	"io"
	"log"
	"shtem-web/sources/internal/configs"
	"shtem-web/sources/internal/core/domain"
	pb "shtem-web/sources/pkg/proto/storage"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	chunkSize       = 4 << 10 // 4Kb
	downloadTimeout = 5 * time.Minute
)

type StorageClient struct {
	ctx        context.Context
	cfg        *configs.Configs
	client     pb.ServiceFilesClient
	connection *grpc.ClientConn
}

func (s *StorageClient) Stop() error {
	err := s.connection.Close()
	if err == nil {
		log.Println("disconnected form Storage")
	}
	return err
}

func (s *StorageClient) ProvideFile(file *domain.File) domain.Error {
	md := metadata.New(map[string]string{s.cfg.Storage.HeaderKey: s.cfg.Storage.FileServiceKey})
	ctx, cancel := context.WithTimeout(s.ctx, downloadTimeout)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)
	req := FileToProvideFileRequest(file)
	stream, err := s.client.ProvideFile(ctx, req)
	if err != nil {
		return domain.NewError().SetError(err)
	}
ReadLoop:
	for {
		req, err := stream.Recv()
		if err != nil && err != io.EOF {
			return domain.NewError().SetError(err)
		}
		if _, err := file.Write(req.GetChunk()); err != nil {
			return domain.NewError().SetError(err)
		}
		if err == io.EOF {
			break ReadLoop
		}
		select {
		case <-stream.Context().Done():
			break ReadLoop
		default:
		}
	}
	return nil
}

func NewStorageClient(ctx context.Context, cfg *configs.Configs) (*StorageClient, error) {
	log.Println("connecting to Storage")
	log.Println(cfg.Storage)
	cx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	con, err := grpc.DialContext(cx, cfg.Storage.Addr, creds, grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	c := pb.NewServiceFilesClient(con)
	return &StorageClient{ctx, cfg, c, con}, nil
}
