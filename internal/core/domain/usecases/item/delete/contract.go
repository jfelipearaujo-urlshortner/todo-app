package delete_contract

type UseCase interface {
	Execute(id string) error
}
