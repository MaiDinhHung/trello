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

	api.Get("/users", controllers.GetAllUsers)

	api.Get("/boards/:id", controllers.GetBoardByID)
	api.Get("/boards/:id/members", controllers.GetBoardMembers)

	api.Post("/lists", controllers.CreateList)
	api.Get("/boards/:id/lists", controllers.GetListsByBoardID)
	api.Put("/lists/:id", controllers.UpdateList)
	api.Delete("/lists/:id", controllers.DeleteList)

	api.Get("/cards", controllers.GetAllCards)
	api.Post("/cards", controllers.CreateCard)
	api.Put("/cards/:id", controllers.UpdateCard)
	api.Get("/lists/:id/cards", controllers.GetCardsByListID)
	api.Delete("/cards/:id", controllers.DeleteCard)

	api.Post("/cards/:card_id/comments", controllers.CreateComment)
	api.Get("/cards/:card_id/comments", controllers.GetComments)
	api.Put("/cards/:card_id/comments/:id", controllers.UpdateComment)
	api.Delete("/cards/:card_id/comments/:id", controllers.DeleteComment)


	api.Get("/cards/:card_id/checklist", controllers.GetChecklistItems)
	api.Post("/cards/:card_id/checklist", controllers.CreateChecklistItem)
	api.Put("/checklist/:id", controllers.UpdateChecklistItem)
	api.Delete("/checklist/:id", controllers.DeleteChecklistItem)

	api.Post("/cards/:card_id/members", controllers.AddMemberToCard)
	api.Get("/cards/:card_id/members", controllers.GetCardMembers)
	api.Delete("/cards/:card_id/members/:user_id", controllers.RemoveMemberFromCard)
}
