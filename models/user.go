package models

type Users struct {
	UserId    int    `json:"user_id" db:"user_id"` //用户ID
	Name      string `json:"name" db:"name"`
	Password  string `json:"password" db:"password"`
	Avatar    string `json:"avatar" db:"avatar"`
	Token     string `json:"token" db:"token"`
	CreatedAt string `json:"created_at" db:"created_at" time_format:"2006-01-02 15:04:05"` //创建时间
	UpdatedAt string `json:"updated_at" db:"updated_at" time_format:"2006-01-02 15:04:05"` //修改时间
}

func (this *Users) DbName() string {
	return "default"
}

func (this *Users) TableName() string {
	return "users"
}

func (this *Users) PK() string {
	return "user_id"
}
