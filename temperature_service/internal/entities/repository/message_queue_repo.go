package repository

type MessageQueue interface {
	Consume() error
	Close() error
}
