package config

import (
    "database/sql"
    "fmt"
    "log"

	_ "github.com/lib/pq"
)

// AppPort adalah port aplikasi
var AppPort = ":8000"

// DB adalah instance koneksi database global
var DB *sql.DB

// Init menampilkan log saat konfigurasi dimuat
func Init() {
    log.Println("Configuration loaded")
}

// InitDB menginisialisasi koneksi ke database PostgreSQL
func InitDB() {
    // Data Source Name (DSN) untuk koneksi ke database
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        "9qasp5v56q8ckkf5dc.leapcellpool.com", // host
        "6438",                                // port
        "eiuugannqyisrfzeqhgv",                // user
        "zliywvsyktuoetzcjjrxqluftmodmi",      // password
        "iwhygaxdzyrbknmjsrkq",                // dbname
        "require",                             // sslmode
    )

    // Membuka koneksi ke database
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Gagal membuka koneksi ke database:", err)
    }

    // Mengecek koneksi ke database
    if err = DB.Ping(); err != nil {
        log.Fatal("Gagal terhubung ke database:", err)
    }

    log.Println("Database connected")
}