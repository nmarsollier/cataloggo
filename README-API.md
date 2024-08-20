# CatalogGo
Microservicio de Catalogo.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /rabbit/article_exist

#### GET
##### Summary

Mensage Rabbit article_exist/article_exist

##### Description

Otros microservicios nos solicitan validar articulos en el catalogo.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| article_exist | body | Message para article_exist | Yes | [service.ConsumeArticleExist](#serviceconsumearticleexist) |

##### Responses

| Code | Description |
| ---- | ----------- |

#### PUT
##### Summary

Mensage Rabbit article_exist

##### Description

Emite respuestas de article_exist

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [service.SendArticleExist](#servicesendarticleexist) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/logout

#### GET
##### Summary

Mensage Rabbit logout

##### Description

Escucha de mensajes logout desde auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [consume.logoutMessage](#consumelogoutmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/order_placed

#### GET
##### Summary

Mensage Rabbit order_placed/order_placed

##### Description

Cuando se recibe el mensage order_placed damos de baja al stock para reservar los articulos. Queda pendiente enviar mensaje confirmando la operacion al MS de Orders.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| order_placed | body | Message order_placed | Yes | [service.ConsumeOrderPlacedMessage](#serviceconsumeorderplacedmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

---
### /v1/articles

#### POST
##### Summary

Crear Artículo

##### Description

Crear Artículo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |
| body | body | Informacion del articulo | Yes | [article.NewArticleData](#articlenewarticledata) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Articulo | [article.ArticleData](#articlearticledata) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

### /v1/articles/:articleId

#### DELETE
##### Summary

Eliminar Artículo

##### Description

Eliminar Artículo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |
| articleId | path | ID de articlo | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | No Content |  |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

#### GET
##### Summary

Obtener un articulo

##### Description

Obtener un articulo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |
| articleId | path | ID de articlo | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Articulo | [article.ArticleData](#articlearticledata) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

#### POST
##### Summary

Actualizar Artículo

##### Description

Actualizar Artículo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |
| body | body | Informacion del articulo | Yes | [article.NewArticleData](#articlenewarticledata) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Articulo | [article.ArticleData](#articlearticledata) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

### /v1/articles/search/:criteria

#### GET
##### Summary

Obtener un articulo

##### Description

Obtener un articulo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |
| articleId | path | ID de articlo | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Articulos | [ [article.ArticleData](#articlearticledata) ] |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

---
### Models

#### article.ArticleData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| _id | string |  | No |
| description | string |  | No |
| enabled | boolean |  | No |
| image | string |  | No |
| name | string |  | No |
| price | number |  | No |
| stock | integer |  | No |

#### article.NewArticleData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | string |  | Yes |
| image | string |  | No |
| name | string |  | Yes |
| price | number |  | Yes |
| stock | integer |  | Yes |

#### consume.logoutMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string | *Example:* `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"` | No |
| type | string | *Example:* `"logout"` | No |

#### engine.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |

#### errs.ValidationErr

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [errs.errField](#errserrfield) ] |  | No |

#### errs.errField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### service.ArticleExistMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string | *Example:* `"ArticleId"` | No |
| price | number |  | No |
| referenceId | string | *Example:* `"Remote Reference Id"` | No |
| stock | integer |  | No |
| valid | boolean |  | No |

#### service.ConsumeArticleExist

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string | *Example:* `"Remote Exchange to Reply"` | No |
| message | [service.ConsumeArticleExistMessage](#serviceconsumearticleexistmessage) |  | No |
| routing_key | string | *Example:* `"Remote RoutingKey to Reply"` | No |

#### service.ConsumeArticleExistMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string | *Example:* `"ArticleId"` | No |
| referenceId | string | *Example:* `"Remote Reference Object Id"` | No |

#### service.ConsumeOrderPlacedArticle

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| quantity | integer |  | No |

#### service.ConsumeOrderPlacedMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | [ [service.ConsumeOrderPlacedArticle](#serviceconsumeorderplacedarticle) ] |  | No |
| cartId | string |  | No |
| orderId | string |  | No |

#### service.SendArticleExist

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | [service.ArticleExistMessage](#servicearticleexistmessage) |  | No |
