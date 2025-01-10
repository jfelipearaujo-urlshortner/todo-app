package complete_controller

type Request struct {
	ID string `json:"id" validate:"required,uuid"`
}
