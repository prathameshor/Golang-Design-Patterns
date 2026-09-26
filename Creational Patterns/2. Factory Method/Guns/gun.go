package guns

type GunFactory interface {
	GunName() string
}

func GiveMeGun(g GunFactory) string {
	return g.GunName()
}
