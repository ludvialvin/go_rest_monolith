package config

import (
	"fmt"
	"os"
)

var (
	ENVIRONMENT = getEnv("ENVIRONMENT", "development")
	PORT        = getEnv("PORT", "3000")
	DB          = getEnv("DB", "test_db")
	TOKENKEY    = getEnv("TOKEN_KEY", "gVkYp3s6v9y$B&E)H+MbQeThWmZq4t7w")
	TOKENEXP    = getEnv("TOKEN_EXP", "10h")
	X_API_KEY   = getEnv("X_API_KEY", "C5OW1GYWvztr0nKeZCuPtoF5DEKAYhOK8vrqOiZhb6jYEZBWgFxf8KP6G0pFcoeVacv6fsrOsuQZoy9t8ycOACMusgJP9jp9PAxGdzbKwWz7dUOKEYv6sCAiUe7RXvaq")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
