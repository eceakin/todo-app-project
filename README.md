![Go](https://img.shields.io/badge/Go-1.20%2B-blue?style=for-the-badge&logo=go)
![Status](https://img.shields.io/badge/Status-Completed-brightgreen?style=for-the-badge)
![License](https://img.shields.io/badge/License-Educational-lightgrey?style=for-the-badge)
![Architecture](https://img.shields.io/badge/Architecture-Clean%20Architecture-yellow?style=for-the-badge)
![Mock](https://img.shields.io/badge/Data-Mock%20Repository-orange?style=for-the-badge)



# 📝 TO-DO App REST API (GoLang - Clean Architecture)

Bu proje, Go programlama dili ile geliştirilmiş bir TO-DO uygulaması REST API’sidir. Clean Architecture yapısına uygun olarak organize edilmiştir. Kullanıcılar yapılacaklar listesi oluşturabilir ve bu listelere adımlar (item) ekleyebilir. Her bir adım tamamlandıkça, listenin tamamlama oranı hesaplanabilir.

---

## 🚀 Özellikler

* 🔐 JWT tabanlı kimlik doğrulama
* 📝 Yapılacaklar listesi oluşturma, güncelleme, silme
* 📌 Listeye adım (item) ekleme, tamamlama ve silme
* 📊 Her liste için tamamlama oranı hesaplama
* 👥 2 tip ön tanımlı kullanıcı: `admin`, `user`

---

## 👤 Ön Tanımlı Kullanıcılar

| Kullanıcı Adı | Şifre | Rol   |
| ------------- | ----- | ----- |
| admin         | admin | Admin |
| user          | user  | User  |
| guest         | guest | User  |

> 🔐 **Admin:** Tüm verilere erişebilir.  
> 👤 **User:** Sadece kendi verileriyle işlem yapabilir.

---

## 📁 Proje Yapısı

```
.
├── cmd                  # Uygulama giriş noktası (main.go)
├── internal
│   ├── config           # Konfigürasyonlar
│   ├── delivery/http    # HTTP handler'lar ve middleware
│   ├── domain           # Entity ve interface tanımları
│   ├── repository/mock  # Mock repository'ler
│   ├── usecase          # İş mantığı katmanı
│   └── utils            # Yardımcı araçlar (JWT, vb.)
├── go.mod
└── go.sum
```

---

## ⚙️ Kurulum

### Gereksinimler

* Go 1.20+
* Git

### Kurulum Adımları

```bash
git clone https://github.com/eceakin/todo-app-project.git
cd todo-app-project
go mod tidy
go run cmd/main.go
```

---

## 🧪 Postman ile Test Etme

### 1️⃣ Giriş Yapma (Token Alma)

* Yöntem: `POST`
* URL: `http://localhost:8080/login`
* Headers:

  * `Content-Type`: `application/json`
* Body (raw > JSON):

```json
{
  "username": "admin",
  "password": "admin"
}
```

* Cevap: Aşağıdaki gibi bir JWT token döner:

```json
{
  "token": "..."
}
```

### 🔐 Token Kullanımı

Tüm korumalı uç noktalar için aşağıdaki header'ı ekleyin:

* Key: `Authorization`
* Value: `Bearer <TOKEN>`

---

## 🔧 API Uç Noktaları

### 📋 Liste İşlemleri

| Metot  | URL                             | Açıklama                     |
| ------ | ------------------------------- | ---------------------------- |
| POST   | /api/lists                      | Yeni liste oluşturur         |
| PUT    | /api/lists/{id}                 | Listeyi günceller            |
| DELETE | /api/lists/{id}                 | Listeyi siler                |
| GET    | /api/lists                      | Tüm listeleri getirir        |
| GET    | /api/lists/{id}/items           | Listeye ait adımları getirir |
| GET    | /api/lists/{id}/completion-rate | Liste tamamlama oranı        |

### ✅ Adım (Item) İşlemleri

| Metot  | URL             | Açıklama                          |
| ------ | --------------- | --------------------------------- |
| POST   | /api/items      | Yeni adım ekler                   |
| PUT    | /api/items/{id} | Adımı günceller                   |
| DELETE | /api/items/{id} | Adımı siler                       |
| PATCH  | /api/items/{id} | Adımı tamamlandı olarak işaretler |
| GET    | /api/items/{id} | Belirli adımı getirir             |

---

## 📌 Kullanım Örnekleri

### 📄 Liste Ekleme

* URL: `http://localhost:8080/api/lists`
* Yöntem: `POST`
* Body:

```json
{
  "name": "Alışveriş Listem"
}
```

### 📝 Liste Güncelleme

* URL: `http://localhost:8080/api/lists/{list_id}`
* Yöntem: `PUT`
* Body:

```json
{
  "name": "Güncellenmiş Liste"
}
```

### ❌ Liste Silme

* URL: `http://localhost:8080/api/lists/{list_id}`
* Yöntem: `DELETE`

### 📈 Tamamlama Oranı

* URL: `http://localhost:8080/api/lists/{id}/completion-rate`
* Yöntem: `GET`

### ➕ Adım Ekleme

* URL: `http://localhost:8080/api/items`
* Yöntem: `POST`
* Body:

```json
{
  "list_id": 1,
  "content": "Süt al"
}
```

### ✅ Adımı Tamamlama

* URL: `http://localhost:8080/api/items/{id}`
* Yöntem: `PATCH`

---

## ⚠️ Hata Yönetimi

| Durum Kodu | Açıklama                              | Örnek Yanıt                                       |
| ---------- | ------------------------------------- | ------------------------------------------------- |
| 400        | Geçersiz istek                        | `{ "error": "invalid request" }`                  |
| 401        | Kimlik doğrulama başarısız            | `{ "error": "invalid credentials" }`              |
| 403        | Yetkisiz erişim                       | `{ "error": "not authorized" }`                   |
| 404        | Kaynak bulunamadı                     | `{ "error": "not found" }` (önerilen)             |
| 500        | Sunucu hatası                         | `{ "error": "failed to ..." }`                    |



---

## 🔒 Middleware

JWT doğrulaması için AuthMiddleware kullanılır. `/login` hariç tüm uç noktalar bu middleware ile korunmaktadır.

---

## 🌐 Proje Bağlantısı

🔗 [GitHub Repo](https://github.com/eceakin/todo-app-project)

---

