package main

import (
	"fmt"
	"math/rand"
)

type CharStats struct {
	nvpesquisa   int
	nvclasse     int
	dex          int
	sor          int
	nvbase       int
	nvprotquim   int
	inteligencia int
}

func checarStats() CharStats {
	var stats CharStats

	fmt.Print("Seu nível de pesquisa: ")
	fmt.Scan(&stats.nvpesquisa)
	fmt.Print("Seu nível de classe: ")
	fmt.Scan(&stats.nvclasse)
	fmt.Print("Sua destreza: ")
	fmt.Scan(&stats.dex)
	fmt.Print("Sua sorte: ")
	fmt.Scan(&stats.sor)
	fmt.Print("Seu nível base: ")
	fmt.Scan(&stats.nvbase)
	fmt.Print("Seu nível de proteção química: ")
	fmt.Scan(&stats.nvprotquim)
	fmt.Print("Sua inteligência: ")
	fmt.Scan(&stats.inteligencia)

	return stats
}

func prepararPoções(stats CharStats) {
	var tipopocao int
	var dificuldade int
	fmt.Print("Qual poção voce quer preparar?\n1 - Vermelha, amarela ou branca\n2 - Alcool\n3 - Acido, Esfera Marinha, Fogo Grego ou Planta Carnivora\n4 - Azul, Analgesico, Aloe Vera, Embriao, Pocoes Anti-Propriedade ou Pocao Compacta Vermelha\n 5 - Compacta Amarela\n6 - Compacta Branca ou Revestimento")
	fmt.Scan(&tipopocao)

	switch tipopocao {
	case 1:
		dificuldade = rand.Intn(11) + 15
	case 2:
		dificuldade = rand.Intn(11) + 5
	case 3:
		dificuldade = rand.Intn(11) - 5
	case 4:
		dificuldade = -5
	case 5:
		dificuldade = rand.Intn(6) - 10
	case 6:
		dificuldade = rand.Intn(11) - 15
	default:
		fmt.Print("Opcao invalida!")

		chance := float64(stats.nvpesquisa) +
			(float64(stats.nvclasse) * 0.2) +
			(float64(stats.dex) * 0.1) +
			(float64(stats.sor) * 0.1) +
			(float64(stats.inteligencia) * 0.05) +
			float64(dificuldade)

		fmt.Printf("A chance de preparar esta poção é de %.2f\n", chance)
	}
}

func farmacologiaAvançada(stats CharStats) {
	randomnum := rand.Intn(120) + 30
	randomsmall := rand.Intn(6) + 4

	chance := stats.inteligencia +
		(stats.dex / 2) +
		stats.sor +
		stats.nvclasse +
		randomnum +
		(stats.nvbase - 100) +
		(stats.nvpesquisa * 5) +
		(stats.nvprotquim * randomsmall)

	fmt.Printf("A chance de Farmacologia Avançada é de %d\n", chance)
}

func main() {
	var opcao int
	fmt.Println("Bem-vindo(a) ao RagCalc!\nDigite as seguintes informações:")
	stats := checarStats()

	for {
		fmt.Println("\nO que você quer fazer agora?")
		fmt.Println("1 - Preparar Poção")
		fmt.Println("2 - Farmacologia Avançada")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha uma opção: ")

		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			prepararPoções(stats)
		case 2:
			farmacologiaAvançada(stats)
		case 3:
			fmt.Println("Boa sorte, aventureiro! Obrigado por usar o RagCalc!")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
