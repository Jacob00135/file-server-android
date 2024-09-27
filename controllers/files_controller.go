package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v3"

	db "github.com/Jacob00135/file-server-android/database"
	"github.com/Jacob00135/file-server-android/models"
)

func UploadFile(c fiber.Ctx) error {
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

func DownloadFile(c fiber.Ctx) error {
	// 获取文件名
	filename := c.Params("path")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Filename is required",
		})
	}

	// 获取路径
	path := filepath.Clean(c.Query("visible_dir"))

	// 检查文件是否存在 && judge no dir
	fullPath := filepath.Join(path, filename)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "File not found",
		})
	}

	// 发送文件
	return c.SendFile(fullPath)
}

func ListFiles(c fiber.Ctx) error {
	userPermission := c.Locals("userPermission").(uint)
	dir := filepath.Clean(c.Query("visible_dir"))

	if dir != "." {
		target := c.Locals("target").(string)
		targeInfo, _ := os.Stat(target)
		if !targeInfo.IsDir() {
			return c.SendFile(target)
		} else {
			files, err := getFilesByPath(target)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
					"code":    fiber.StatusInternalServerError,
					"message": fmt.Sprintf("Could not read directory: %v", err.Error()),
				})
			}

			return c.JSON(fiber.Map{
				"father": target,
				"files":  files,
			})
		}
	} else {
		files, err := db.DB.GetDirsByPermission(userPermission)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"father": "",
			"files":  files,
		})
	}
}

// func ListIndex(c fiber.Ctx) error {
// 	// auth check
// 	sess, err := db.Storage.Get(c)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Could not Get session",
// 		})
// 	}

// 	var username string
// 	if name := sess.Get("username"); name != nil {
// 		username = name.(string)
// 	}

// 	var userPermission uint = 1
// 	if username != "" {
// 		var err error
// 		userPermission, err = db.DB.GetUserPermission(username)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"message": err.Error(),
// 			})
// 		}
// 	}

// 	dir := filepath.Clean(c.Query("visible_dir"))

// 	if dir != "." {
// 		// check dir auth
// 		dirPermission, err := db.DB.GetFilePermission(dir)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"message": err.Error(),
// 			})
// 		}
// 		if userPermission < dirPermission {
// 			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 				"message": "visible_dir permission denied",
// 			})
// 		}

// 		// check path
// 		path := filepath.Clean(c.Query("path"))
// 		if path == "." {
// 			if !osFileExists(dir) {
// 				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 					"message": "Path not found",
// 				})
// 			}

// 			files, err := getFilesByPath(dir)
// 			if err != nil {
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 					"message": fmt.Sprintf("Could not read directory: %v", err.Error()),
// 				})
// 			}

// 			return c.JSON(fiber.Map{
// 				"father": dir,
// 				"files":  files,
// 			})
// 		} else {
// 			target := filepath.Join(dir, path)
// 			if !securePath(dir, target) {
// 				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 					"message": "Invalid path",
// 				})
// 			}
// 			if !osFileExists(target) {
// 				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 					"message": "Path not found",
// 				})
// 			}

// 			targeInfo, _ := os.Stat(target)
// 			if !targeInfo.IsDir() {
// 				return c.SendFile(target)
// 			} else {
// 				files, err := getFilesByPath(target)
// 				if err != nil {
// 					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 						"message": fmt.Sprintf("Could not read directory: %v", err.Error()),
// 					})
// 				}

// 				return c.JSON(fiber.Map{
// 					"father": target,
// 					"files":  files,
// 				})
// 			}
// 		}
// 	} else {
// 		files, err := db.DB.GetDirsByPermission(userPermission)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"message": err.Error(),
// 			})
// 		}

// 		return c.JSON(fiber.Map{
// 			"father": "",
// 			"files":  files,
// 		})
// 	}
// }

// func osFileExists(path string) bool {
// 	_, err := os.Stat(path)
// 	return !os.IsNotExist(err)
// }

func getFilesByPath(path string) ([]models.File, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []models.File
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		files = append(files, models.File{
			Name:     info.Name(),
			FileType: info.IsDir(),
			FileSize: info.Size(),
		})
	}
	return files, nil
}

// func insertFileTable(filePath string) (id int, permission int, err error) {
// 	res, err := db.DB.Conn.Exec("INSERT INTO directory (directorypath) VALUES (?)", filePath)
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

func InsertFileTableWithP(filePath string, p int) (id int, permission int, err error) {
	res, err := db.DB.Conn.Exec("INSERT INTO directory (directorypath, permission) VALUES (?, ?)", filePath, p)
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
