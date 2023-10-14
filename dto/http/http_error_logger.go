package dtohttp

type errorLoggerData struct {
	Info       string
	RequestId  string
	StackTrace string
}

func NewErrorLoggerData(info string, requestId string, stackTrace string) LoggingData {
	return &errorLoggerData{
		Info:       info,
		RequestId:  requestId,
		StackTrace: stackTrace,
	}
}

func (e *errorLoggerData) GetParam() map[string]interface{} {
	return map[string]interface{}{
		"request_id":  e.RequestId,
		"stack_trace": e.StackTrace,
	}
}

func (e *errorLoggerData) GetInfo() string {
	return e.Info
}
