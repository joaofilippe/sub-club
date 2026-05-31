package module

type ModuleDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"createdAt"`
}

type ModuleInputDTO struct {
	Name string `json:"name"`
}

type UpdateModuleInputDTO struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
