package routes

import (
	"news/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/news/add", controllers.NewsAdd)
	app.Post("/api/news/edit", controllers.NewsEdit)
	app.Delete("/api/news/delete", controllers.NewsDelete)
	app.Get("/api/news/list/:status/:topic", controllers.NewsList)

	app.Post("/api/tag/add", controllers.TagAdd)
	app.Post("/api/tag/edit", controllers.TagEdit)
	app.Delete("/api/tag/delete", controllers.TagDelete)
	app.Get("/api/tag/list", controllers.TagList)
}
