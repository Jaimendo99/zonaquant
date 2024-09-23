package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountType struct {
	gorm.Model
	AccountType string    `gorm:"unique;not null"`
	Accounts    []Account `gorm:"foreignKey:AccountTypeID"`
}

type Account struct {
	gorm.Model
	AccountTypeID    uint          `gorm:"not null;index"`
	AccountType      AccountType   `gorm:"foreignKey:AccountTypeID"`
	Balance          float64       `gorm:"not null"`
	Users            []User        `gorm:"many2many:users_accounts;constraint:OnDelete:CASCADE;"`
	TransactionsFrom []Transaction `gorm:"foreignKey:FromAccountID"`
	TransactionsTo   []Transaction `gorm:"foreignKey:ToAccountID"`
}

type Role struct {
	gorm.Model
	Role  string `gorm:"unique;not null"`
	Users []User `gorm:"foreignKey:RoleID"`
}

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"` // Ensure passwords are hashed
	RoleID   uint      `gorm:"not null;index"`
	Role     Role      `gorm:"foreignKey:RoleID"`
	Accounts []Account `gorm:"many2many:users_accounts;constraint:OnDelete:CASCADE;"`
}

type UsersAccounts struct {
	UserID    uint `gorm:"primaryKey"`
	AccountID uint `gorm:"primaryKey"`
}

type Producer struct {
	gorm.Model
	FullName  string `gorm:"not null"`
	RUC       string `gorm:"unique"`
	Nickname  string
	AccountID uint      `gorm:"not null;unique;index"`
	Account   Account   `gorm:"foreignKey:AccountID"`
	Batches   []Batch   `gorm:"foreignKey:ProducerID"`
	Advances  []Advance `gorm:"foreignKey:ProducerID"`
}
type Empacadora struct {
	gorm.Model
	BusinessName string  `gorm:"not null"`
	RUC          string  `gorm:"unique;not null"`
	AccountID    uint    `gorm:"not null;unique;index"`
	Account      Account `gorm:"foreignKey:AccountID"`
	Batches      []Batch `gorm:"foreignKey:EmpacadoraID"`
}
type PurchaseReport struct {
	gorm.Model
	Date           time.Time              `gorm:"not null"`
	WeightDiscount float64                `gorm:"not null"`
	MoneyDiscount  float64                `gorm:"not null"`
	Total          float64                `gorm:"not null"`
	Commission     float64                `gorm:"not null"`
	Checked        bool                   `gorm:"not null;default:false"`
	Details        []PurchaseReportDetail `gorm:"foreignKey:PurchaseReportID"`
	Batches        []Batch                `gorm:"foreignKey:PurchaseReportID"`
}
type PurchaseReportDetail struct {
	gorm.Model
	SizeGrams        float64        `gorm:"not null"`
	Pounds           float64        `gorm:"not null"`
	PricePerPound    float64        `gorm:"not null"`
	PurchaseReportID uint           `gorm:"not null;index"`
	PurchaseReport   PurchaseReport `gorm:"foreignKey:PurchaseReportID"`
}

type SellingReport struct {
	gorm.Model
	Date               time.Time             `gorm:"not null"`
	FactoryBatch       string                `gorm:"not null"`
	TailWeight         float64               `gorm:"not null"`
	HeadWeight         float64               `gorm:"not null"`
	TrashWeight        float64               `gorm:"not null"`
	LogisticsFee       float64               `gorm:"not null"`
	LogisticsRetention float64               `gorm:"not null"`
	ProductRetention   float64               `gorm:"not null"`
	Details            []SellingReportDetail `gorm:"foreignKey:SellingReportID"`
	Batches            []Batch               `gorm:"foreignKey:SellingReportID"`
}
type Size struct {
	gorm.Model
	SizeValue            string                `gorm:"unique;not null"`
	SellingReportDetails []SellingReportDetail `gorm:"foreignKey:SizeID"`
}
type Class struct {
	gorm.Model
	ClassName            string                `gorm:"unique;not null"`
	SellingReportDetails []SellingReportDetail `gorm:"foreignKey:ClassID"`
}
type SellingReportDetail struct {
	gorm.Model
	SizeID          uint          `gorm:"not null;index"`
	Size            Size          `gorm:"foreignKey:SizeID"`
	ClassID         uint          `gorm:"not null;index"`
	Class           Class         `gorm:"foreignKey:ClassID"`
	BoxQuantity     int           `gorm:"not null"`
	Pounds          float64       `gorm:"not null"`
	PricePerPound   float64       `gorm:"not null"`
	SellingReportID uint          `gorm:"not null;index"`
	SellingReport   SellingReport `gorm:"foreignKey:SellingReportID"`
}
type DeliveryReport struct {
	gorm.Model
	Date               time.Time `gorm:"not null"`
	DeliveryTracking   string    `gorm:"not null"`
	DeliveryWeight     float64   `gorm:"not null"`
	ProductUnit        string    `gorm:"not null"`
	UnitAmount         float64   `gorm:"not null"`
	PricePerUnit       float64   `gorm:"not null"`
	FactoryTotalWeight float64   `gorm:"not null"`
	Batches            []Batch   `gorm:"foreignKey:DeliveryReportID"`
}
type Batch struct {
	gorm.Model
	BatchName        string          `gorm:"unique;not null"`
	ProducerID       uint            `gorm:"not null;index"`
	Producer         Producer        `gorm:"foreignKey:ProducerID"`
	EmpacadoraID     uint            `gorm:"not null;index"`
	Empacadora       Empacadora      `gorm:"foreignKey:EmpacadoraID"`
	PurchaseReportID *uint           `gorm:"index"`
	PurchaseReport   *PurchaseReport `gorm:"foreignKey:PurchaseReportID"`
	SellingReportID  *uint           `gorm:"index"`
	SellingReport    *SellingReport  `gorm:"foreignKey:SellingReportID"`
	DeliveryReportID *uint           `gorm:"index"`
	DeliveryReport   *DeliveryReport `gorm:"foreignKey:DeliveryReportID"`
	Logistics        []Logistics     `gorm:"foreignKey:BatchID"`
	Transactions     []Transaction   `gorm:"foreignKey:BatchID"`
}
type Logistics struct {
	gorm.Model
	Amount      float64   `gorm:"not null"`
	Description string    `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	BatchID     uint      `gorm:"not null;index"`
	Batch       Batch     `gorm:"foreignKey:BatchID"`
}
type TransactionState struct {
	gorm.Model
	StateName    string        `gorm:"unique;not null"`
	Transactions []Transaction `gorm:"foreignKey:StateID"`
}
type Transaction struct {
	gorm.Model
	FromAccountID uint             `gorm:"not null;index"`
	FromAccount   Account          `gorm:"foreignKey:FromAccountID"`
	ToAccountID   uint             `gorm:"not null;index"`
	ToAccount     Account          `gorm:"foreignKey:ToAccountID"`
	Amount        float64          `gorm:"not null"`
	Date          time.Time        `gorm:"not null"`
	StateID       uint             `gorm:"not null;index"`
	State         TransactionState `gorm:"foreignKey:StateID"`
	Description   string
	BatchID       *uint               `gorm:"index"`
	Batch         *Batch              `gorm:"foreignKey:BatchID"`
	Details       []TransactionDetail `gorm:"foreignKey:TransactionID"`
}
type TransactionDetail struct {
	gorm.Model
	TransactionID uint        `gorm:"not null;index"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID"`
	Amount        float64     `gorm:"not null"`
	Description   string
}
type Advance struct {
	gorm.Model
	Amount     float64   `gorm:"not null"`
	ProducerID uint      `gorm:"not null;index"`
	Producer   Producer  `gorm:"foreignKey:ProducerID"`
	Date       time.Time `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

func (Account) TableName() string {
	return "accounts"
}

func (UsersAccounts) TableName() string {
	return "users_accounts"
}
