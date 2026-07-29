package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Transaction struct {
	ID uint `gorm:"primaryKey" queryParam:"id"`

	UserID uint `gorm:"not null" queryParam:"user_id"`
	User   User

	Name              string            `gorm:"size:100;not null" queryParam:"name"`
	Amount            float64           `gorm:"type:numeric(12,2);not null" queryParam:"amount"`
	PaymentMethodType PaymentMethodType `gorm:"type:payment_method;not null" queryParam:"payment_method_type"`
	TransactionType   TransactionType   `gorm:"type:transaction_type;not null" queryParam:"transaction_type"`
	StatusType        StatusType        `gorm:"type:status;not null" queryParam:"status_type"`
	CreatedAt         time.Time `queryParam:"created_at"`
	CompletedAt         time.Time
	Metadata          TransactionMetaData `gorm:"type:jsonb"`
}

type PaymentMethodType string

const (
	TypeCard         PaymentMethodType = "CARD"
	TypeBankTransfer PaymentMethodType = "BANK_TRANSFER"
	TypePayPal       PaymentMethodType = "PAYPAL"
	TypeCrypto       PaymentMethodType = "CRYPTO"
	TypeWallet       PaymentMethodType = "WALLET"
)

type TransactionType string

const (
	TypeDeposit        TransactionType = "DEPOSIT"
	TypeWithdrawal     TransactionType = "WITHDRAWAL"
	TypeProgramPayment TransactionType = "PROGRAM_PAYMENT"
	TypeRefund         TransactionType = "REFUND"
	TypePayout         TransactionType = "PAYOUT"
)

type StatusType string

const (
	StatusTypePending   StatusType = "PENDING"
	StatusTypeCompleted StatusType = "COMPLETED"
	StatusTypeFailed    StatusType = "FAILED"
	StatusTypeCanceled  StatusType = "CANCELLED"
)

type TransactionMetaData struct {
}

func (m *TransactionMetaData) Scan(value any) error {
	if value == nil {
		*m = TransactionMetaData{}
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(data, m)
}

func (m *TransactionMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}
