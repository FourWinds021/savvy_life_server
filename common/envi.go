package common

var GrpcAddr string
var GatewayAddr string
var SwaggerEnable bool
var SwaggerAddr string
var SwaggerUrl string
var ReadTimeout int64
var CheckAuthEnable bool

func InitEnvi() (err error) {
	GrpcAddr = ConfigFile.Section("service").Key("grpc_addr").
		MustString(ConfigDefaultGrpcAddr)

	GatewayAddr = ConfigFile.Section("service").Key("gateway_addr").
		MustString(ConfigDefaultGatewayAddr)

	SwaggerEnable = ConfigFile.Section("service").Key("swagger_enable").
		MustBool(ConfigDefaultSwaggerEnable)

	SwaggerAddr = ConfigFile.Section("service").Key("swagger_addr").
		MustString(ConfigDefaultSwaggerAddr)

	SwaggerUrl = ConfigFile.Section("service").Key("swagger_url").
		MustString(ConfigDefaultSwaggerUrl)

	ReadTimeout = ConfigFile.Section("service").Key("read_timeout").
		MustInt64(ConfigDefaultReadTimeout)

	CheckAuthEnable = ConfigFile.Section("service").Key("check_auth_enable").
		MustBool(ConfigDefaultCheckAuthEnable)

	return nil
}
