package storage

import (
	"fmt"
	pet "goAPI/pkg"
	"sync"
)

type petRepository struct {
	mtx	sync.RWMutex
	pets map[string]*pet.Pet
}

func NewPetRepository(pets map[string]*pet.Pet) pet.PetRepository {
	if pets == nil{
		pets = make(map[string]*pet.Pet)
	}

	return &petRepository{
		pets: pets,
	}
}

func (r *petRepository) CreatePet(p *pet.Pet) error  {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if err := r.checkIfExists(p.ID); err != nil {
		return err
	}
	r.pets[p.ID] = p
	return nil
}

func (r *petRepository) FetchPets() ([]*pet.Pet, error)  {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]*pet.Pet, 0, len(r.pets))
	for _, value := range r.pets{
		values = append(values,value)
	}
	return values, nil
}

func (r *petRepository) DeletePet(ID string) error  {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.pets, ID)
	return nil
}

func (r *petRepository) UpdatePet(ID string, p *pet.Pet) error  {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.pets[ID] = p
	return nil
}

func (r *petRepository) FetchPetByID(ID string) (*pet.Pet, error)  {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	for _, v := range r.pets{
		if v.ID == ID{
			return v, nil
		}
	}
	return nil, fmt.Errorf("The ID %s doesn't exist", ID)
}

func (r *petRepository) checkIfExists(id string) error {
	for _, v := range  r.pets{
		if v.ID == id {
			return fmt.Errorf("The pet %s is already exist",id)
		}
	}	
	return nil
}
