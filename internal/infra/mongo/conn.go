package mongo

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}
