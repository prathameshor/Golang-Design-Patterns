package guns

type AssaultRifle struct {
	gun_name string
}

func (a *AssaultRifle) Name() string {
	return a.gun_name
}

func NewAssaultRifle() AssaultRifle {
	return AssaultRifle{
		gun_name: "AK47",
	}
}

// func (a *AssaultRifle) M416() {}

// func (a *AssaultRifle) Sig716() {}
