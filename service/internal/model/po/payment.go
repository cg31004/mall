package po

import "time"

type Payment struct {
	Id       string `gorm:"column:id;primary_key"`
	MemberId string  `gorm:"column:member_id"`
	Number   string `gorm:"column:number"`

	CreatedAt time.Time `gorm:"<-:create;column:created_at"`
	UpdatedAt time.Time `gorm:"<-:create;column:updated_at"`
}

func (Payment) TableName() string {
	return "payment"
}
