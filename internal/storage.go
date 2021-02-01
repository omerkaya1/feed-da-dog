package internal

import "context"

type Store interface {
	Create(ctx context.Context) (string, error)
	Read(ctx context.Context) (interface{}, error)
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}

type DummyDB struct {
}

func NewDummyDB() *DummyDB {
	return &DummyDB{}
}

func (d *DummyDB) Create(ctx context.Context) (string, error) {
	return "", nil
}

func (d *DummyDB) Read(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (d *DummyDB) Update(ctx context.Context) error {
	return nil
}

func (d *DummyDB) Delete(ctx context.Context) error {
	return nil
}
