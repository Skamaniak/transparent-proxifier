package conf

import (
	"github.com/spf13/viper"
)

const TcpTransparentProxyPort = "TCP_TRANSPARENT_PROXY_PORT"
const TlsTransparentProxyPort = "TLS_TRANSPARENT_PROXY_PORT"
const ProxyLocation = "PROXY_LOCATION"

func InitConfig() {
	viper.AutomaticEnv()

	viper.SetDefault(TcpTransparentProxyPort, 1080)
	viper.SetDefault(TlsTransparentProxyPort, 1081)
	viper.SetDefault(ProxyLocation, "0.0.0.0:8080")

}
