package utils

// 错误码定义
const (
	// 通用错误码 (0-99)
	SuccessCode      = iota // 0
	UnknownError            // 1
	InvalidParameter        // 2
	Unauthorized            // 3
	Forbidden               // 4
	RecordNotFound          // 5

	// 用户操作的相关错误
	UserAlreadyExists = 100 + iota
	UserNotFound
	InvalidCredentials
	UserNotActive

	// 分类相关错误码 (200-299)
	CategoryAlreadyExists = 200 + iota // 200
	CategoryNotFound                   // 201

	// 标签相关错误码 (300-399)
	TagAlreadyExists = 300 + iota // 300
	TagNotFound                   // 301

	// 文章相关错误码 (400-499)
	PostAlreadyExists = 400 + iota // 400
	PostNotFound                   // 401

	// 数据库相关错误码 (500-599)
	DatabaseError          = 500 + iota // 500
	DatabaseRecordNotFound              // 501

)

// 错误码对应的消息映射
var ErrorCodeMessages = map[int]string{
	SuccessCode:      "success",
	UnknownError:     "未知错误",
	InvalidParameter: "参数错误",
	Unauthorized:     "未授权访问",
	Forbidden:        "权限不足",
	RecordNotFound:   "记录不存在",

	// 用户相关错误
	UserAlreadyExists:  "用户已存在",
	UserNotFound:       "用户不存在",
	InvalidCredentials: "用户名或密码错误",
	UserNotActive:      "用户未激活",

	// 分类相关错误
	CategoryAlreadyExists: "分类已存在",
	CategoryNotFound:      "分类不存在",

	// 标签相关错误
	TagAlreadyExists: "标签已存在",
	TagNotFound:      "标签不存在",

	// 文章相关错误
	PostAlreadyExists: "文章已存在",
	PostNotFound:      "文章不存在",

	// 数据库相关错误
	DatabaseError:          "数据库错误",
	DatabaseRecordNotFound: "数据库记录不存在",
}
