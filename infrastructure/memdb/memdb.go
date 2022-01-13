package memdb

import "context"

type MemDB struct{}

func (m MemDB) Insert(_ context.Context, _ interface{}) error { return nil }
