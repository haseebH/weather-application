package repository

//go:generate mockery --name MessageQueue --inpackage --filename=message_queue_mock.go
type MessageQueue interface {
	Publish(message *User) error
	Close() error
}
