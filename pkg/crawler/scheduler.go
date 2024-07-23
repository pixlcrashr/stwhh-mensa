package crawler

import (
	"context"
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Scheduler struct {
	crawler *Crawler
	storage *storage.Storage
	logger  *zap.Logger
}

func NewScheduler(
	c *Crawler,
	s *storage.Storage,
	l *zap.Logger,
) *Scheduler {
	return &Scheduler{
		crawler: c,
		storage: s,
		logger:  l.Named("scheduler"),
	}
}

func (s *Scheduler) StartAndBlock(
	interval time.Duration,
) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		wg.Add(1)
		defer wg.Done()

		s.process(ctx)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.process(ctx)
			}
		}
	}()

	s.logger.Info("started crawler scheduler")

	<-c
	cancel()
	wg.Wait()

	s.logger.Info("SIGINT signaled... stopped crawler scheduler gracefully")

	return nil
}

func (s *Scheduler) process(ctx context.Context) {
	s.logger.Debug("crawling data")

	days, err := s.crawler.Crawl(ctx)

	if err != nil {
		s.logger.Error("could not crawl data", zap.Error(err))
		return
	}

	s.logger.Debug("finished crawling data")

	s.logger.Debug("adding crawled data to db")

	for _, day := range days {
		if err := s.storage.AddDay(ctx, day); err != nil {
			s.logger.Info("could not add data", zap.Error(err))
		}
	}

	s.logger.Debug("added crawled data to db")
}
