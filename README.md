 # 📝 TO-DO App REST API (GoLang - Clean Architecture)  

 Bu proje, Go programlama dili ile geliştirilmiş bir TO-DO uygulaması REST API’sidir.   
 Clean Architecture yapısına uygun olarak organize edilmiştir.   
 Kullanıcılar yapılacaklar listesi oluşturabilir ve bu listelere adımlar (item) ekleyebilir.  
 Her bir adım tamamlandıkça, listenin tamamlama oranı hesaplanabilir.  

 ## 📌 Özellikler  
 
- ✅ Kullanıcılar için kimlik doğrulama (JWT tabanlı)  
- ✅ Yapılacaklar listesi oluşturma, güncelleme, silme  
- ✅ Listeye adım (item) ekleme, tamamlama ve silme  
- ✅ Her liste için tamamlama oranı hesaplama  
- ✅ 2 tip ön tanımlı kullanıcı: `admin`, `user`

  
## 🧑‍💻 Ön Tanımlı Kullanıcılar

| Kullanıcı Adı | Şifre   | Rol    |
|---------------|---------|--------|
| admin         | admin   | Admin  |
| user          | user    | User   |
| guest         | guest   | User   |

> 🔐 Admin: Tüm verilere erişebilir.    
> 👤 User: Sadece kendi verileriyle işlem yapabilir.

## 📁 Proje Yapısı  
.
├── cmd  
│ └── main.go # Uygulama başlangıç noktası  
├── internal  
│ ├── config # Uygulama konfigürasyonları  
│ ├── delivery/http # HTTP handler'lar ve middleware  
│ ├── domain # Entity ve interface tanımları  
│ ├── repository/mock # Mock repository'ler (veri katmanı)  
│ ├── usecase # İş mantığı katmanı  
│ └── utils # Yardımcı araçlar (JWT, vb.)  
├── go.mod  
└── go.sum  

## 🚀 Kurulum

### Gereksinimler

- Go 1.20+
- Git

### Kurulum Adımları

```bash
git clone https://github.com/eceakin/todo-app-project.git
cd todo-app-project
go mod tidy
go run cmd/main.go

```
Kurulumu yapıp main.go dosyasını çalıştırdıktan sonra Postman gibi bir uygulamayı açın.   
"POST" metodunu seçip url kısmına "http://localhost:8080/login" yazın.  
Sonrasında body kısmında "raw json" seçip "username" : "admin", "password" : "admin" yazın
Başarılı giriş sonrasında output kısmında size bir token dönecektir. 
Headers bölümünde "Key" kısmına "Authorization" ,"Value" kısmına ise Bearer boşluk aldığınız tokeni yazın.  

| Yöntem | URL                             | Açıklama                     |
| ------ | ------------------------------- | ---------------------------- |
| POST   | /api/lists                      | Yeni liste oluşturur         |
| PUT    | /api/lists/{id}                 | Listeyi günceller            |
| DELETE | /api/lists/{id}                 | Listeyi siler                |
| GET    | /api/lists                      | Tüm listeleri getirir        |
| GET    | /api/lists/{id}/items           | Listeye ait adımları getirir |
| GET    | /api/lists/{id}/completion-rate | Liste tamamlama oranı        |

| Yöntem | URL             | Açıklama                          |
| ------ | --------------- | --------------------------------- |
| POST   | /api/items      | Yeni adım ekler                   |
| PUT    | /api/items/{id} | Adımı günceller                   |
| DELETE | /api/items/{id} | Adımı siler                       |
| PATCH  | /api/items/{id} | Adımı tamamlandı olarak işaretler |
| GET    | /api/items/{id} | Belirli adımı getirir             |



### Liste ve Madde İşlemleri 
Başarılı şekilde giriş yaptıktan sonra   
  Liste eklemek için   
    url kısmına http://localhost:8080/api/lists yazın  
    "POST" işlemini seçin
    body kısmına ise "name" : "liste adı" yazın
    sonrasında "send" butonuna basın. 
    eğer her şey doğru ise 201 CREATED dönecektir. 

  Liste güncellemek için 
    url kısmına http://localhost:8080/api/lists/{güncellemek istediğiniz list_id} yazın  
    "PUT" işlemini seçin
    body ksımına ise "name" : "yeni liste adı" yazın 
    sonrasında "send" butonuna basın.  

  Liste silmek için 
    url kısmına http://localhost:8080/api/lists/{silmek istediğiniz list_id} yazın  
    "DELETE" işlemini seçin
    sonrasında "send" butonuna basın.  

  Liste tamamlanma oranı için 
    url kısmına http://localhost:8080/api/lists/id/completion-rate yazın  
    "GET" işlemini seçin
    sonrasında "send" butonuna basın. 

  Madde eklemek için 
     url kısmına http://localhost:8080/api/items yazın  
    "POST" işlemini seçin
    body kısmına ise "list_id" : "liste idsi"  , "content" : "madde içeriği" yazın
    sonrasında "send" butonuna basın. 

  Maddeyi tamamlandı olarak göstermek için 
    url kısmına http://localhost:8080/api/items/id yazın  
    "PATCH" işlemini seçin
    sonrasında "send" butonuna basın. 

## ⚠️ Hata Yönetimi

API, çeşitli senaryolarda anlamlı hata yanıtları döndürmeyi hedefler. İşte sık karşılaşılabilecek bazı hata durumları ve beklenen HTTP durum kodları:

| Durum Kodu | Açıklama                               | Örnek Hata Mesajı (JSON)                      | İlgili Kod (Örnek)                                                                 |
|------------|----------------------------------------|---------------------------------------------|------------------------------------------------------------------------------------|
| 400 Bad Request | İstemci tarafından gönderilen istek hatalı veya geçersiz. | `{"error": "invalid request"}`             | `http.Error(w, "invalid request", http.StatusBadRequest)` (Örn: `AuthHandler.Login`, `TodoHandler.CreateList`) |
| 401 Unauthorized | Kimlik doğrulama başarısız. Geçersiz kimlik bilgileri veya token. | `{"error": "invalid credentials"}`         | `http.Error(w, "invalid credentials", http.StatusUnauthorized)` (`AuthHandler.Login`)   |
| 403 Forbidden  | İstemcinin kaynağa erişim izni yok.   | `{"error": "not authorized"}`              | `errors.New("not authorized")` (`TodoItemUseCase.Create`, `GetByListID`, `GetByID`) |
| 404 Not Found  | İstenen kaynak bulunamadı.            | (Genellikle Go'nun default 404'ü veya özel bir mesaj) | (Kodda direkt 404 dönüşü görülmemekle birlikte, repository katmanından gelebilir) |
| 409 Conflict   | İstek, sunucudaki mevcut durumla çakışıyor. (Örn: Aynı isimde bir liste oluşturulmaya çalışılması) | (Bu senaryo kodda doğrudan ele alınmıyor)        |                                                                                    |
| 500 Internal Server Error | Sunucuda beklenmedik bir hata oluştu. | `{"error": "failed to ...", ...}`        | `http.Error(w, "failed to create list", http.StatusInternalServerError)` (Örn: `TodoHandler.CreateList`) |

**Örnek Hata Senaryoları ve Kod İlişkisi:**

* **Geçersiz İstek Verisi (`400 Bad Request`):**
    * `/login` endpoint'ine eksik veya hatalı JSON gönderildiğinde, `json.NewDecoder(r.Body).Decode(&req)` hata döndürür ve `http.Error(w, "invalid request", http.StatusBadRequest)` ile yanıt verilir.
    * Benzer şekilde, `/api/lists` veya `/api/items` gibi endpointlere geçersiz formatta veya eksik alanlarla istek yapıldığında da bu hata kodu döner.

* **Kimlik Doğrulama Başarısız (`401 Unauthorized`):**
    * `/login` endpoint'ine yanlış kullanıcı adı veya şifre gönderildiğinde, `h.authUseCase.Login` hata döndürür ve `http.Error(w, "invalid credentials", http.StatusUnauthorized)` ile yanıt verilir.

* **Yetkisiz Erişim (`403 Forbidden`):**
    * `TodoItemUseCase` içindeki `Create`, `Update`, `SoftDelete`, `GetByListID`, `CompleteItem` ve `GetByID` gibi fonksiyonlarda, kullanıcının ilgili kaynağa (liste veya madde) erişim yetkisi kontrol edilir. Yetkisiz bir durumda `errors.New("not authorized")` döndürülür ve bu hata handler katmanında `http.Error(w, "userID not found in context", http.StatusUnauthorized)` veya `http.Error(w, "failed to ...", http.StatusInternalServerError)` gibi yanıtlarla sonuçlanabilir (context'ten `userID` alınamaması durumu da `401`'e yol açabilir). **Önemli Not:** Yetkisiz erişim durumlarında `http.StatusForbidden` (403) dönmek daha语义 olarak doğru olabilir.

* **Kaynak Bulunamadı (`404 Not Found`):**
    * Şu anki kodda doğrudan `http.StatusNotFound` dönüşü olmamakla birlikte, usecase veya repository katmanlarında bir kaynağın (liste, madde, kullanıcı) ID'sine göre bulunamaması durumunda hata dönebilir. Bu hatalar handler katmanında yakalanıp `http.Error(w, "failed to get item", http.StatusInternalServerError)` gibi genel sunucu hatası olarak döndürülüyor. İlerleyen aşamalarda, bu gibi durumlarda `http.StatusNotFound` dönmek daha iyi bir uygulama olabilir.

* **Sunucu Hatası (`500 Internal Server Error`):**
    * Usecase katmanındaki fonksiyonlar (örneğin, liste veya madde oluşturma, güncelleme, silme işlemleri sırasında) beklenmedik bir hata ile karşılaştığında (örneğin, veritabanı hatası), `http.Error(w, "failed to create list", http.StatusInternalServerError)` gibi genel bir sunucu hatası döndürülür.

**Hata Yanıt Yapısı:**

Hata durumlarında API genellikle aşağıdaki JSON yapısında bir yanıt döner:

```json
{
  "error": "Hata mesajı."
}

🔧 Middleware
JWT doğrulama için AuthMiddleware kullanılır. /login dışındaki tüm uç noktalar bu middleware tarafından korunur.

## 🔗 Proje Bağlantısı

[GitHub Repo](https://github.com/eceakin/todo-app-project)
  
 
