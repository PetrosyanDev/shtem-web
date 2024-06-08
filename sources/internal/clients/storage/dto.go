// HRACH_DEV © iMed Cloud Services, Inc.
package storageclient

import (
	"shtem-web/sources/internal/core/domain"
	pb "shtem-web/sources/pkg/proto/storage"
)

func FileToProvideFileRequest(file *domain.File) *pb.ProvideFileRequest {
	var policy pb.CachePolicy
	switch file.CachePolicy {
	case domain.FileCachePolicyMaximal:
		policy = pb.CachePolicy_CACHE_POLICY_MAXIMAL
	case domain.FileCachePolicyMedium:
		policy = pb.CachePolicy_CACHE_POLICY_MEDIUM
	case domain.FileCachePolicyMinimal:
		policy = pb.CachePolicy_CACHE_POLICY_MINIMAL
	default:
		policy = pb.CachePolicy_CACHE_POLICY_NEVER
	}
	return &pb.ProvideFileRequest{
		Key:         file.Key,
		Group:       file.OwnerID,
		CachePolicy: policy,
	}
}
