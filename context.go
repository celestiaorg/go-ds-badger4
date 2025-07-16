package badger4

import (
	"context"

	"github.com/ipfs/go-datastore"
)

type batchKey int

var batchKeyV batchKey

func WithBatch(ctx context.Context, batch datastore.Batch) context.Context {
	return context.WithValue(ctx, batchKeyV, batch)
}

func GetBatch(ctx context.Context) (datastore.Batch, bool) {
	batch, ok := ctx.Value(batchKeyV).(datastore.Batch)
	return batch, ok
}
