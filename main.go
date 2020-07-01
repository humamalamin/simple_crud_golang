package main

func main() {
	a := App{}

	// Configurasi database disini
	a.Initialize("root", "root", "kum_astra")

	a.Run(":8080")
}
