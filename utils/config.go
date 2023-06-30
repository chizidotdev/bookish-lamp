package utils

import "os"

// Config stores all the configuration for the application
// using values read by viper from the config file or env variables
type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	PORT       string `mapstructure:"PORT"`
	AuthSecret string `mapstructure:"AUTH_SECRET"`

	Auth0Domain      string `mapstructure:"AUTH0_DOMAIN"`
	Auth0Audience    string `mapstructure:"AUTH0_AUDIENCE"`
	Auth0ClientID    string `mapstructure:"AUTH0_CLIENT_ID"`
	Auth0CallbackURL string `mapstructure:"AUTH0_CALLBACK_URL"`
}

var EnvVars Config

// LoadConfig loads the configuration from the config file or env variable
func LoadConfig() {
	EnvVars.DBDriver = os.Getenv("DB_DRIVER")
	EnvVars.DBSource = os.Getenv("DB_SOURCE")
	EnvVars.PORT = os.Getenv("PORT")
	EnvVars.AuthSecret = os.Getenv("AUTH_SECRET")
}
