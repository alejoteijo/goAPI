package sample

import (
	pet "goAPI/pkg"
)

var Pets = map[string]*pet.Pet{
	"64ASDF64ASD64F55A43GA4": &pet.Pet{
		ID:    "64ASDF64ASD64F55A43GA4",
		Name:  "Laika",
		Image: "https://en.wikipedia.org/wiki/Laika#/media/File:Laika_(Soviet_dog).jpg",
		Age:   3,
	},

	"ASDF6757ASD4F3ASDFA34": &pet.Pet{
		ID:    "ASDF6757ASD4F3ASDFA34",
		Name:  "Rex",
		Image: "https://upload.wikimedia.org/wikipedia/commons/6/6e/Megan_Leavey_and_Rex.jpg",
		Age:   5,
	},
}
