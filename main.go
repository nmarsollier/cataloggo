package main

import (
	"github.com/nmarsollier/cataloggo/rabbit/consume"
	routes "github.com/nmarsollier/cataloggo/rest"
)

//	@title			CatalogGo
//	@version		1.0
//	@description	Microservicio de Catalogo.
//	@contact.name	Nestor Marsollier
//	@contact.email	nmarsollier@gmail.com
//
//	@host			localhost:3002
//	@BasePath		/v1
//
// Main Method
func main() {
	consume.Init()
	routes.Start()
}
