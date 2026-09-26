package guns

type MachineGun struct {
	gun_name string
}

func (m MachineGun) GunName() string {
	return m.gun_name
}

func NewMachineGun() GunFactory {
	return MachineGun{
		gun_name: "M60",
	}
}
