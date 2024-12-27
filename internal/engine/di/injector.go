package di

import (
	"github.com/nmarsollier/cataloggo/internal/article"
	"github.com/nmarsollier/cataloggo/internal/engine/db"
	"github.com/nmarsollier/cataloggo/internal/engine/env"
	"github.com/nmarsollier/cataloggo/internal/engine/httpx"
	"github.com/nmarsollier/cataloggo/internal/engine/log"
	"github.com/nmarsollier/cataloggo/internal/rabbit/consume"
	"github.com/nmarsollier/cataloggo/internal/rabbit/emit"
	"github.com/nmarsollier/cataloggo/internal/security"
	"github.com/nmarsollier/cataloggo/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// Singletons
var database *mongo.Database
var httpClient httpx.HTTPClient
var catalogCollection db.Collection
var articleConsumer consume.ArticleExistConsumer
var logoutConsumer consume.LogoutConsumer
var orderPlacedConsumer consume.OrderPlacedConsumer
var catalogService services.CatalogService
var emiter emit.RabbitEmitter

type Injector interface {
	Logger() log.LogRusEntry
	Database() *mongo.Database
	HttpClient() httpx.HTTPClient
	SecurityRepository() security.SecurityRepository
	SecurityService() security.SecurityService
	CatalogCollection() db.Collection
	ArticleRepository() article.ArticleRepository
	ArticleService() article.ArticleService
	ArticleExistConsumer() consume.ArticleExistConsumer
	LogoutConsumer() consume.LogoutConsumer
	OrderPlacedConsumer() consume.OrderPlacedConsumer
	CatalogService() services.CatalogService
	RabbitEmit() emit.RabbitEmitter
}

type Deps struct {
	CurrLog               log.LogRusEntry
	CurrHttpClient        httpx.HTTPClient
	CurrDatabase          *mongo.Database
	CurrSecRepo           security.SecurityRepository
	CurrSecSvc            security.SecurityService
	CurrCatalogColl       db.Collection
	CurrArticleRepository article.ArticleRepository
	CurrArticleService    article.ArticleService
	CurrArticleConsumer   consume.ArticleExistConsumer
	CurrLogoutConsumer    consume.LogoutConsumer
	CurrOrderPlaced       consume.OrderPlacedConsumer
	CurrCatalogServices   services.CatalogService
	CurrEmit              emit.RabbitEmitter
}

func NewInjector(log log.LogRusEntry) Injector {
	return &Deps{
		CurrLog: log,
	}
}

func (i *Deps) Logger() log.LogRusEntry {
	return i.CurrLog
}

func (i *Deps) Database() *mongo.Database {
	if i.CurrDatabase != nil {
		return i.CurrDatabase
	}

	if database != nil {
		return database
	}

	database, err := db.NewDatabase(env.Get().MongoURL, "catalog")
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}

	return database
}

func (i *Deps) HttpClient() httpx.HTTPClient {
	if i.CurrHttpClient != nil {
		return i.CurrHttpClient
	}

	if httpClient != nil {
		return httpClient
	}

	httpClient = httpx.Get()
	return httpClient
}

func (i *Deps) SecurityRepository() security.SecurityRepository {
	if i.CurrSecRepo != nil {
		return i.CurrSecRepo
	}
	i.CurrSecRepo = security.NewSecurityRepository(i.Logger(), i.HttpClient())
	return i.CurrSecRepo
}

func (i *Deps) SecurityService() security.SecurityService {
	if i.CurrSecSvc != nil {
		return i.CurrSecSvc
	}
	i.CurrSecSvc = security.NewSecurityService(i.Logger(), i.SecurityRepository())
	return i.CurrSecSvc
}

func (i *Deps) CatalogCollection() db.Collection {
	if i.CurrCatalogColl != nil {
		return i.CurrCatalogColl
	}

	if catalogCollection != nil {
		return catalogCollection
	}

	userCollection, err := db.NewCollection(i.CurrLog, i.Database(), "catalog")
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}

	return userCollection
}

func (i *Deps) ArticleRepository() article.ArticleRepository {
	if i.CurrArticleRepository != nil {
		return i.CurrArticleRepository
	}
	i.CurrArticleRepository = article.NewArticleRepository(i.Logger(), i.CatalogCollection())
	return i.CurrArticleRepository
}

func (i *Deps) ArticleService() article.ArticleService {
	if i.CurrArticleService != nil {
		return i.CurrArticleService
	}
	i.CurrArticleService = article.NewArticleService(i.Logger(), i.ArticleRepository())
	return i.CurrArticleService
}

func (i *Deps) ArticleExistConsumer() consume.ArticleExistConsumer {
	if i.CurrArticleConsumer != nil {
		return i.CurrArticleConsumer
	}
	if articleConsumer != nil {
		return articleConsumer
	}
	articleConsumer = consume.NewArticleExistConsumer(env.Get().FluentUrl, env.Get().RabbitURL, i.CatalogService())
	return articleConsumer
}

func (i *Deps) LogoutConsumer() consume.LogoutConsumer {
	if i.CurrLogoutConsumer != nil {
		return i.CurrLogoutConsumer
	}
	if logoutConsumer != nil {
		return logoutConsumer
	}
	logoutConsumer = consume.NewLogoutConsumer(env.Get().FluentUrl, env.Get().RabbitURL, i.SecurityService())
	return logoutConsumer
}

func (i *Deps) OrderPlacedConsumer() consume.OrderPlacedConsumer {
	if i.CurrOrderPlaced != nil {
		return i.CurrOrderPlaced
	}
	if orderPlacedConsumer != nil {
		return orderPlacedConsumer
	}
	orderPlacedConsumer = consume.NewOrderPlacedConsumer(env.Get().FluentUrl, env.Get().RabbitURL, i.CatalogService())
	return orderPlacedConsumer
}

func (i *Deps) CatalogService() services.CatalogService {
	if i.CurrCatalogServices != nil {
		return i.CurrCatalogServices
	}
	if catalogService != nil {
		return catalogService
	}
	catalogService = services.NewCatalogService(i.ArticleService(), i.RabbitEmit())
	return catalogService
}

func (i *Deps) RabbitEmit() emit.RabbitEmitter {
	if i.CurrEmit != nil {
		return i.CurrEmit
	}
	if emiter != nil {
		return emiter
	}
	emiter = emit.NewRabbitEmitter(env.Get().RabbitURL)
	return emiter
}
