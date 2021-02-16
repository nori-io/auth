package postgres

type Auth struct {
	Email    string `gorm:"column:email; type: VARCHAR(32)"`
	Password string `gorm:"column:email; type: VARCHAR(32)"`
}
