package main

import (
	"fmt"
	. "savvy_life/common"
	. "savvy_life/server"
	. "savvy_life/timer"
)

var (
	Version   = "1.0.0"
	BuildTime = "20230310202956"
)

func main() {
	fmt.Println("Start InitConfig")
	err := InitConfig()
	if err != nil {
		fmt.Println("init config failed. error: ", err)
		return
	}
	fmt.Println("Config loaded successfully")

	fmt.Println("Start InitMySQL")
	err = InitMySQL()
	if err != nil {
		fmt.Println("init mysql failed. error: ", err)
		return
	}
	fmt.Println("MySQL initialized successfully")

	fmt.Println("Start InitEnvi")
	err = InitEnvi()
	if err != nil {
		fmt.Println("init envi failed. error: ", err)
		return
	}
	fmt.Println("Environment initialized successfully")

	fmt.Println("Start InitVersion")
	err = InitVersion(Version, BuildTime)
	if err != nil {
		fmt.Println("init version failed. error: ", err)
		return
	}
	fmt.Println("Version initialized successfully")

	fmt.Println("Start InitPid")
	err = InitPid()
	if err != nil {
		fmt.Println("init pid failed. error: ", err)
		return
	}
	fmt.Println("PID initialized successfully")

	fmt.Println("Start InitBusiness")
	err = InitBusiness()
	if err != nil {
		fmt.Println("init business failed. error: ", err)
		return
	}
	fmt.Println("Business initialized successfully")

	fmt.Println("Start Timer")
	Timer()

	fmt.Println("Service is starting")
	Logger.Info("Service start. good luck!")
	Run()
}
