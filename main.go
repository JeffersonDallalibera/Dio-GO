package main
import "fmt"



func kelvinParaCelsius(temperaturaKelvin float64) float64 {
    return temperaturaKelvin - 273
}

func main() {
	var kelvin float64
	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scan(&kelvin)

	celsius := kelvinParaCelsius(kelvin)

    fmt.Printf("A temperatura em Celsius é: %.2f\n", celsius)
}
