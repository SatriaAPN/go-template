package dtogrpc

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
}

func NewResponseGrpcLogger(fullMethod string, requestId string, info string, status int) *responseGrpcLogger {
	return &responseGrpcLogger{
		fullMethod: fullMethod,
		requestId:  requestId,
		status:     status,
		info:       info,
	}
}

func (rgl *responseGrpcLogger) GetFields() map[string]interface{} {
	return map[string]interface{}{
		"full_method": rgl.fullMethod,
		"request_id":  rgl.requestId,
		"status":      rgl.status,
	}
}

func (rgl *responseGrpcLogger) GetInfo() string {
	return rgl.info
}
