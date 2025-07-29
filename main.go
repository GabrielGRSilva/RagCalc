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
	var dificuldade int
	fmt.Print("Qual a dificuldade para preparar a poção desejada? ")
	fmt.Scan(&dificuldade)

	chance := float64(stats.nvpesquisa) +
		(float64(stats.nvclasse) * 0.2) +
		(float64(stats.dex) * 0.1) +
		(float64(stats.sor) * 0.1) +
		(float64(stats.inteligencia) * 0.05) +
		float64(dificuldade)

	fmt.Printf("A chance de preparar esta poção é de %.2f\n", chance)
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
