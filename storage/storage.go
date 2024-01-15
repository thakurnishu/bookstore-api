package storage

type Storage interface {
	AddBook() error
	GetBook() error
}

func (db *PostgresStore) AddBook() error {
	return nil
}

func (db *PostgresStore) GetBook() error {
	return nil
}
