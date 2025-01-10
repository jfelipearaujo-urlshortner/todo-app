package list_usecase

import (
	"sort"

	"github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/entities"
	item_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/repositories/item"
	list_contract "github.com/jfelipearaujo-urlshortner/todo-app/internal/core/domain/usecases/item/list"
)

type useCase struct {
	repository item_contract.Repository
}

func New(repository item_contract.Repository) list_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (u *useCase) Execute() []entities.Item {
	items := u.repository.List()

	if len(items) == 0 {
		return []entities.Item{}
	}

	sort.Sort(ItemSorter(items))

	return items
}

type ItemSorter []entities.Item

func (s ItemSorter) Len() int {
	return len(s)
}

func (s ItemSorter) Less(i, j int) bool {
	return s[i].CreatedAt.Before(s[j].CreatedAt)
}

func (s ItemSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
