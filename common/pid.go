package common

import (
	"fmt"
	"os"
)

func InitPid() (err error) {
	Logger.Debug("Func: InitPid start.")
	pidFile := Path +
		ConfigFile.Section("service").Key("pid_file").
			MustString(ConfigDefaultPidFile)
	var f *os.File

	f, err = os.OpenFile(pidFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		Logger.Error("open pid file failed. error:", err)
		return err
	}

	_, err = f.WriteString(fmt.Sprintf("%d", os.Getpid()))
	if err != nil {
		Logger.Error("write pid file failed. error:", err)
		return err
	}
	Logger.Info("pid file init success")
	Logger.Debug("Func: InitPid end.")
	return nil
}
