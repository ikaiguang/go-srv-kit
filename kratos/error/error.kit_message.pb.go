package errorpkg

import (
	stdhttp "net/http"
)

// ErrMessage ...
func ErrMessage(key string) string {
	if v, ok := errMessage[key]; ok {
		return v
	}
	return ""
}

var (
	_          = stdhttp.StatusOK
	errMessage = map[string]string{
		"UNKNOWN":                            "未知错误",
		"REQUEST_FAILED":                     "请求失败",
		"RECORD_NOT_FOUND":                   "请求的内容未找到或已删除",
		"RECORD_ALREADY_EXISTS":              "数据已存在",
		"NETWORK_ERROR":                      "网络开小差了，请稍后再试",
		"NETWORK_TIMEOUT":                    "请求超时",
		"CONNECTION":                         "链接错误",
		"UNIMPLEMENTED":                      "未实现",
		"REQUEST_NOT_SUPPORT":                "暂不支持此次请求",
		"THIRD_PARTY_SERVICE_INTERNAL_ERROR": "第三方服务错误",
		"THIRD_PARTY_SERVICE_INVALID_CODE":   "第三方服务响应结果有误",
		"DB":                                 "DB错误",
		"MYSQL":                              "MYSQL错误",
		"MONGO":                              "MONGO错误",
		"CACHE":                              "CACHE错误",
		"REDIS":                              "REDIS错误",
		"MQ":                                 "MQ错误",
		"RABBIT_MQ":                          "RABBIT_MQ错误",
		"KAFKA":                              "KAFKA错误",
		"PANIC":                              "服务内部错误：PANIC",
		"FATAL":                              "服务内部错误：FATAL",

		"CONTINUE":            "请继续请求",
		"SWITCHING_PROTOCOLS": "请切换协议",
		"PROCESSING":          "将继续执行请求",

		"OK":                     "请求成功",
		"CREATED":                "请求已被接受，等待资源响应",
		"ACCEPTED":               "请求已被接受，但尚未处理",
		"NON_AUTHORITATIVE_INFO": "请求已成功处理，结果来自第三方拷贝",
		"NO_CONTENT":             "请求已成功处理，但无返回内容",
		"RESET_CONTENT":          "请求已成功处理，但需重置内容",
		"PARTIAL_CONTENT":        "请求已成功处理，但仅返回了部分内容",
		"MULTI_STATUS":           "请求已成功处理，返回了多个状态的XML消息",
		"ALREADY_REPORTED":       "响应已发送",
		"I_M_USED":               "已完成响应",

		"MULTIPLE_CHOICES":   "返回多条重定向供选择",
		"MOVED_PERMANENTLY":  "永久重定向",
		"FOUND":              "临时重定向",
		"SEE_OTHER":          "当前请求的资源在其它地址",
		"NOT_MODIFIED":       "请求资源与本地缓存相同，未修改",
		"USE_PROXY":          "必须通过代理访问",
		"EMPTY306":           "请切换代理", // Switch Proxy
		"TEMPORARY_REDIRECT": "临时重定向",
		"PERMANENT_REDIRECT": "永久重定向，且禁止改变http方法",

		"BAD_REQUEST":                     "请求错误或参数错误",
		"UNAUTHORIZED":                    "请先进行身份认证",
		"PAYMENT_REQUIRED":                "code=402",
		"FORBIDDEN":                       "无权限访问",
		"NOT_FOUND":                       "请求的内容未找到或已删除",
		"METHOD_NOT_ALLOWED":              "不允许的请求方法",
		"NOT_ACCEPTABLE":                  "无法响应，因资源无法满足客户端条件",
		"PROXY_AUTH_REQUIRED":             "请通过代理的身份认证",
		"REQUEST_TIMEOUT":                 "请求超时",
		"CONFLICT":                        "请求存在冲突",
		"GONE":                            "资源已被删除",
		"LENGTH_REQUIRED":                 "无法处理该请求",
		"PRECONDITION_FAILED":             "请求条件错误",
		"REQUEST_ENTITY_TOO_LARGE":        "请求的实体过大",
		"REQUEST_URI_TOO_LONG":            "请求的URI过长",
		"UNSUPPORTED_MEDIA_TYPE":          "无法处理的媒体格式",
		"REQUESTED_RANGE_NOT_SATISFIABLE": "请求范围不符合要求",
		"EXPECTATION_FAILED":              "请求未满足期望值",
		"TEAPOT":                          "TEAPOT",
		"MISDIRECTED_REQUEST":             "错误的请求",
		"UNPROCESSABLE_ENTITY":            "请求的语义错误",
		"LOCKED":                          "当前资源被锁定",
		"FAILED_DEPENDENCY":               "当前请求失败",
		"TOO_EARLY":                       "未知的请求",
		"UPGRADE_REQUIRED":                "请切换或升级请求协议",
		"PRECONDITION_REQUIRED":           "请求未带条件",
		"TOO_MANY_REQUESTS":               "请求过于频繁，请稍后再试",
		"REQUEST_HEADER_FIELDS_TOO_LARGE": "请求头过大",
		"UNAVAILABLE_FOR_LEGAL_REASONS":   "访问被拒绝(法律的要求)",

		"INTERNAL_SERVER":                 "网络开小差了，请稍后再试",
		"NOT_IMPLEMENTED":                 "服务器不支持的请求方法",
		"BAD_GATEWAY":                     "网关无响应",
		"SERVICE_UNAVAILABLE":             "服务器端临时错误",
		"GATEWAY_TIMEOUT":                 "网关超时",
		"HTTP_VERSION_NOT_SUPPORTED":      "服务器不支持的HTTP版本",
		"VARIANT_ALSO_NEGOTIATES":         "服务器内部配置错误",
		"INSUFFICIENT_STORAGE":            "服务器无法存储请求",
		"LOOP_DETECTED":                   "服务器因死循环而终止操作",
		"NOT_EXTENDED":                    "获取资源策略未被满足",
		"NETWORK_AUTHENTICATION_REQUIRED": "请携带许可凭证进行访问",
	}
)
