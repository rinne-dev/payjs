package payjs

import (
	"rinne.dev/payjs/config"
)

// PayJS 实例
type PayJS struct {
	Config *config.Config
}

// New 新建 PayJS 实例
func New(merchID, key string) *PayJS {
	return &PayJS{
		Config: &config.Config{
			MerchID: merchID,
			Key:     key,
		},
	}
}
