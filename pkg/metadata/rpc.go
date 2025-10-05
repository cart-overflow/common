package metadata

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

func (b *Builder) RpcCtx(ctx context.Context) context.Context {
	return RpcCtxWithMetadata(ctx, b.md)
}

func RpcCtxWithMetadata(ctx context.Context, md *Metadata) context.Context {
	mdmap := map[string]string{}
	mdmap["user-id"] = md.UserId
	mdmap["role"] = md.Role
	mdmap["email"] = md.Email
	mdmap["name-bin"] = md.Name
	mdmap["register-date"] = strconv.FormatInt(md.RegisterTimestamp, 10)

	return metadata.NewOutgoingContext(ctx, metadata.New(mdmap))
}

func FromRpcCtx(ctx context.Context) *Metadata {
	md := Metadata{}

	ctxMd, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &md
	}

	userId := ctxMd.Get("user-id")
	if len(userId) > 0 {
		md.UserId = userId[0]
	}

	role := ctxMd.Get("role")
	if len(userId) > 0 {
		md.Role = role[0]
	}

	email := ctxMd.Get("email")
	if len(email) > 0 {
		md.Email = email[0]
	}

	name := ctxMd.Get("name-bin")
	if len(name) > 0 {
		md.Name = name[0]
	}

	regTimestamp := ctxMd.Get("register-date")
	if len(regTimestamp) > 0 {
		i, err := strconv.ParseInt(regTimestamp[0], 10, 64)
		if err == nil {
			md.RegisterTimestamp = i
		}
	}

	return &md
}
