package security

import (
	"github.com/nmarsollier/cataloggo/tools/log"
)

// Invalidate invalida un token del cache
func Invalidate(token string, ctx ...interface{}) {
	if len(token) <= 7 {
		log.Get(ctx...).Info("Token no valido: ", token)
		return
	}

	cache.Delete(token[7:])
}