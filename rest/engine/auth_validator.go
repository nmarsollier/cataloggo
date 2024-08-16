package engine

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/nmarsollier/cataloggo/security"
	"github.com/nmarsollier/cataloggo/tools/errs"
)

// ValidateAuthentication validate gets and check variable body to create new variable
// puts model.Variable in context as body if everything is correct
func ValidateAuthentication(c *gin.Context) {
	if err := validateToken(c); err != nil {
		c.Error(err)
		c.Abort()
		return
	}
}

var securityValidate func(token string) (*security.User, error) = security.Validate

func validateToken(c *gin.Context) error {
	tokenString, err := GetHeaderToken(c)
	if err != nil {
		glog.Error(err)
		return errs.NotFound
	}

	if _, err = securityValidate(tokenString); err != nil {
		glog.Error(err)
		return errs.Invalid
	}

	return nil
}

// get token from Authorization header
func GetHeaderToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "bearer ") != 0 {
		glog.Error(errs.Unauthorized)

		return "", errs.Unauthorized
	}
	return tokenString[7:], nil
}