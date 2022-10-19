package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type Pagination struct {
	Data  interface{} `json:"data" binding:"required"`
	Count int         `json:"count" binding:"required"`
}

func CustomResponse(statusCode int, data interface{}, err string) Response {
	if statusCode < 300 {
		return Response{statusCode, data, ""}
	}
	return Response{statusCode, nil, err}
}

func PaginationResponse(statusCode int, data *Pagination, err string) Response {
	if statusCode < 300 {
		return Response{statusCode, data, ""}
	}
	return Response{statusCode, nil, err}
}
