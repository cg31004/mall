package po

import "time"

type Member struct {
	Id       string `gorm:"column:id;primary_key"`
	Account  string `gorm:"column:account"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	Salt     string `gorm:"column:salt"`

	CreatedAt *time.Time `gorm:"<-:create;column:created_at"`
	UpdatedAt time.Time  `gorm:"<-:create;column:updated_at"`
}

func (Member) TableName() string {
	return "member"
}

//
type MemberChart struct {
	Id        string `gorm:"column:id;primary_key"`
	MemberId  string `gorm:"column:member_id"`
	ProductId string `gorm:"column:product_id"`
	Quantity  int    `gorm:"column:quantity"`

	CreatedAt *time.Time `gorm:"<-:create;column:created_at"`
	UpdatedAt time.Time  `gorm:"<-:create;column:updated_at"`
}

func (MemberChart) TableName() string {
	return "member_chart"
}

type MemberChartFirst struct {
	MemberId  string
	ProductId string
}

type MemberChartSearch struct {
	MemberId string
}

type MemberChartDel struct {
	Id       string
	MemberId string
}

type MemberChartUpdate struct {
	Id       string
	MemberId string
	Quantity int
}

func (c *MemberChartUpdate) GetUpdateChart() map[string]interface{} {
	return map[string]interface{}{
		"quantity": c.Quantity,
	}
}
