package dtohttp

type httpRequestLogging struct {
	Endpoint  string
	Method    string
	RequestId string
	Info      string
}

func NewHttpRequestLogging(endpoint string, method string, requestId string, info string) *httpRequestLogging {
	return &httpRequestLogging{
		Endpoint:  endpoint,
		Method:    method,
		RequestId: requestId,
		Info:      info,
	}
}

func (h *httpRequestLogging) GetFields() map[string]interface{} {
	return map[string]interface{}{
		"endpoint":   h.Endpoint,
		"method":     h.Method,
		"request_id": h.RequestId,
	}
}

func (h *httpRequestLogging) GetInfo() string {
	return h.Info
}

type httpResponseLogging struct {
	Endpoint  string
	Method    string
	RequestId string
	Info      string
	Status    int
}

func NewHttpResponseLogging(endpoint string, method string, requestId string, info string, status int) *httpResponseLogging {
	return &httpResponseLogging{
		Endpoint:  endpoint,
		Method:    method,
		RequestId: requestId,
		Info:      info,
		Status:    status,
	}
}

func (h *httpResponseLogging) GetFields() map[string]interface{} {
	return map[string]interface{}{
		"endpoint":   h.Endpoint,
		"method":     h.Method,
		"request_id": h.RequestId,
		"status":     h.Status,
	}
}

func (h *httpResponseLogging) GetInfo() string {
	return h.Info
}
