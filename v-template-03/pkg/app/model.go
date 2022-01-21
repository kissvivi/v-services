package app

type Response struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"status_msg"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
}

func (r Response) ReturnError(code int) interface{} {
	//todo:错误日志
	r.StatusCode = code
	r.Success = false
	return r
}

func (r Response) ReturnOK() interface{} {
	r.StatusCode = 200
	r.Success = true
	return r
}

type PageData struct {
	List     interface{} `json:"list"`
	Count    int         `json:"count"`
	PageNo   int         `json:"page_no"`
	PageSize int         `json:"page_size"`
}

type PageResponse struct {
	StatusCode int      `json:"status_code"`
	Msg        string   `json:"status_msg"`
	Success    bool     `json:"success"`
	Data       PageData `json:"data"`
}

func (r PageResponse) ReturnOK() interface{} {
	r.StatusCode = 200
	r.Success = true
	return r
}

func (r PageResponse) ReturnError(code int) interface{} {
	r.StatusCode = code
	r.Success = false
	return r
}
