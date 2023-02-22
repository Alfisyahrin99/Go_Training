package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	db *sqlx.DB
}

func (c *Config) initDb() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	// tampung nilai ENV dari terminal
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbDriver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")

	// data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	// buka koneksi
	db, err := sqlx.Open(dbDriver, dsn)
	// menutup koneksi
	// defer untuk menjalankan suatu statement diakhir block function / atau dieksekusi sebelum fungsi selesai berjalan
	if err != nil {
		fmt.Println(err.Error())
	}

	// check konfigurasi
	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
	}
	c.db = db
}

func (c *Config) DbConnect() *sqlx.DB { // method untuk mendapatkan Koneksi
	return c.db
}

func NewConfig() Config {
	config := Config{}
	config.initDb() // untuk menjalankan method yang didalamnya membuka koneksi ke DB
	return config   // Mengirimkan Object Config yang didalamnya terdapat attribute bertipe data Koneksi, namun koneksi tidak dapat diakses secara
	// langsung karena attribute bertipe private
}
