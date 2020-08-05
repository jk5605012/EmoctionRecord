package e

var MsgFlags = map[int]string{
	SUCCESS:         "ok",
	SERVER_ERROR:    "fail",
	PARAMETER_ERROR: "請求參數錯誤",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[SERVER_ERROR]
}
