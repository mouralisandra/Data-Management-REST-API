# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete datas. No database implementation yet

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
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
### Get Single Book
``` bash
GET api/datas/{id}
```

### Delete Book
``` bash
DELETE api/datas/{id}
```

### Create Book
``` bash
POST api/datas

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Book Three",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }
```

### Update Book
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

## App Info

### Author

Brad Traversy
[Traversy Media](http://www.traversymedia.com)

### Version

1.0.0

### License

This project is licensed under the MIT License
