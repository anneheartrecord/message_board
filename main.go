package main

import (
	"newProject/messageboard/cmd"
	"newProject/messageboard/dao"
)

func main()  {
	dao.MysqlInit()
	cmd.Entrance()
}
