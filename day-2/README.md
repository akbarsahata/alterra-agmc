# DB Connection, GORM, & MVC

## Tugas 1 - Static API

Pengerjaan tugas ini bisa dilihat di `routes/book.route.go`, `controllers/book.controller.go`, `modes/book.model.go` dan beberapa file penunjang di folder `lib`.

Port server dan `{{base_url}}` disesuaikan menjadi 3001 karena menggunakan [gin](https://github.com/codegangsta/gin). 

Hasil dari pemanggilan setiap endpoint bisa dilihat di examples pada postman collection di folder `v1/books`.


## Tugas 2 - Dynamic API

### GET /v1/users

get many users

### GET /v1/users/:id

get one specific user by its ID

### POST /v1/users

create one user

#### expected request payload


```
{
    "name": <string, required>,
    "email": <string, required, valid email format>,
    "password <string, required, min length 6 chars>
}
```

### PUT /v1/users/:id

update specific user by its ID

```
{
    "name": <string, required>,
    "email": <string, required, valid email format>,
}
```

### DELETE /v1/users/:id

delete (soft) specific user by its ID. will return 204 is successfully deleting record
