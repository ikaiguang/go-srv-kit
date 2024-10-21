package configutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
)

func AppConfig(bootstrap *configpb.Bootstrap) *configpb.App {
	return bootstrap.GetApp()
}

func SettingConfig(bootstrap *configpb.Bootstrap) *configpb.Setting {
	return bootstrap.GetSetting()
}
func SettingCaptchaConfig(bootstrap *configpb.Bootstrap) *configpb.Setting_Captcha {
	return bootstrap.GetSetting().GetCaptcha()
}
func SettingLoginConfig(bootstrap *configpb.Bootstrap) *configpb.Setting_Login {
	return bootstrap.GetSetting().GetLogin()
}

func HTTPConfig(bootstrap *configpb.Bootstrap) *configpb.Server_HTTP {
	return bootstrap.GetServer().GetHttp()
}
func GRPCConfig(bootstrap *configpb.Bootstrap) *configpb.Server_GRPC {
	return bootstrap.GetServer().GetGrpc()
}

func LogConfig(bootstrap *configpb.Bootstrap) *configpb.Log {
	return bootstrap.GetLog()
}
func LogConsoleConfig(bootstrap *configpb.Bootstrap) *configpb.Log_Console {
	return bootstrap.GetLog().GetConsole()
}
func LogFileConfig(bootstrap *configpb.Bootstrap) *configpb.Log_File {
	return bootstrap.GetLog().GetFile()
}

func MysqlConfig(bootstrap *configpb.Bootstrap) *configpb.MySQL {
	return bootstrap.GetMysql()
}
func PostgresConfig(bootstrap *configpb.Bootstrap) *configpb.PSQL {
	return bootstrap.GetPsql()
}
func MongoConfig(bootstrap *configpb.Bootstrap) *configpb.Mongo {
	return bootstrap.GetMongo()
}
func RedisConfig(bootstrap *configpb.Bootstrap) *configpb.Redis {
	return bootstrap.GetRedis()
}
func RabbitmqConfig(bootstrap *configpb.Bootstrap) *configpb.Rabbitmq {
	return bootstrap.GetRabbitmq()
}
func ConsulConfig(bootstrap *configpb.Bootstrap) *configpb.Consul {
	return bootstrap.GetConsul()
}
func EtcdConfig(bootstrap *configpb.Bootstrap) *configpb.Etcd {
	return bootstrap.GetEtcd()
}
func JaegerConfig(bootstrap *configpb.Bootstrap) *configpb.Jaeger {
	return bootstrap.GetJaeger()
}

func TransferEncryptConfig(bootstrap *configpb.Bootstrap) *configpb.Encrypt_TransferEncrypt {
	return bootstrap.GetEncrypt().GetTransferEncrypt()
}
func ServiceEncryptConfig(bootstrap *configpb.Bootstrap) *configpb.Encrypt_ServiceEncrypt {
	return bootstrap.GetEncrypt().GetServiceEncrypt()
}
func TokenEncryptConfig(bootstrap *configpb.Bootstrap) *configpb.Encrypt_TokenEncrypt {
	return bootstrap.GetEncrypt().GetTokenEncrypt()
}

func ClusterServiceApis(bootstrap *configpb.Bootstrap) []*configpb.ClusterServiceApi {
	return bootstrap.GetClusterServiceApi()
}
func ThirdPartyApis(bootstrap *configpb.Bootstrap) []*configpb.ThirdPartyApi {
	return bootstrap.GetThirdPartyApi()
}

func SnowflakeConfig(bootstrap *configpb.Bootstrap) *configpb.Snowflake {
	return bootstrap.GetSnowflake()
}
