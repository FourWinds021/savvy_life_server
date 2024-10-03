package common

import (
	"fmt"
	"os"
)

func InitVersion(version, buildTime string) (err error) {
	Logger.Debug("Func: InitVersion start.")
	versionFile := Path +
		ConfigFile.Section("service").Key("version_file").
			MustString(ConfigDefaultVersionFile)
	// 检查文件是否存在
	if _, err := os.Stat(versionFile); os.IsNotExist(err) {
		// 如果文件不存在，创建一个空的文件
		Logger.Debug(fmt.Sprintf("Version file not found, creating new file: %s", versionFile))
		f, err := os.Create(versionFile)
		if err != nil {
			Logger.Error(fmt.Sprintf("Failed to create version file: %s, error: %v", versionFile, err))
			return err
		}
		defer f.Close() // 确保文件在使用完后关闭
		Logger.Debug("Version file created successfully.")
	}

	var f *os.File
	f, err = os.OpenFile(versionFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		Logger.Error("open version file failed. error:", err)
		return err
	}
	_, err = f.WriteString(fmt.Sprintf("%s\n%s", version, buildTime))
	if err != nil {
		Logger.Error("write version file failed. error:", err)
		return err
	}
	Logger.Info("version file init success")
	Logger.Debug("Func: InitVersion end.")
	return nil
}
