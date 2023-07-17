package except

import (
	// "net/http"
	"google.golang.org/grpc/codes"
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

func Error2StatusCode(err error)(statusCode codes.Code){
	switch err.(type){
		case *InvalidArgument:
			return codes.InvalidArgument

		case *SystemError:
			return 	codes.Internal
			
		default:
			return codes.InvalidArgument
	}
}