package rest

import (
	"fmt"

	"github.com/nmarsollier/cataloggo/rest/engine"
	"github.com/nmarsollier/cataloggo/tools/env"
)

// Start this server
func Start() {
	engine.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
}
