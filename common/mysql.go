package common

import (
	"github.com/arthurkiller/rollingwriter"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var Engine *xorm.Engine

func InitMySQL() (err error) {
	addr := ConfigFile.Section("mysql").Key("addr").MustString("")
	user := ConfigFile.Section("mysql").Key("user").MustString("")
	passwd := ConfigFile.Section("mysql").Key("passwd").MustString("")
	db := ConfigFile.Section("mysql").Key("db").MustString("")
	dataSourceName := user + ":" + passwd + "@tcp(" + addr + ")/" + db + "?charset=utf8"
	Engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		print("xorm.NewEngine command failed. err:", err, " dataSourceName: ", dataSourceName)
		return err
	}

	Engine.DatabaseTZ = time.Local
	Engine.TZLocation = time.Local
	maxOpenConnections := ConfigFile.Section("mysql").Key("max_open_connections").
		MustInt(ConfigDefaultMysqlMaxOpenConnections)
	maxIdleConnections := ConfigFile.Section("mysql").Key("max_idle_connections").
		MustInt(ConfigDefaultMysqlMaxIdleConnections)
	Engine.SetMaxOpenConns(maxOpenConnections)
	Engine.SetMaxIdleConns(maxIdleConnections)

	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, "t_")
	Engine.SetTableMapper(tbMapper)
	Engine.SetColumnMapper(names.SnakeMapper{})

	config := rollingwriter.Config{
		LogPath:                "../logs",                   //日志路径
		TimeTagFormat:          "200601021504",              //时间格式串
		FileName:               "sql",                       //日志文件名
		MaxRemain:              14,                          //配置日志最大存留数
		RollingPolicy:          rollingwriter.VolumeRolling, //配置滚动策略 norolling timerolling volumerolling
		RollingTimePattern:     "0 0 0 * * *",               //配置时间滚动策略
		RollingVolumeSize:      "512M",                      //配置截断文件下限大小
		WriterMode:             "none",
		BufferWriterThershould: 256,
		Compress:               true, //Compress will compress log file with gzip
	}
	writer, err := rollingwriter.NewWriterFromConfig(&config)
	if err != nil {
		print("rollingWriter.NewWriterFromConfig command failed. err:", err)
		return err
	}

	var sqlLogger *log.SimpleLogger
	sqlLogger = log.NewSimpleLogger(writer)
	Engine.SetLogger(sqlLogger)
	Engine.ShowSQL(true)
	Logger.Debug("Func: InitMySQL end.")

	return nil
}
