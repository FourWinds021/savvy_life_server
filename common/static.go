package common

// server
const (
	ConfigDefaultGrpcAddr        = ":50053"
	ConfigDefaultGatewayAddr     = ":39254"
	ConfigDefaultSwaggerEnable   = false
	ConfigDefaultSwaggerAddr     = ":9092"
	ConfigDefaultSwaggerUrl      = "http://127.0.0.1:9092/swagger/savvy.swagger.json"
	ConfigDefaultReadTimeout     = 60
	ConfigDefaultCheckAuthEnable = true
	ConfigDefaultPidFile         = "/../proc/savvy_life.pid"
	ConfigDefaultVersionFile     = "/../proc/savvy_life.version"
)

// mysql
const (
	ConfigDefaultMysqlMaxOpenConnections = 2000
	ConfigDefaultMysqlMaxIdleConnections = 100
)

// business
const (
	FullMethodUserGet = "/savvy_life.proto.api.v1.User"
)
