package ports

import "github.com/isaias-dgr/story-balance/src/internal/core/domain"

type Mailer interface {
	Send(dest string, acc domain.Account) error
}
