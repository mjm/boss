package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mjm/boss/api"
	"github.com/mjm/boss/cfg"
)

type Server struct {
	Project        *cfg.Project
	ConfigFilePath string

	grpcServer *grpc.Server
	tasks      map[string]*taskWorker
	wg         sync.WaitGroup
}

func (s *Server) Run() error {
	s.createMissingTaskWorkers()
	defer s.stopTaskWorkers()

	// TODO make this a flag
	socketName := fmt.Sprintf("/tmp/boss.%s.sock", s.Project.Name)
	lis, err := net.Listen("unix", socketName)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()
	api.RegisterBossServer(s.grpcServer, s)

	return s.grpcServer.Serve(lis)
}

func (s *Server) Shutdown() {
	s.grpcServer.GracefulStop()
}

func (s *Server) createMissingTaskWorkers() {
	if s.tasks == nil {
		s.tasks = map[string]*taskWorker{}
	}

	for _, t := range s.Project.Tasks {
		if _, ok := s.tasks[t.Name]; ok {
			continue
		}

		tw := &taskWorker{
			Task: t,
			Ch:   make(chan interface{}),
			Done: make(chan struct{}),
		}
		s.wg.Add(1)
		go func() {
			tw.Run()
			s.wg.Done()
		}()
		s.tasks[t.Name] = tw
	}
}

func (s *Server) stopTaskWorkers() {
	for _, tw := range s.tasks {
		tw.Stop()
	}
	s.tasks = nil
	s.wg.Wait()
}

var _ api.BossServer = (*Server)(nil)

func (s *Server) StartTasks(ctx context.Context, req *api.StartTasksRequest) (*api.StartTasksResponse, error) {
	for _, taskName := range req.GetNames() {
		tw, ok := s.tasks[taskName]
		if !ok {
			return nil, status.Errorf(codes.NotFound, "no task named %s", taskName)
		}

		tw.Ch <- startTaskMsg{}
	}

	return &api.StartTasksResponse{}, nil
}
