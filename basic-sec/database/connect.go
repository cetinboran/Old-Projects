package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Config struct {
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBName string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	// Üstteki gibi return yazarsak o değerleri := gibi oluşturmuş oluyor
	// aynı zamanda düz return dersek onları döndürüyor.

	viper.AddConfigPath(path)  // config dosyasının olduğu yer
	viper.SetConfigName("db")  // config dosyasının adı.
	viper.SetConfigType("env") // config dosyasının extentionu.

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

var DBConn *sql.DB

func InitDB() error {
	// viper ile env dosyasıyla önemli bilgileri .env'e sakladım onu git'e atmıcam mesela.
	config, configErr := LoadConfig(".")
	if configErr != nil {
		log.Fatal("Cannot load config:", configErr)
	}

	// Connecting the mysql
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	DBConn, err = sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	// Bağlantı sağlandı mı kontrol ediyoruz.
	err = DBConn.Ping()
	if err != nil {
		return fmt.Errorf("Database connection error %v", err)
	}
	return nil
}
