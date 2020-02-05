package logger

// Storage storage
type Storage struct{}

// NewStorage new storage
func NewStorage() *Storage {
	return &Storage{}
}

// File storage log to file
func (s *Storage) File() {

}

// DB storage log to db
func (s *Storage) DB() {

}

// Cache storage log to cache
func (s *Storage) Cache() {

}
