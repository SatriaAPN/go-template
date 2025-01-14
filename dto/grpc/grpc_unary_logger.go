package dtogrpc

import "time"

type requestGrpcLogger struct {
	fullMethod string
	requestId  string
	info       string
}

func NewRequestGrpcLogger(fullMethod string, requestId string, info string) *requestGrpcLogger {
	return &requestGrpcLogger{
		fullMethod: fullMethod,
		requestId:  requestId,
		info:       info,
	}
}

func (rgl *requestGrpcLogger) GetFields() map[string]interface{} {
	return map[string]interface{}{
		"full_method": rgl.fullMethod,
		"request_id":  rgl.requestId,
	}
}

func (rgl *requestGrpcLogger) GetInfo() string {
	return rgl.info
}

type responseGrpcLogger struct {
	fullMethod string
	requestId  string
	status     int
	info       string
	time       time.Duration
}

func NewResponseGrpcLogger(fullMethod string, requestId string, info string, status int, usedTime time.Duration) *responseGrpcLogger {
	return &responseGrpcLogger{
		fullMethod: fullMethod,
		requestId:  requestId,
		status:     status,
		info:       info,
		time:       usedTime,
	}
}

func (rgl *responseGrpcLogger) GetFields() map[string]interface{} {
	return map[string]interface{}{
		"full_method": rgl.fullMethod,
		"request_id":  rgl.requestId,
		"status":      rgl.status,
		"timePassed":  rgl.time,
	}
}

func (rgl *responseGrpcLogger) GetInfo() string {
	return rgl.info
}

type errorGrpcLogger struct {
	info       string
	requestId  string
	stackTrace string
}

func NewErrorLoggerData(info string, requestId string, stackTrace string) *errorGrpcLogger {
	return &errorGrpcLogger{
		info:       info,
		requestId:  requestId,
		stackTrace: stackTrace,
	}
}

func (e *errorGrpcLogger) GetFields() map[string]interface{} {
	return map[string]interface{}{
		"request_id":  e.requestId,
		"stack_trace": e.stackTrace,
	}
}

func (e *errorGrpcLogger) GetInfo() string {
	return e.info
}
