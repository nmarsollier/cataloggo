package main

import (
	"github.com/nmarsollier/cataloggo/internal/engine/di"
	"github.com/nmarsollier/cataloggo/internal/engine/env"
	"github.com/nmarsollier/cataloggo/internal/engine/log"
	server "github.com/nmarsollier/cataloggo/internal/graph"
	"github.com/nmarsollier/cataloggo/internal/rabbit/consume"
	"github.com/nmarsollier/cataloggo/internal/rest"
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
	deps := di.NewInjector(log.Get(env.Get().FluentUrl))

	go server.Start(deps.Logger())
	consume.Init(deps.Logger(), deps.ArticleExistConsumer(), deps.LogoutConsumer(), deps.OrderPlacedConsumer())
	rest.Start()
}
