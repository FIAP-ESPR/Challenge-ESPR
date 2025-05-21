package dao

import "ancora/model"

func GetProducts() ([]model.Product, error) {
	return []model.Product{
		{ID: 12, Name: "Pastilha de Freio", Price: 89.90, Image: "peca1.jpg"},
		{ID: 3, Name: "Amortecedor Dianteiro", Price: 220.00, Image: "peca1.jpg"},
		{ID: 27, Name: "Filtro de Ar", Price: 45.00, Image: "peca1.jpg"},
		{ID: 8, Name: "Bateria 60Ah", Price: 430.00, Image: "peca1.jpg"},
		{ID: 21, Name: "Correia Dentada", Price: 110.00, Image: "peca1.jpg"},
		{ID: 5, Name: "Pneu Aro 15", Price: 350.00, Image: "peca1.jpg"},
		{ID: 30, Name: "Velas de Ignição", Price: 65.00, Image: "peca1.jpg"},
		{ID: 16, Name: "Disco de Freio", Price: 140.00, Image: "peca1.jpg"},
		{ID: 1, Name: "Retrovisor Elétrico", Price: 280.00, Image: "peca1.jpg"},
		{ID: 24, Name: "Farol Direito", Price: 520.00, Image: "peca1.jpg"},
		{ID: 18, Name: "Lanterna Traseira", Price: 150.00, Image: "peca1.jpg"},
		{ID: 10, Name: "Sensor de Estacionamento", Price: 120.00, Image: "peca1.jpg"},
		{ID: 7, Name: "Reservatório de Água", Price: 60.00, Image: "peca1.jpg"},
		{ID: 29, Name: "Coxim do Motor", Price: 130.00, Image: "peca1.jpg"},
		{ID: 13, Name: "Radiador", Price: 470.00, Image: "peca1.jpg"},
		{ID: 2, Name: "Parachoque Dianteiro", Price: 590.00, Image: "peca1.jpg"},
		{ID: 20, Name: "Filtro de Óleo", Price: 35.00, Image: "peca1.jpg"},
		{ID: 6, Name: "Comando de Setas", Price: 180.00, Image: "peca1.jpg"},
		{ID: 25, Name: "Módulo de Injeção", Price: 980.00, Image: "peca1.jpg"},
		{ID: 11, Name: "Braço Axial", Price: 95.00, Image: "peca1.jpg"},
		{ID: 4, Name: "Capô", Price: 890.00, Image: "peca1.jpg"},
		{ID: 22, Name: "Caixa de Direção", Price: 1150.00, Image: "peca1.jpg"},
		{ID: 9, Name: "Motor de Arranque", Price: 670.00, Image: "peca1.jpg"},
		{ID: 28, Name: "Cilindro Mestre", Price: 240.00, Image: "peca1.jpg"},
		{ID: 15, Name: "Interruptor Luz Freio", Price: 30.00, Image: "peca1.jpg"},
		{ID: 26, Name: "Bobina de Ignição", Price: 160.00, Image: "peca1.jpg"},
		{ID: 19, Name: "Filtro de Combustível", Price: 38.00, Image: "peca1.jpg"},
		{ID: 14, Name: "Cabo de Embreagem", Price: 55.00, Image: "peca1.jpg"},
		{ID: 17, Name: "Ventoinha do Radiador", Price: 195.00, Image: "peca1.jpg"},
		{ID: 23, Name: "Junta do Cabeçote", Price: 75.00, Image: "peca1.jpg"},
	}, nil
}
