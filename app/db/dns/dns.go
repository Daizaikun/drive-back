package dns

import (
	"fmt"

	"github.com/Daizaikun/wompi-back/app/lib"
)

func New() string {

	host := lib.CheckEnv("DB_HOST", "localhost")

	user := lib.CheckEnv("DB_USER", "root")

	password := lib.CheckEnv("DB_PASSWORD", "12323")

	name := lib.CheckEnv("DB_USER", "root")

	port := lib.CheckEnv("DB_PORT", "5432")


	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s
	sslmode=disable TimeZone=America/Bogota`, host, user, password, name, port)

	return dsn

}