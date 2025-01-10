package server

func (s *Server) registerRoutes() {
	api := s.App.Group("/api/v1")

	api.Post("/items/:id/complete", s.container.CompleteItemController.Handle)
	api.Post("/items", s.container.CreateItemController.Handle)
	api.Delete("/items/:id", s.container.DeleteItemController.Handle)
	api.Get("/items/:id", s.container.GetItemController.Handle)
	api.Get("/items", s.container.ListItemController.Handle)
	api.Put("/items", s.container.UpdateItemController.Handle)
}
