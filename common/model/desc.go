package model

var m = map[int]string{
	SUCCESS:        "操作成功",
	ERROR:          "操作失败",
	INVALID_PARAMS: "请求参数错误",

	UnauthorizedAuthFail: "未登录/权限不足",

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token过期，请重新登录",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "无效的token，请重新登录",
}

func GetDesc(code int) string {
	if desc, ok := m[code]; ok {
		return desc
	}

	return m[ERROR]
}
