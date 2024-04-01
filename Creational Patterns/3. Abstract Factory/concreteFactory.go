package main

type Modern struct {
}

func (m *Modern) createChair() IChair {
	return &ModernChair{
		Chair: Chair{
			Furniture: "Chair",
			Type:      "Modern",
		},
	}
}

func (m *Modern) createSofa() ISofa {
	return &Sofa{
		Furniture: "Sofa",
		Type:      "Modern",
	}
}

type Victorian struct {
}

func (m *Victorian) createChair() IChair {
	return &Chair{
		Furniture: "Chair",
		Type:      "Victorian",
	}
}

func (m *Victorian) createSofa() ISofa {
	return &Sofa{
		Furniture: "Sofa",
		Type:      "Victorian",
	}
}
