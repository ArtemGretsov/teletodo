package entities

import "gorm.io/gorm"

type AuthType string

const TelegramAuthType AuthType = "telegram"

type UserAuthentication struct {
	gorm.Model
	ExternalID string   `gorm:"type:varchar;index:idx_external_id_user_id_auth_type"`
	UserID     uint     `gorm:"index:idx_external_id_user_id_auth_type"`
	AuthType   AuthType `gorm:"type:user_auth_type;index:idx_external_id_user_id_auth_type"`
}
