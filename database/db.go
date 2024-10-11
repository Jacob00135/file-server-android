package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"

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
	fileHome := os.Getenv("FILE_HOME")

	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	Storage = session.New()
	DB = &Database{Conn: conn}

	// if db exits return
	if _, err := os.Stat(dbPath); err == nil {
		return
	}

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
		{"red", "red", 2},
		{"blue", "blue", 2},
		{"green", "green", 2},
		{"fakeadmin", "admin", 2},
	}
	for _, user := range testUser {
		// _, err = DB.Conn.Exec("INSERT OR IGNORE INTO users (username, password, permission) VALUES (?, ?, ?)", user.username, user.password, user.permission)
		err := DB.InsertUser(user.username, user.password, uint(user.permission))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = DB.InsertDir(fileHome, 1)
	if err != nil {
		log.Fatal(err)
	}

}

func (db *Database) Connect(dbPath string) {
	if db.Conn != nil {
		return
	}
	conn, err := sql.Open("sqlite3", dbPath)

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

// func (db *Database) InsertUser(username, password string, p uint) error {
// 	if db.Conn == nil {
// 		return errors.New("database connection is not open")
// 	}
// 	_, err := db.Conn.Exec("INSERT INTO users (username, password, permission) VALUES (?, ?, ?)", username, password, p)
// 	return err
// }

func (db *Database) InsertUser(username, password string, p uint) error {
	if db.Conn == nil {
		return errors.New("database connection is not open")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	res, err := db.Conn.Exec("INSERT OR IGNORE INTO users (username, password, permission) VALUES (?, ?, ?)", username, string(hashedPassword), p)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected > 0 {
		return nil
	}

	return fmt.Errorf("duplicate username: %s", username)
}

func (db *Database) DeleteUserByName(username string) error {
	if db.Conn == nil {
		return errors.New("database connection is not open")
	}
	res, err := db.Conn.Exec("DELETE FROM users WHERE username = ? AND username != 'admin'", username)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected > 0 {
		return nil
	}
	return fmt.Errorf("delete %s faild", username)
}

func (db *Database) DeleteUserById(id uint) error {
	if db.Conn == nil {
		return errors.New("database connection is not open")
	}
	res, err := db.Conn.Exec("DELETE FROM users WHERE userid = ? AND username != 'admin'", id)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected > 0 {
		return nil
	}
	return fmt.Errorf("delete id:%d faild", id)
}

func (db *Database) UpdateUser(username, password string, p uint) error {
	if db.Conn == nil {
		return errors.New("database connection is not open")
	}
	// check pwd in db
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	res, err := db.Conn.Exec("UPDATE users SET password = ?, permission = ? WHERE username = ?", hashedPassword, p, username)
	if err != nil {
		return err
	}
	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected > 0 {
		return nil
	}
	return fmt.Errorf("update %s faild", username)
}

func (db *Database) InsertDir(path string, p uint) error {
	if db.Conn == nil {
		return errors.New("database connection is not open")
	}
	path = filepath.Clean(path)
	_, err := db.Conn.Exec("INSERT OR IGNORE INTO directory (directorypath, permission) VALUES (?, ?)", path, p)
	return err
}

func (db *Database) GetDirsByPermission(p uint) ([]models.File, error) {
	if db.Conn == nil {
		return nil, errors.New("database connection is not open")
	}
	rows, err := db.Conn.Query("SELECT directorypath FROM directory WHERE permission <= ?", p)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	files := make([]models.File, 0)
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			return nil, err
		}
		files = append(files, models.File{
			Name:     path,
			FileType: true,
			FileSize: 0,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return files, nil
}

func (db *Database) CheckUserExists(username, password string) (bool, error) {
	if db.Conn == nil {
		return false, errors.New("database connection is not open")
	}

	var hashedPassword string
	err := db.Conn.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil // user not found
	} else if err != nil {
		return false, err // database error
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == nil {
		return true, nil // password correct
	} else if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil // password incorrect
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

func (db *Database) GetUserInfo(username string) (*models.User, error) {
	if db.Conn == nil {
		return nil, errors.New("database connection is not open")
	}
	user := new(models.User)
	err := db.Conn.QueryRow("SELECT userid, username, permission FROM users WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.Permission)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *Database) GetFileInfo(path string) (*models.DbFile, error) {
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

func (db *Database) GetUserPermission(username string) (uint, error) {
	user, err := db.GetUserInfo(username)
	if err != nil {
		return 0, err
	}
	return user.Permission, nil
}

func (db *Database) GetFilePermission(path string) (uint, error) {
	file, err := db.GetFileInfo(path)
	if err != nil {
		return 0, err
	}
	return file.Permission, nil
}

func (db *Database) GetAllUsers() ([]models.User, error) { // only name and permission
	if db.Conn == nil {
		log.Fatal("database connection is not open")
	}

	rows, err := db.Conn.Query("SELECT userid, username, permission FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Permission); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *Database) GetAllDir() ([]models.DbFile, error) {
	if db.Conn == nil {
		log.Fatal("database connection is not open")
	}

	rows, err := db.Conn.Query("SELECT * FROM directory")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	files := make([]models.DbFile, 0)
	for rows.Next() {
		var file models.DbFile
		if err := rows.Scan(&file.ID, &file.Path, &file.Permission); err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func (db *Database) DeleteDirById(id uint) error {
	if db.Conn == nil {
		return errors.New("database connection is not open")
	}
	res, err := db.Conn.Exec("DELETE FROM users WHERE userid = ? AND username != 'admin'", id)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected > 0 {
		return nil
	}
	return fmt.Errorf("delete id:%d faild", id)
}

func (db *Database) UpdateDir(id, p uint, path string) error {
	if db.Conn == nil {
		return errors.New("database connection is not open")
	}
	res, err := db.Conn.Exec("UPDATE directory SET permission = ?, directorypath = ? WHERE directoryid = ?", p, path, id)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected > 0 {
		return nil
	}

	return fmt.Errorf("update id:%d faild", id)
}
