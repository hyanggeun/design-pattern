package main

type CharacterFactory struct {
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{}
}

func (f *CharacterFactory) createCharacter(charType string) CharacterIf {
	switch charType {
	case "king":
		return &King{}
	case "queen":
		return &Queen{}
	}
	return nil
}

