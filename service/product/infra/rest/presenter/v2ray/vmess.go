package v2ray

import (
	"fmt"
	"stockland/config"
	"stockland/entity"
)

type Vmess struct {
	V2ray1 string `json:"v_2_ray_1"`
}

func get1(account *entity.Account, config *config.Config) string {
	return fmt.Sprintf("amir test ://%s@%s:%d?encryption=none&security=tls&sni=%s&fp=chrome&type=ws&host=%s&path=%%2Fgraphql%%2F%d#%s",
		account.ClientId, config.MciAddress, config.Port443, getSni(account.Server.Alias), account.Server.Alias, account.Port, account.Remark,
	)
}

func (v *Vmess) Present() *Link {
	var l *Link

	l = (*Link)(&map[string]string{
		"v_2_ray_1": v.V2ray1,
	})
	return l
}
func NewVMess(account *entity.Account, config *config.Config) *Vmess {
	return &Vmess{
		V2ray1: get1(account, config),
	}
}
