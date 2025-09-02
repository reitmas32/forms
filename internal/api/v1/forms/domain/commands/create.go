package commands

type CreateFormCommand struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (c CreateFormCommand) Validate() error {
	return nil
}
