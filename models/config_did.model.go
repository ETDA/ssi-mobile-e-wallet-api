package models

type ConfigDID struct {
	DIDAddress    string `json:"did_address" gorm:"column:did_address"`
	PublicKeyPEM  string `json:"public_key_pem" gorm:"column:public_key_pem"`
	PrivateKeyPEM string `json:"private_key_pem" gorm:"column:private_key_pem"`
}

func (ConfigDID) TableName() string {
	return "configs_did"
}
