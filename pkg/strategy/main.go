package main


func main(){
	factory := NewCharacterFactory()
	king := factory.createCharacter("king")
	king.setName("king hg")
	king.setWeapon(&AxeBehavior{})
	king.fight()

	queen := factory.createCharacter("queen")
	queen.setName("queen hg")

	queen.setWeapon(&KnifeBehavior{})
	queen.fight()
}
