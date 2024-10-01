package model

type Blog struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title" gorm:"not null comment 'title'"`
	Desc   string `json:"desc" gorm:"not null comment 'desc'"`
	Image  string `json:"image" gorm:"null comment 'image'"`
	UserId string `json:"userid"`
	User   User   `json:"user";gorm:"foreignkey:UserId"`
}
