package v2ray

import (
	"stockland/config"
	"stockland/entity"
)

type Link map[string]string

type Presenter interface {
	Present() *Link
}

func NewLink(account *entity.Account, config *config.Config) *Link {
	var presenter Presenter
	// create presenter
	switch account.Type {
	case entity.AccountTypeVless:
		presenter = NewV2ray(account, config)
	case entity.AccountTypeVMess:
		presenter = NewVMess(account, config)

	default:
		presenter = NewV2ray(account, config)
	}

	return presenter.Present()
}
