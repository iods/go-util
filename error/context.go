package error

//type ErrorContext struct {
//	ctx context.Context
//	err error
//}
//
//func (err ErrorContext) Error() string {
//	return err.err.Error()
//}
//
//func (err ErrorContext) Ctx() context.Context {
//	return err.ctx
//}
//
//// Errorf will wrap the error in the errors.Errorf package.
//func Errorf(ctx context.Context, format string, args ...interface{}) error {
//	return ErrorContext{
//		ctx: ctx,
//		err: errors.Errorf(format, args),
//	}
//}
//
//// New will extend the traditional function in the Go errgo package.
//func New(ctx context.Context, message string) error {
//	return ErrorContext{
//		ctx: ctx,
//		err: errgo.New(message),
//	}
//}
//
//// Newf will be for managing the format of an error?
//func Newf(ctx context.Context, format string, args ...interface{}) error {
//	return ErrorContext{
//		ctx: ctx,
//		err: errgo.Newf(format, args...),
//	}
//}
//
//// Notef will format an existing error
//func Notef(ctx context.Context, err error, format string, args ...interface{}) error {
//	return ErrorContext{
//		ctx: ctx,
//		err: errgo.Notef(err, format, args...),
//	}
//}
//
//// NoteMask protects an underlying error by rendering something else when thrown.
//func NoteMask(ctx context.Context, err error, message string) error {
//	return ErrorContext{
//		ctx: ctx,
//		err: errgo.NoteMask(err, message),
//	}
//}
//
//// Wrap can support wrapping error messages for simpler control.
//func Wrap(ctx context.Context, err error, message string) error {
//	return ErrorContext{
//		ctx: ctx,
//		err: errors.Wrap(err, message),
//	}
//}
//
//func Wrapf(ctx context.Context, err error, format string, args ...interface{}) error {
//	return ErrorContext{
//		ctx: ctx,
//		err: errors.Wrapf(err, format, args...),
//	}
//}
//
//// UnwrapContext will unwrap all wrapped errors from err to get the deepest available context.
//func UnwrapContext(ctx context.Context, err error) context.Context {
//	var previousCtx context.Context
//
//	type causer interface {
//		Cause() error
//	}
//
//	for err != nil {
//		errgoErr, ok := err.(*errgo.Err)
//		if ok {
//			err = errgoErr.Underlying()
//			continue
//		}
//
//		cause, ok := err.(causer)
//		if ok {
//			err = cause.Cause()
//			continue
//		}
//
//		errContext, ok := err.(ErrorContext)
//		if ok {
//			err = errContext.err
//			previousCtx = errContext.Ctx()
//			continue
//		}
//
//		break
//	}
//
//	if previousCtx == nil {
//		return ctx
//	}
//
//	return previousCtx
//}
