package tools

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/nmarsollier/cataloggo/tools/log"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func newLogger(ctx context.Context, deps ...interface{}) *logrus.Entry {
	operationContext := graphql.GetOperationContext(ctx)

	return log.Get(deps...).
		WithField(log.LOG_FIELD_CORRELATION_ID, getCorrelationId(ctx)).
		WithField(log.LOG_FIELD_CONTROLLER, "Rest").
		WithField(log.LOG_FIELD_HTTP_METHOD, operationContext.OperationName).
		WithField(log.LOG_FIELD_HTTP_PATH, operationContext.OperationName)
}

func gqlLogger(ctx context.Context) *logrus.Entry {
	operationContext := graphql.GetOperationContext(ctx)

	logger, exist := operationContext.Variables["logger"]
	if !exist {
		return newLogger(ctx)
	}
	return logger.(*logrus.Entry)
}

func getCorrelationId(ctx context.Context) string {
	operationContext := graphql.GetOperationContext(ctx)
	value := operationContext.Headers.Get("Authorization")

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}
