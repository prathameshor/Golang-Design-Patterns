package guns

type AssaultRifle struct {
	gun_name string
}

func (a AssaultRifle) GunName() string {
	return a.gun_name
}

func NewAssaultRifle() GunFactory {
	return AssaultRifle{
		gun_name: "AK47",
	}
}
