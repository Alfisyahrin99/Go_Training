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
	// Tampung nilai ENV dari
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

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
	// defer untuk menjalankan suatu statement diakhir blok function / atau dieksekusi sebelum fungsi selesai berjalan
	// defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	// check Konfigurasi
	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	c.db = db
}

func (c *Config) DbConnect() *sqlx.DB { // Method untuk mendapatkan koneksi
	return c.db
}

func NewConfig() Config {
	config := Config{}
	config.initDb() // untuk menjalankan method yang didalmnya membuka koneksi ke DB
	return config   // Mengirimkan Object Config yang dalamnya terdapat attribute bertipe data koneksi, namun koneksi tidak dapat diakses secara
	// langsung karena attribute bertipe private
}
