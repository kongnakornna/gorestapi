package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id                 uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name               string     `gorm:"type:varchar(100);not null"`
	Email              string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password           string     `gorm:"type:varchar(100);not null"`
	CreatedAt          time.Time  `gorm:"not null;default:now()"`
	UpdatedAt          time.Time  `gorm:"not null;default:now()"`
	IsActive           bool       `gorm:"not null;default:true"`
	IsSuperUser        bool       `gorm:"not null;default:false"`
	Verified           bool       `gorm:"not null;default:false"`
	VerificationCode   *string    `gorm:"type:varchar(32);default:null"`
	PasswordResetToken *string    `gorm:"type:varchar(32);default:null"`
	PasswordResetAt    *time.Time `gorm:"default:null"`
	Items              []Item     `gorm:"foreignKey:OwnerId;references:Id"`
	DeleteDate            *time.Time `gorm:"column:deletedate;type:date;default:null"`
	RoleID                int        `gorm:"column:role_id;not null;default:3"`
	PasswordTemp          *string    `gorm:"column:password_temp;default:null"`
	Firstname             *string    `gorm:"column:firstname;default:null"`
	Lastname              *string    `gorm:"column:lastname;default:null"`
	Fullname              *string    `gorm:"column:fullname;default:null"`
	Nickname              *string    `gorm:"column:nickname;default:null"`
	IDCard                *string    `gorm:"column:idcard;default:null"`
	Lastsignindate        time.Time  `gorm:"column:lastsignindate;not null;default:now()"`
	Status                int16      `gorm:"column:status;not null;default:0"`          // 1 = active, 0 = inactive
	ActiveStatus          *int16     `gorm:"column:active_status;default:0"`
	NetworkID             *int       `gorm:"column:network_id;default:1"`
	Remark                *string    `gorm:"column:remark;default:null"`
	InfomationAgreeStatus *int16     `gorm:"column:infomation_agree_status;default:0"`
	Gender                *string    `gorm:"column:gender;default:null"`
	Birthday              *time.Time `gorm:"column:birthday;type:date;default:null"`
	OnlineStatus          *string    `gorm:"column:online_status;default:'0'"`
	Message               *string    `gorm:"column:message;default:null"`
	NetworkTypeID         *int       `gorm:"column:network_type_id;default:0"`
	PublicStatus          *int16     `gorm:"column:public_status;default:0"`
	TypeID                *int       `gorm:"column:type_id;default:0"`
	Avatarpath            *string    `gorm:"column:avatarpath;default:null"`
	Avatar                *string    `gorm:"column:avatar;default:null"`
	RefreshToken          *string    `gorm:"column:refresh_token;type:text;default:null"`
	Loginfailed           *int16     `gorm:"column:loginfailed";default:null"`
	PublicNotification    *int16     `gorm:"column:public_notification;default:0"`
	SmsNotification       *int16     `gorm:"column:sms_notification;default:0"`
	EmailNotification     *int16     `gorm:"column:email_notification;default:0"`
	LineNotification      *int16     `gorm:"column:line_notification;default:0"`
	MobileNumber          *string    `gorm:"column:mobile_number;default:'0'"`
	PhoneNumber           *string    `gorm:"column:phone_number;default:'0'"`
	LineID                *string    `gorm:"column:lineid;default:'0'"`
	SystemID              *string    `gorm:"column:system_id;default:'1'"`
	LocationID            *string    `gorm:"column:location_id;default:'1'"`   
} 

func (User) TableName() string {
	return "user"
}
