package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func CustomResponse(statusCode int, data interface{}, err string) Response {
	if statusCode < 300 {
		return Response{statusCode, data, ""}
	}
	return Response{statusCode, nil, err}
}
