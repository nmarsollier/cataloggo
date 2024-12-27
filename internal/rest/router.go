package rest

import (
	"fmt"

	"github.com/nmarsollier/cataloggo/internal/engine/env"
	"github.com/nmarsollier/cataloggo/internal/rest/engine"
)

// Start this server
func Start() {
	engine.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
}
