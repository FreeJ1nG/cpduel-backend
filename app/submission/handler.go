package submission

type handler struct{}

type Handler interface{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SubmitCode() {

}
