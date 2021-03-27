package env

import "os"

func GetEnv() (string, string, string, string, string) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	return host, user, password, port, name
}
