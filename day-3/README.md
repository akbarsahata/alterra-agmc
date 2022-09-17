# Middleware

Penerapan middleware pada framework Echo dan endpoint `/login`

Untuk mengakses endpoint yang di-_protect_, perlu menambahkan header `Authorization` dengan format `Bearer <valid token>`. String token diperoleh dari dari login yang berhasil. Detail login dibawah.

## POST /login

### Expexted Request Payload

```
{
    "email": <string, required, valid email>,
    "password": <string, required>
}
```

### Response Payload

```
{
    "token": <string>
}
```