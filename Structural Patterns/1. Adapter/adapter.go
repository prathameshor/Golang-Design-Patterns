package main

type Adapter struct {
	ser Service
}

func (a *Adapter) Calculate() int {
	return int(a.ser.CalculateFloat()) 
}
