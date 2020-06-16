package src
 
import ("time"
		"github.com/jinzhu/gorm"
		uuid "github.com/satori/go.uuid")

type user struct {
	ID int `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"not null" json:"name"`
	BirthDate string `gorm:"type:varchar(10)" json:"birth_date"`
	Email string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
	Debts []debt `gorm:"foreignkey:user_id;"`
	CreatedAt time.Time `gorm:"default:NOW()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:NOW()" json:"updated_at"`
}

func (user) TableName() string {
	return "users"
}

type debt struct {
	ID string `gorm:"primary_key" json:"id"`
	CompanyName string `gorm:"not null" json:"company_name"`
	Value float32  `gorm:"not null;" sql:"type:decimal(10,2)" json:"value"`
	Date string `gorm:"type:varchar(10)" json:"date"`
	Status int `gorm:"not null" json:"status"`
	UserId int `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"default:NOW()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:NOW()" json:"updated_at"`
}

func (debt) TableName() string {
	return "debts"
}

//BeforeCreate execute before insert data on DB
func (o *debt) BeforeCreate(scope *gorm.Scope) error {

	uuid := uuid.NewV4().String()
	return scope.SetColumn("ID", uuid)
}
