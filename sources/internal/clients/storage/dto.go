// HRACH_DEV © iMed Cloud Services, Inc.
package storageclient

import (
	"shtem-web/sources/internal/core/domain"
	pb "shtem-web/sources/pkg/proto/storage"
)

func FileToProvideFileRequest(file *domain.File) *pb.ProvideFileRequest {

	return &pb.ProvideFileRequest{
		Key:   file.Key,
		Group: file.OwnerID,
	}
}
