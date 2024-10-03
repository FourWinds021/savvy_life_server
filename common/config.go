package common

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"path/filepath"
)

var ConfigFile *ini.File
var Path string

func InitConfig() (err error) {
	fmt.Printf("Init Config Load start.\n")
	Path, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	//ConfigFile, err = ini.Load(Path + "/../config/config.ini")
	ConfigFile, err = ini.Load(Path + "/config/config.ini")
	if err != nil {
		fmt.Printf("Ini Config Load failed. err:%s\n", err)
		return err
	}
	fmt.Printf("Init Config Load end.\n")
	return nil
}
