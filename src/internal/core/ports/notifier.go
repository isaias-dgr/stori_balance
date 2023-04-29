package ports

type Notifier interface {
	Send_Messages(subject, telephone string) error
}
