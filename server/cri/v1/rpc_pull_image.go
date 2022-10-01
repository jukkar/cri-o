package v1

import (
	"context"

	pb "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func (s *service) PullImage(
	ctx context.Context, req *pb.PullImageRequest,
) (*pb.PullImageResponse, error) {
	return s.server.PullImage(ctx, req)
}

func (s *service) PullImageWithProgress(req *pb.PullImageWithProgressRequest, srv pb.ImageService_PullImageWithProgressServer) error {
	return s.server.PullImageWithProgress(req, srv)
}
