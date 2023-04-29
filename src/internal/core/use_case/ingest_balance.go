package usecase

import (
	log "github.com/sirupsen/logrus"

	"github.com/isaias-dgr/story-balance/src/internal/core/ports"
)

type IngestUseCase struct {
	log     *log.Logger
	storage ports.Storage
	queue   ports.Queue
}

func NewIngestUseCase(l *log.Logger, s ports.Storage, q ports.Queue) *IngestUseCase {
	return &IngestUseCase{
		log:     l,
		storage: s,
		queue:   q,
	}
}

func (is IngestUseCase) Execute() (err error) {
	is.log.Info("👨🏿 Ingesting values on queue")
	files_logs, err := is.storage.GetListFiles("")
	if err != nil {
		is.log.Errorf("fail storage %s", err.Error())
		return err
	}
	is.log.Infof("✉️ %d", len(files_logs))
	for _, file := range files_logs {
		is.log.Infof("✉️ %s", file)
		err := is.queue.Send(&file)
		if err != nil {
			is.log.Errorf("fail storage %s", err.Error())
			return err
		}
	}
	return nil
}
