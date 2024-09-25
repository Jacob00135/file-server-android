package database

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/Jacob00135/file-server-android/models"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// var DB *sql.DB

type Database struct {
	Conn *sql.DB
}

var DB *Database
var Storage *session.Store

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")
	HOME := os.Getenv("TEST_ROOT")

	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	Storage = session.New()
	DB = &Database{Conn: conn}

	// 创建用户表
	_, err = DB.Conn.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			userid INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			permission INTEGER NOT NULL DEFAULT 1
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// 创建文件表
	_, err = DB.Conn.Exec(`
		CREATE TABLE IF NOT EXISTS directory (
			directoryid INTEGER PRIMARY KEY AUTOINCREMENT,
			directorypath TEXT UNIQUE NOT NULL,
			permission INTEGER NOT NULL DEFAULT 4
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// create test user
	testUser := []struct {
		username   string
		password   string
		permission int
	}{
		{"admin", "admin", 4},
		{"user", "user", 2},
		{"guest", "guest", 1},
	}
	for _, user := range testUser {
		_, err = DB.Conn.Exec("INSERT OR IGNORE INTO users (username, password, permission) VALUES (?, ?, ?)", user.username, user.password, user.permission)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = DB.InsertDir(HOME, 1)
	if err != nil {
		log.Fatal(err)
	}

}

func (db *Database) Connect() {
	if db.Conn != nil {
		return
	}
	conn, err := sql.Open("sqlite3", "./filesystem.db")
	if err != nil {
		log.Fatal(err)
	}
	DB = &Database{Conn: conn}
}

func (db *Database) Close() {
	if db.Conn != nil {
		db.Conn.Close()
		db.Conn = nil
	}
}

func (db *Database) InsertUser(username, password string) error {
	if db.Conn == nil {
		return errors.New("database connection is not open")
	}
	_, err := db.Conn.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}

func (db *Database) InsertDir(path string, p int) error {
	if db.Conn == nil {
		return errors.New("database connection is not open")
	}
	path = filepath.Clean(path)
	_, err := db.Conn.Exec("INSERT OR IGNORE INTO directory (directorypath, permission) VALUES (?, ?)", path, p)
	return err
}

func (db *Database) GetDirsByPermission(p int) ([]string, error) {
	if db.Conn == nil {
		return nil, errors.New("database connection is not open")
	}
	rows, err := db.Conn.Query("SELECT directorypath FROM directory WHERE permission = ?", p)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paths []string
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return paths, nil
}

func (db *Database) CheckUserExists(username string) (bool, error) {
	if db.Conn == nil {
		return false, errors.New("database connection is not open")
	}

	err := db.Conn.QueryRow("SELECT 1 FROM users WHERE username = ?", username).Scan(new(int))
	if err == nil {
		return true, nil // 用户存在
	} else if errors.Is(err, sql.ErrNoRows) {
		return false, nil // 用户不存在
	}
	return false, err // 其他错误
}

func (db *Database) CheckFileExists(path string) (bool, error) {
	if db.Conn == nil {
		return false, errors.New("database connection is not open")
	}

	err := db.Conn.QueryRow("SELECT 1 FROM directory WHERE directorypath = ?", path).Scan(new(int))
	if err == nil {
		return true, nil // 存在
	} else if errors.Is(err, sql.ErrNoRows) {
		return false, nil // 不存在
	}
	return false, err // 其他错误
}

func (db *Database) getUserInfo(username string) (*models.User, error) {
	if db.Conn == nil {
		return nil, errors.New("database connection is not open")
	}
	user := new(models.User)
	err := db.Conn.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Permission)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *Database) getFileInfo(path string) (*models.DbFile, error) {
	if db.Conn == nil {
		return nil, errors.New("database connection is not open")
	}
	file := new(models.DbFile)
	err := db.Conn.QueryRow("SELECT * FROM directory WHERE directorypath = ?", path).Scan(&file.ID, &file.Path, &file.Permission)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (db *Database) GetUserPermission(username string) (int, error) {
	user, err := db.getUserInfo(username)
	if err != nil {
		return 0, err
	}
	return user.Permission, nil
}

func (db *Database) GetFilePermission(path string) (int, error) {
	file, err := db.getFileInfo(path)
	if err != nil {
		return 0, err
	}
	return file.Permission, nil
}
