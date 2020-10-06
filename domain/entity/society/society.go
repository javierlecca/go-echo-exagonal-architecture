package society

type Society struct {
	ID         uint    `json:"id" db:"id" gorm:"primaryKey; autoIncrement; column:id"`
	Name       string  `json:"name" db:"name" gorm:"varchar(255); column:name"`
	Logo       string  `json:"logo" db:"logo" gorm:"varchar(255); column:logo"`
	Country    string  `json:"country" db:"country" gorm:"varchar(255); column:country"`
	Department string  `json:"department" db:"department" gorm:"varchar(255); column:department"`
	Longitude  float32 `json:"longitude" db:"longitude" gorm:"decimal(); column:longitude"`
	Latitude   float32 `json:"latitude" db:"latitude" gorm:"decimal(); column:latitude"`
	Web        string  `json:"web" db:"web" gorm:"varchar(255); column:web"`
}

func (Society) TableName() string {
	return "society"
}
