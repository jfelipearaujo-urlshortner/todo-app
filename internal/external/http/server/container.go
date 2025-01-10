package server

import (
	"github.com/go-playground/validator/v10"
	complete_usecase "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/application/usecases/item/complete"
	create_usecase "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/application/usecases/item/create"
	delete_usecase "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/application/usecases/item/delete"
	get_usecase "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/application/usecases/item/get"
	list_usecase "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/application/usecases/item/list"
	update_usecase "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/application/usecases/item/update"
	item_repository "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/infrastructure/repositories/item"
	complete_controller "github.com/jfelipearaujo-urlshortner/todo-app/internal/external/http/controllers/item/complete"
	create_controller "github.com/jfelipearaujo-urlshortner/todo-app/internal/external/http/controllers/item/create"
	delete_controller "github.com/jfelipearaujo-urlshortner/todo-app/internal/external/http/controllers/item/delete"
	get_controller "github.com/jfelipearaujo-urlshortner/todo-app/internal/external/http/controllers/item/get"
	list_controller "github.com/jfelipearaujo-urlshortner/todo-app/internal/external/http/controllers/item/list"
	update_controller "github.com/jfelipearaujo-urlshortner/todo-app/internal/external/http/controllers/item/update"
)

type container struct {
	CompleteItemController *complete_controller.Controller
	CreateItemController   *create_controller.Controller
	DeleteItemController   *delete_controller.Controller
	GetItemController      *get_controller.Controller
	ListItemController     *list_controller.Controller
	UpdateItemController   *update_controller.Controller
}

func newContainer() *container {
	validator := validator.New()

	repository := item_repository.New()

	completeUseCase := complete_usecase.New(repository)
	createUseCase := create_usecase.New(repository)
	deleteUseCase := delete_usecase.New(repository)
	getUseCase := get_usecase.New(repository)
	listUseCase := list_usecase.New(repository)
	updateUseCase := update_usecase.New(repository)

	return &container{
		CompleteItemController: complete_controller.New(validator, completeUseCase),
		CreateItemController:   create_controller.New(validator, createUseCase),
		DeleteItemController:   delete_controller.New(validator, deleteUseCase),
		GetItemController:      get_controller.New(validator, getUseCase),
		ListItemController:     list_controller.New(validator, listUseCase),
		UpdateItemController:   update_controller.New(validator, updateUseCase),
	}
}
