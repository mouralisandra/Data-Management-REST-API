# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete datas. No database implementation yet

## Quick Start


``` bash
# Install mux router
go mod init projet
go mod tidy
```

``` bash
go build
./go_restapi
```

## Endpoints

### Get All Datas
``` bash
GET api/datas
```
### Get Single Data
``` bash
GET api/datas/{id}
```

### Delete Data
``` bash
DELETE api/datas/{id}
```

### Create Data
``` bash
POST api/datas

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Data Three",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }
```

### Update Data
``` bash
PUT api/datas/{id}

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Updated Title",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }

```


```
