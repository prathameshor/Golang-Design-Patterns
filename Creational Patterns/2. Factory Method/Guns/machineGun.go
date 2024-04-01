package guns

type MachineGun struct {
	gun_name string
}

func (m *MachineGun) Name() string {
	return m.gun_name
}

func NewMachineGun() MachineGun {
	return MachineGun{
		gun_name: "M60",
	}
}

// func(m *MachineGun) IWINegev(){}

// func(m *MachineGun) M60(){}
