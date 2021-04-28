package main

type WeaponBehavior interface {
	useWeapon() string
}

type KnifeBehavior struct {
}
func (b *KnifeBehavior) useWeapon() string{
	return "knife"
}

type AxeBehavior struct {
}

func (b *AxeBehavior) useWeapon() string{
	return "axe"
}


