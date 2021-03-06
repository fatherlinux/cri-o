package server

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// StopPodSandbox stops the sandbox. If there are any running containers in the
// sandbox, they should be force terminated.
func (s *Server) StopPodSandbox(ctx context.Context, req *pb.StopPodSandboxRequest) (resp *pb.StopPodSandboxResponse, err error) {
	// platform dependent call
	return s.stopPodSandbox(ctx, req)
}

// stopAllPodSandboxes removes all pod sandboxes
func (s *Server) stopAllPodSandboxes(ctx context.Context) {
	logrus.Debugf("stopAllPodSandboxes")
	for _, sb := range s.ContainerServer.ListSandboxes() {
		pod := &pb.StopPodSandboxRequest{
			PodSandboxId: sb.ID(),
		}
		if _, err := s.StopPodSandbox(ctx, pod); err != nil {
			logrus.Warnf("could not StopPodSandbox %s: %v", sb.ID(), err)
		}
	}
}
