package models

// User 定义请求参数结构体
type User struct {
	UserID       uint64 `json:"user_id,string" db:"user_id"` // 指定json序列化/反序列化时使用小写user_id
	UserName     string `json:"username" db:"username"`
	Password     string `json:"password" db:"password"`
	Email        string `json:"email" db:"gender"`  // 邮箱
	Gender       int    `json:"gender" db:"gender"` // 性别
	AccessToken  string
	RefreshToken string
}
