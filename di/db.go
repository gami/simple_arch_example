package di

import "github.com/gami/simple_arch_example/mysql"

func InjectUserDB() *mysql.DB {
	return mysql.UserDB()
}
