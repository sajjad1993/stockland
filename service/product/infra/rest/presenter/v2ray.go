package presenter

import (
	"fmt"
	"math/rand"
	"stockland/api/presenter/v2ray"
	"stockland/config"
	"stockland/entity"
	"strings"
	"time"
)

type V2ray struct {
	V2ray1 string `json:"v_2_ray_1"`
	V2ray2 string `json:"v_2_ray_2"`
	V2ray3 string `json:"v_2_ray_3"`
	V2ray4 string `json:"v_2_ray_4"`
	V2ray5 string `json:"v_2_ray_5"`
}

func getV2ray1(account *entity.Account, config *config.Config) string {
	return fmt.Sprintf("vless://%s@%s:%d?encryption=none&security=tls&sni=%s&fp=chrome&type=ws&host=%s&path=%%2Fgraphql%%2F%d#%s",
		account.ClientId, config.MciAddress, config.Port443, getSni(account.Server.Alias), account.Server.Alias, account.Port, account.Remark,
	)
}

func getV2ray2(account *entity.Account, config *config.Config) string {
	return fmt.Sprintf("vless://%s@%s:%d?encryption=none&security=none&sni=%s&fp=chrome&type=ws&host=%s&path=%%2Fgraphql%%2F%d#%s",
		account.ClientId, config.MciXAddress, config.Port80, getSni(account.Server.Alias), account.Server.Alias, account.Port, account.Remark,
	)
}

func getV2ray3(account *entity.Account, config *config.Config) string {
	return fmt.Sprintf("vless://%s@%s:%d?encryption=none&security=tls&sni=%s&fp=chrome&type=ws&host=%s&path=%%2Fgraphql%%2F%d#%s",
		account.ClientId, config.MtnAddress, config.Port443, getSni(account.Server.Alias), account.Server.Alias, account.Port, account.Remark,
	)
}

func getV2ray4(account *entity.Account, config *config.Config) string {
	return fmt.Sprintf("vless://%s@%s:%d?encryption=none&security=none&sni=%s&fp=chrome&type=ws&host=%s&path=%%2Fgraphql%%2F%d#%s",
		account.ClientId, config.MtnXAddress, config.Port80, getSni(account.Server.Alias), account.Server.Alias, account.Port, account.Remark,
	)
}

func getV2ray5(account *entity.Account, config *config.Config) string {
	return fmt.Sprintf("vless://%s@%s:%d?encryption=none&security=tls&sni=%s&fp=chrome&type=ws&host=%s&path=%%2Fgraphql%%2F%d#%s",
		account.ClientId, config.WifiAddress, config.Port2083, getSni(account.Server.Alias), account.Server.Alias, account.Port, account.Remark,
	)
}

func getSni(alias string) string {
	part := strings.Split(alias, ".")
	random := randomString(10)
	return fmt.Sprintf("%s.%s.%s", random, part[1], part[2])
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}

func (v *V2ray) Present() *v2ray.Link {
	return nil
}

func NewV2ray(account *entity.Account, config *config.Config) *V2ray {
	return &V2ray{
		V2ray1: getV2ray1(account, config),
		V2ray2: getV2ray2(account, config),
		V2ray3: getV2ray3(account, config),
		V2ray4: getV2ray4(account, config),
		V2ray5: getV2ray5(account, config),
	}
}
