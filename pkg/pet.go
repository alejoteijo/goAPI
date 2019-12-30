package pkg

type Pet struct {
	ID string `json:"ID"`
	Name string `json:"name, omitempty"`
	Image string `json:"image,omitempty"`
	Age int `json:"age,omitempty"`
}

type PetRepository interface {
	CreatePet(p *Pet) error
	FetchPets() ([]*Pet, error)
	DeletePet(ID string) error
	UpdatePet(ID string, p *Pet) error
	FetchPetByID(ID string) (*Pet, error)
}