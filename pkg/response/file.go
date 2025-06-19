package response

import "github.com/gofiber/fiber/v2"

func FileResponse(c *fiber.Ctx, path string) error {
	return c.SendFile(path)
}

func DownloadFileResponse(c *fiber.Ctx, path, filename string) error {
	return c.Download(path, filename)
}
