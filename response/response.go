package response

const (
	SUCCESS      = 200
	ERROR        = 500
	FAIL         = 400
	NOTFOUND     = 404
	UNAUTHORIZED = 401
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Ok(data any) Response {
	return Response{
		Code: SUCCESS,
		Msg:  "成功",
		Data: data,
	}
}

func Error(msg string) Response {
	return Response{
		Code: ERROR,
		Msg:  msg,
		Data: "",
	}
}

func ErrorWithCode(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: "",
	}
}
