package interceptor

import (
	"context"

	"github.com/go-errors/errors"

	dto "go-template/dto/general"
	dtogrpc "go-template/dto/grpc"
	errorapp "go-template/share/general/error"
	"go-template/share/general/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorHandling(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)

	if err != nil {
		stackTrace := "unhandled"

		switch e := err.(type) {
		case *errors.Error:
			err = errorCase(err)
			stackTrace = e.ErrorStack()
		}

		egl := dtogrpc.NewErrorLoggerData("error-grpc", ctx.Value(dto.RequestIdKey).(string), stackTrace)
		util.GetLogger().Errorf(egl)
	}

	return resp, err
}

func errorCase(err error) error {
	if status.Code(err) != codes.OK {
		return err
	} else {
		switch {
		case errors.Is(err, errorapp.ErrEmailAlreadyExist):
			return status.Error(codes.AlreadyExists, err.Error())
		case errors.Is(err, errorapp.ErrEmailIsNotValid):
			return status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, errorapp.ErrEmailNotFound):
			return status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, errorapp.ErrForgetPasswordTokenLength):
			return status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, errorapp.ErrMaximumPasswordLength):
			return status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, errorapp.ErrMinimumPasswordLength):
			return status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, errorapp.ErrResetCodeNotFound):
			return status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, errorapp.ErrWrongPassword):
			return status.Error(codes.InvalidArgument, err.Error())
		default:
			return status.Error(codes.Internal, "something wrong happened")
		}
	}
}
