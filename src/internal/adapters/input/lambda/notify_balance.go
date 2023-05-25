package adapter

import (
	"context"

	log "github.com/sirupsen/logrus"

	dto "github.com/isaias-dgr/stori-balance/src/internal/adapters/dto/input"
	usecase "github.com/isaias-dgr/stori-balance/src/internal/core/use_case"
)

type funcNotify struct {
	l        *log.Logger
	use_case *usecase.NotifyBalance
}

func NewfuncNotify(l *log.Logger, use_case *usecase.NotifyBalance) *funcNotify {
	return &funcNotify{
		l:        l,
		use_case: use_case,
	}
}

func (fs funcNotify) Handler(
	ctx context.Context, event dto.EventNotify) (string, error) {
	fs.l.Infof("🧐 %+v", event)
	record := event.Records[0]
	file_ingest := record.Body
	err := fs.use_case.Execute(file_ingest)
	if err != nil {
		fs.l.Errorf("Handler 😭 %s ", err.Error())
		return "", err
	}
	return "Handler 😎", nil
}
