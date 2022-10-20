package web

type response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type Pagination struct {
	Data  interface{} `json:"data" binding:"required"`
	Count int         `json:"count" binding:"required"`
}

func CustomResponse(statusCode int, data interface{}, err string) response {
	if statusCode < 300 {
		return response{statusCode, data, ""}
	}
	return response{statusCode, nil, err}
}

func PaginationResponse(statusCode int, data *Pagination, err string) response {
	if statusCode < 300 {
		return response{statusCode, data, ""}
	}
	return response{statusCode, nil, err}
}
