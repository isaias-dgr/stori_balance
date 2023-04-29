package domain

type User struct {
	Id                 string
	Email              string
	Tel                string
	Notification_email bool
	Notification_sms   bool
}
