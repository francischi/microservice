package except

import (
	"net/http"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InvalidArgument struct{
	Message string
} 
func (e *InvalidArgument) Error() string {
    return e.Message
}

type SystemError struct{
	Message string
}
func (e *SystemError) Error() string {
    return e.Message
}

func Error2StatusCode(err error)(statusCode int , errMsg string){
	if e,ok := status.FromError(err);ok{
		switch e.Code() {
			case codes.InvalidArgument:
				return http.StatusBadRequest , e.Proto().Message

			case codes.Internal:
				return http.StatusInternalServerError , e.Proto().Message

			default:
				return http.StatusBadRequest , e.Proto().Message

		}
	}
	return http.StatusBadRequest , "bad request"
}