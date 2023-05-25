package adapter

import (
	"context"

	log "github.com/sirupsen/logrus"

	dto "github.com/isaias-dgr/stori-balance/src/internal/adapters/dto/input"
	usecase "github.com/isaias-dgr/stori-balance/src/internal/core/use_case"
)

type funcIngest struct {
	l        *log.Logger
	use_case *usecase.IngestUseCase
}

func NewfuncIngest(l *log.Logger, use_case *usecase.IngestUseCase) *funcIngest {
	return &funcIngest{
		l:        l,
		use_case: use_case,
	}
}

func (fs funcIngest) Handler(
	ctx context.Context, event dto.EventIngest) (string, error) {
	fs.l.Infof("🧐 %s", event.Name)
	err := fs.use_case.Execute()
	if err != nil {
		fs.l.Errorf("Handler 😭 %s ", err.Error())
		return "", err
	}
	return "Handler 😎", nil
}
