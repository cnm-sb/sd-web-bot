package plugininterface

import "github.com/gin-gonic/gin"

// PluginInterface 定义插件接口
type PluginInterface interface {
	RegisterRoutes() []gin.RouteInfo
	GetName() string
}
