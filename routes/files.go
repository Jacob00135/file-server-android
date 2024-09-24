package routes

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v3"

	"github.com/Jacob00135/file-server-android/database"
)

func SetupFileRoutes(app *fiber.App) {
	// 上传文件
	app.Post("/upload", uploadFile)

	// 下载文件
	app.Get("/download/:filename", downloadFile)

	app.Get("/api/index", ListFilesNew)
}

func uploadFile(c fiber.Ctx) error {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// 获取路径
	path := filepath.Clean(c.Query("path"))

	// 检查路径是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Path not found",
		})
	}

	// 保存文件到指定路径
	err = c.SaveFile(file, filepath.Join(path, file.Filename))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "File uploaded successfully",
	})
}

func downloadFile(c fiber.Ctx) error {
	// 获取文件名
	filename := c.Params("filename")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Filename is required",
		})
	}

	// 获取路径
	path := filepath.Clean(c.Query("path"))

	// 检查文件是否存在
	fullPath := filepath.Join(path, filename)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "File not found",
		})
	}

	// 发送文件
	return c.SendFile(fullPath)
}

func ListFilesNew(c fiber.Ctx) error {

	// auth check
	sess, err := database.Storage.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not Get session",
		})
	}

	var username string
	if name := sess.Get("username"); name != nil {
		username = name.(string)
	}

	userPermission := 1
	if username != "" {
		var err error
		userPermission, err = database.DB.GetUserPermission(username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	files, err := database.DB.GetDirsByPermission(userPermission)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"files": files,
	})
}

func checkFileExists(path string) (bool, error) {
	_, err := database.DB.Conn.Query("SELECT 1 FROM directory WHERE directorypath = ?", path)
	if err == nil {
		return true, nil
	} else if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return false, err
}

// func insertFileTable(filePath string) (id int, permission int, err error) {
// 	res, err := database.DB.Conn.Exec("INSERT INTO directory (directorypath) VALUES (?)", filePath)
// 	if err != nil {
// 		return id, permission, err
// 	}
// 	id64, err := res.LastInsertId()
// 	if err != nil {
// 		return id, permission, err
// 	}
// 	id = int(id64)
// 	permission = 4
// 	return id, permission, nil
// }

func insertFileTableWithP(filePath string, p int) (id int, permission int, err error) {
	res, err := database.DB.Conn.Exec("INSERT INTO directory (directorypath, permission) VALUES (?, ?)", filePath, p)
	if err != nil {
		return id, permission, err
	}
	id64, err := res.LastInsertId()
	if err != nil {
		return id, permission, err
	}
	id = int(id64)
	permission = p
	return id, permission, nil
}
