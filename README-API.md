# CatalogGo
Microservicio de Catalogo.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /rabbit/article-data

#### GET
##### Summary

Mensage Rabbit article-data o article-exist

##### Description

Otros microservicios nos solicitan validar articulos en el catalogo, respondemos encviando direct al Exchange/Queue proporcionado.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| article-data | body | Message para Type = article-data | Yes | [service.ConsumeArticleValidation](#serviceconsumearticlevalidation) |

##### Responses

| Code | Description |
| ---- | ----------- |

#### PUT
##### Summary

Mensage Rabbit

##### Description

Emite respuestas de article-data or article-exist

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [service.EmitArticleValidation](#serviceemitarticlevalidation) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/logout

#### GET
##### Summary

Mensage Rabbit

##### Description

Escucha de mensajes logout desde auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [r_consume.logoutMessage](#r_consumelogoutmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/order-placed

#### GET
##### Summary

Mensage Rabbit order/order-placed

##### Description

Cuando se recibe el mensage order-placed damos de baja al stock para reservar los articulos. Queda pendiente enviar mensaje confirmando la operacion al MS de Orders.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| article-data | body | Message para Type = article-data | Yes | [service.ConsumeOrderPlacedMessage](#serviceconsumeorderplacedmessage) |

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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

---
### Models

#### apperr.ValidationErr

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [apperr.errField](#apperrerrfield) ] |  | No |

#### apperr.errField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

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

#### engine.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |

#### r_consume.logoutMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string | *Example:* `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"` | No |
| type | string | *Example:* `"logout"` | No |

#### service.ConsumeArticleValidation

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string | *Example:* `"Remote Exchange to Reply"` | No |
| message | [service.ConsumeArticleValidationMessage](#serviceconsumearticlevalidationmessage) |  | No |
| queue | string | *Example:* `"Remote Queue to Reply"` | No |
| type | string | *Example:* `"article-data"` | No |

#### service.ConsumeArticleValidationMessage

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
| cartId | integer |  | No |
| orderId | string |  | No |

#### service.EmitArticleValidation

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string | *Example:* `"ArticleId"` | No |
| price | number |  | No |
| referenceId | string | *Example:* `"Remote Reference Id"` | No |
| stock | integer |  | No |
| valid | boolean |  | No |
