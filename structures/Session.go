package structures

type Session struct {
	Id uint32
	PackageId uint32 `gorm:"column:package_id"`
	SessionKey string `gorm:"column:session_key"`
}

func (Session) TableName() string {
	return "download_queue"
}
