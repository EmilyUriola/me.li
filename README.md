#### Meli Test

Operaciones:

- Get.
- Add (Post).
- Add Bulk.
- Delete.

### Como Usar

##Agregar URL:

``` shell
 curl -X POST http://localhost:8011/urls -d "{\"long\": \"https://github.com\"}"
```
#Respuesta:
{"message": "URL shortified.", "short_url": "http://localhost:8011/2931944842"}

##Agregar Bulk URL:

``` shell
curl -X POST http://localhost:8011/urls/bulk -d "[{\"long\": \"http://www.amazon.com\"},{\"long\": \"http://www.gmail.com\"},{\"long\": \"http://as.com\"},{\"long\": \"https://circleci.com\"}]"
```
#Respuesta:
{"message": "URL shortified.", "short_url": "["http://localhost:8011/37044323","http://localhost:8011/1171990929","http://localhost:8011/1736520920","http://localhost:8011/1421392713"]"}

##Eliminar URL:
``` shell
 curl -X GET http://localhost:8011/urls/37044323
```
#Respuesta:
{"long_url": "http://www.amazon.com", "short_url": "http://localhost:8011/37044323"}

##Eliminar URL:
``` shell
 curl -X DELETE http://localhost:8011/urls/37044323
```
#Respuesta:
{"message": "Delete URL shortified.", "short_url": "http://localhost:8011/37044323"}

##Uso de URL: