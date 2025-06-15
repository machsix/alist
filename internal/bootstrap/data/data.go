package data

import "github.com/machsix/alist/v3/cmd/flags"

func InitData() {
	initUser()
	initSettings()
	initTasks()
	if flags.Dev {
		initDevData()
		initDevDo()
	}
}
