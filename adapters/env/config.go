package env

import "os"

var (
	AdminPhoneNumber string
)

func init() {
	AdminPhoneNumber = os.Getenv("ADMIN_PHONE_NUMBER")
}
