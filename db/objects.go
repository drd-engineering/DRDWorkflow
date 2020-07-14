package db

import (
	"time"
)

type User struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Gender       string
	Email        string
	KtpNumber    int64 `gorm:"unique;not null"`
	Address      string
	PhoneNumber  string
	DateOfBirth  time.Time
	Cityzenship  string
	PlaceOfBirth string
	IsActive     bool
	Signature    string
	Signature2   string
	Signature3   string
}
type Company struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	IsActive  bool
	Name      string
	Email     string
}
type Division struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Level     int
	ParentID  int
	CompanyID string
	Parent    *Division `gorm:"foreignkey:ParentID"`
	Company   Company   `gorm:"foreignkey:CompanyID"`
}
type AppToken struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	IsActive  bool
	Token     string
	CompanyID string
	Company   Company `gorm:"foreignkey:CompanyID"`
}
type ApiLog struct {
	ID                int `gorm:"primary_key"`
	Timestamp         time.Time
	Ttl               int
	ApiResponseStatus int
	AppTokenID        int
	ApiTypeID         int
	AppToken          AppToken `gorm:"foreignkey:AppTokenID"`
	ApiType           ApiType  `gorm:"foreignkey:ApiTypeID"`
}
type ApiType struct {
	ID          int `gorm:"primary_key"`
	Path        string
	Method      string
	Url         string
	Description string
}
type Activity struct {
	ID                    int
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time
	Name                  string
	Permission            bool
	IsDone                bool
	CssProperty           string
	ExpiredDate           time.Time
	UserID                string
	DivisionID            int
	AlternativeUserID     string
	AlternativeDivisionID int
	User                  User      `gorm:"foreignkey:UserID"`
	Division              *Division `gorm:"foreignkey:DivisionID"`
	AlternativeUser       *User     `gorm:"foreignkey:AlternativeUserID"`
	AlternativeDivision   *Division `gorm:"foreignkey:AlternativeDivisionID"`
}
type Workflow struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	IsActive    bool
	IsTemplate  bool
	Name        string
	IsStatic    bool
	RevisedType int
	TotalUsed   int
	FirstID     int
	LastID      int
	CreatorID   string
	AppTokenID  int
	First       Activity `gorm:"foreignkey:FirstID"`
	Last        Activity `gorm:"foreignkey:LastID"`
	Creator     User     `gorm:"foreignkey:CreatorID"`
	AppToken    AppToken `gorm:"foreignkey:AppTokenID"`
}
type ActivityLink struct {
	ID         int `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Type       string
	PreviousID int
	NextID     int
	Previous   Activity `gorm:"foreignkey:PreviousID"`
	Next       Activity `gorm:"foreignkey:NextID"`
}
type ActivityLog struct {
	ID         int `gorm:"primary_key"`
	Timestamp  time.Time
	Action     string
	ActivityID int
	UserID     string
	Activity   Activity `gorm:"foreignkey:ActivityID"`
	User       User     `gorm:"foreignkey:UserID"`
}
type AuditTrails struct {
	ID        int `gorm:"primary_key"`
	Timestamp time.Time
	Activity  string
	UserID    string
	CompanyID string
	User      User    `gorm:"foreignkey:UserID"`
	Company   Company `gorm:"foreignkey:CompanyID"`
}
type UserActivity struct {
	ID         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	IsActive   bool
	Message    string
	IsRead     bool
	Status     int
	UserID     string
	ActivityID int
	WorkflowID string
	User       User     `gorm:"foreignkey:UserID"`
	Activity   Activity `gorm:"foreignkey:ActivityID"`
	Workflow   Workflow `gorm:"foreignkey:WorkflowID"`
}
