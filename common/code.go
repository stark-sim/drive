package common

type MyCode int64

const (
	CodeSuccess              MyCode = 20000
	CodeUserExist            MyCode = 20001
	CodeUserNotExist         MyCode = 20002
	CodeInvalidPassword      MyCode = 20003
	CodeServerBusy           MyCode = 20004
	CodeNotLogin             MyCode = 20005
	CodeRoleNameExist        MyCode = 20006
	CodeUserAuthNotEnough    MyCode = 20007

	CodeInvalidParams        MyCode = 30000
	CodeInvalidToken         MyCode = 30001
	CodeInvalidAuthFormat    MyCode = 30002

	CodeServerError          MyCode = 50000
	CodeGrpcError            MyCode = 50001
	CodeGrpcDialError        MyCode = 50002
	CodeServerJwtError       MyCode = 50003
	CodeServerDBError        MyCode = 50004
	CodeServerIpfsError      MyCode = 50005
)

var msgFlags = map[MyCode]string{
	CodeSuccess:              "success",
	CodeUserExist:            "用户名重复",
	CodeUserNotExist:         "用户不存在",
	CodeInvalidPassword:      "用户名或密码错误",
	CodeServerBusy:           "服务繁忙",
	CodeNotLogin:             "未登录",
	CodeRoleNameExist:        "角色名重复",
	CodeUserAuthNotEnough:    "用户权限不足",

	CodeInvalidParams:        "非法参数或缺失",
	CodeInvalidToken:         "无效的cookie或token",
	CodeInvalidAuthFormat:    "认证格式有误",

	CodeServerError:          "服务器错误",
	CodeGrpcError:            "grpc调用错误",
	CodeGrpcDialError:        "grpc连接错误",
	CodeServerJwtError:       "服务器JWT错误",
	CodeServerDBError:        "服务器数据库错误",
	CodeServerIpfsError:      "服务器IPFS错误",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
