package routes
import (
	"trello/controllers"
	"github.com/gofiber/fiber/v2"
	
)
func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    auth := api.Group("/auth")
    auth.Post("/register", controllers.Register)
    auth.Post("/login", controllers.Login)

	api.Get("/boards/:id", controllers.GetBoardByID)

	api.Post("/lists", controllers.CreateList)
	api.Get("/boards/:id/lists", controllers.GetListsByBoardID)
	api.Put("/lists/:id", controllers.UpdateList)
	api.Delete("/lists/:id", controllers.DeleteList)


	api.Post("/cards", controllers.CreateCard)
	api.Put("/cards/:id", controllers.UpdateCard)
	api.Get("/lists/:id/cards", controllers.GetCardsByListID)

	api.Post("/cards/:card_id/comments", controllers.CreateComment)
	api.Get("/cards/:card_id/comments", controllers.GetComments)
}