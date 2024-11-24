# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete datas. No database implementation yet

## Quick Start


``` bash
docker-compose up --build
```

## Endpoints

### Get All Datas
``` bash
GET localhost:8000/datas
```
### Get Single Data
``` bash
GET localhost:8000/datas/{id}
```

### Delete Data
``` bash
DELETE localhost:8000/datas/{id}
```

### Create Data
``` bash
POST localhost:8000/datas
```
```json
{
  "isbn":"4545454",
  "title":"Data Three",
  "author":{"firstname":"Harry",  "lastname":"White"}
}
```

### Update Data
``` bash
PUT localhost:8000/datas/{id}

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Updated Title",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }

```


```
