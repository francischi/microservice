package main

import (
	"server/pkg/helpers"
	"server/pkg/migrate/migrations"
)

func main(){
	helpers.InitMySql()
	migrations.Member(helpers.SqlSession)
}