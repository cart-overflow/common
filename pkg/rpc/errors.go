package rpc

import (
	"errors"

	"github.com/cart-overflow/common/pkg/core"

	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapServiceErr(err error) error {
	var aerr *core.Error
	ok := errors.As(err, &aerr)
	if !ok {
		return status.Error(codes.Internal, err.Error())
	}

	var st *status.Status
	if aerr.Code == core.ErrUnauthenticated {
		st = status.New(codes.Unauthenticated, aerr.Message)
	}
	if aerr.Code == core.ErrInvalidArgument {
		st = status.New(codes.InvalidArgument, aerr.Message)
	}
	if aerr.Code == core.ErrInternal {
		st = status.New(codes.Internal, aerr.Message)
	}
	if aerr.Code == core.ErrUnauthorized {
		st = status.New(codes.PermissionDenied, aerr.Message)
	}
	if st == nil {
		return status.Error(codes.Internal, err.Error())
	}

	withDs, err := st.WithDetails(
		&epb.ErrorInfo{
			Reason: aerr.Reason,
			Domain: "cartoverflow.user",
		},
		&epb.DebugInfo{
			Detail: aerr.Details,
		},
	)
	if err != nil {
		return st.Err()
	}

	return withDs.Err()
}
