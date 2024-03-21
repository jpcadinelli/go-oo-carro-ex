package carros

type carro struct {
	Marca string
	Modelo string
	Ano int
	Cor string
	VelocidadeAtual float64
}

func (car *carro) Acelerar(quantidade float64) {
	car.VelocidadeAtual += quantidade
}