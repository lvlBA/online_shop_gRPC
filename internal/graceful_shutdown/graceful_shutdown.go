package gracefulshutdown

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/lvlBA/online_shop/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StopHandle func()

type StopWithErrorHandle func() error

type GracefulShutDown struct {
	ch     chan os.Signal
	log    logger.Logger
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
	stop   []StopHandle
}

type Config struct {
	Ctx  context.Context
	Log  logger.Logger
	Stop []StopHandle
}

func New(cfg *Config) *GracefulShutDown {
	ch := make(chan os.Signal)
	ctx, cancel := context.WithCancel(cfg.Ctx)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	return &GracefulShutDown{
		ch:     ch,
		log:    cfg.Log,
		ctx:    ctx,
		cancel: cancel,
		wg:     new(sync.WaitGroup),
		stop:   cfg.Stop,
	}
}

func (s *GracefulShutDown) Observe() {
	sig := <-s.ch
	s.log.Info(s.ctx, "received signal %s", sig)
	s.cancel()
	s.wg.Wait()
	for i := range s.stop {
		s.stop[i]()
	}
}

func (s *GracefulShutDown) AddStop(handle StopHandle) {
	s.stop = append(s.stop, handle)
}

func (s *GracefulShutDown) GetContext() context.Context {
	return s.ctx
}

func (s *GracefulShutDown) GetWG() *sync.WaitGroup {
	return s.wg
}

func (s *GracefulShutDown) GrpcInterceptor(
	ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	s.wg.Add(1)
	defer s.wg.Done()

	if err := s.ctx.Err(); err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	return handler(ctx, req)
}
