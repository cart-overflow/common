package metadata

import "context"

type mdKey string

const metadataKey = mdKey("metadata")

func MetadataFromCtx(ctx context.Context) *Metadata {
	if md, ok := ctx.Value(metadataKey).(*Metadata); ok {
		return md
	}
	return &Metadata{}
}

func CtxWithMetadata(ctx context.Context, md *Metadata) context.Context {
	return context.WithValue(ctx, metadataKey, md)
}
