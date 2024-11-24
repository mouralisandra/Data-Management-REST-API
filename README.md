# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete datas. No database implementation yet

## Quick Start


``` bash
docker-compose up --build
```

# Endpoints
## Customer Data :
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
  "title":"Data 3",
  "author":{"firstname":"Sandra",  "lastname":"Mourali"}
}
```

### Update Data
``` bash
PUT localhost:8000/datas/{id}
```
```json
{
  "isbn":"4545454",
  "title":"Updated Title",
  "author":{"firstname":"Sandra",  "lastname":"Mourali"}
}
```
## Documents :
echo "# Documents-Management-REST-API" >> README.md
## Endpoints

### Get All Documentss
``` bash
GET localhost:8000/documents
```
### Get Single Documents
``` bash
GET localhost:8000/documents/{id}
```

### Delete Documents
``` bash
DELETE localhost:8000/documents/{id}
```

### Create Documents
``` bash
POST localhost:8000/documents
```
```json
{
  "isbn":"4545454",
  "title":"Documents 3",
  "author":{"firstname":"Sandra",  "lastname":"Mourali"}
}
```

### Update Documents
``` bash
PUT localhost:8000/documents/{id}
```
```json
{
  "isbn":"4545454",
  "title":"Updated Title",
  "author":{"firstname":"Sandra",  "lastname":"Mourali"}
}
```