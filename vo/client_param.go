package vo

import "github.com/bidaya0/nacos-sdk-go/common/constant"

type NacosClientParam struct {
	ClientConfig  *constant.ClientConfig  // optional
	ServerConfigs []constant.ServerConfig // optional
}
