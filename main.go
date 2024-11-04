package main

import (
  "fmt"
  "flag"
)

// Busca por uma pasta .git no diretório
func scan(path string) {
  fmt.Println("Scannig", path)
}

// Gera os status de contribuição
func stats(email string) {
  fmt.Println("Stats de ", email)
}
func main() {
  var email string
  var folder string

  flag.StringVar(&email, "email", "seuemail@email.com", "Email do qual será retirado as contribuições")
  flag.StringVar(&folder, "add", "", "Pasta na qual será buscado um diretório .git")

  flag.Parse()

  if folder != "" {
    scan(folder)
  } else {
    stats(email)
  }

  fmt.Println("Fim do arquivo")
}
