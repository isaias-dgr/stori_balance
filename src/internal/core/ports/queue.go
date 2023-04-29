package ports

type Queue interface {
	Send(message *string) error
}
