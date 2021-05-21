package structures

type Package struct {
	Id uint32
	Name string
	Version string
	Ord uint8
	PathToFile string `gorm:"column:path"`
	OnServer string `gorm:"column:on_server"`
}

func (Package) TableName() string {
	return "packages"
}