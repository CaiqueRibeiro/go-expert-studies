package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// O viper trabalha com annotations, então é possível mapeá-los para os campos do .env
type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config") // nome do arquivo de configuração
	viper.SetConfigType("env")        // tipo do arquivo que será lido
	viper.AddConfigPath(path)         // caminho do arquivo que será lido
	viper.SetConfigFile(".env")       // nome do arquivo que será lido
	viper.AutomaticEnv()              // Isso permite que, caso exista uma variável de ambiente definida no sistema, sobrescreva a definida no .env
	err := viper.ReadInConfig()       // lê o arquivo .env
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg) // transforma o arquivo .env em um struct
	if err != nil {
		panic(err)
	}
	// Cria uma instância do TokenAuth para gerar JWT
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil) // Adiciona o token JWT nas configs
	return cfg, nil
}
