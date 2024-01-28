# -Golang Route App

Go'da bir rota çizim uygulaması planlayalım. </br> </br> </br>
![image](https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg)  </br>
</br>->Proje Fiber Freamwork'ü kullanılarak, MongoDB'deki DB içerisindeki verileri kontrol edecek şekilde yazılmış olan bir webService uygulamasıdır.</br></br>
->Uygulama içerisinde Golang'in Fiber web çatısıyla katmanlı mimari yazılmış olup, 6 adet endpoint bulunmaktadır. Rotalarla ilgili yeni lokasyon ekleme, mevcut lokasyonu düzenleme, lokasyon silme, 1 adet lokasyon ait bilgileri detaylı bir şekilde görme, tüm lokasyonları listeleme ve veritabanındaki tüm lokasyonlarla girilecek olan yeni bir merkez arasında rota çizme özellikleri uygulama içerisinde vardır.</br></br>
->Uygulama içerisinde uniq test birimleri vardır.</br></br>
->Uygulama istek limitine sahiptir. Max 20 tane istek hafızada 30sn olacak şekilde uygulama dışarıdan yapılacak olağan dışı trafiklere karşı koruma katmanına sahiptir.</br></br>
->Uygulama Dockerize bir şekilde çalışma imkanı sunmaktadır.</br></br>
->Uygulama CI/CD süreçlerine uygun bir şekilde geliştirilmiştir. Github Actions kullanılarak CI pipeline süreçleri yönetilmiş olup uygulama geliştrme ve test ortamında çalışmasında süreç main dalı üzerinde otomatize edilmiştir.</br></br>
->CI/CD süreçlerinden deploy kısmında uygulama Heroku üzerinde canlıya alınmıştır.</br>

</br></br></br>


</br>
-> Detaylı Proje isterleri</br>
•	Konum Ekleme Uç noktası: Gönderilen bilgiler içinde konumun enlem ve boylam bilgileri, konumun adı ve hexadecimal şekilde marker rengini veritabanına ekleyebilmelidir.</br>
•	Konumları Listeleme Uç Noktası:  Kaydedilmiş konumları listeleyen uç noktadır.  Listede konumun adı, kullanıcının seçtiği renk ve  konum bilgileri olacaktır.</br>
•	Konum Detayı Uç Noktası: Kaydedilmiş konumu tek olarak gösteren uç noktadır.</br>
•	Konum Düzenleme Sayfası: Kaydedilmiş konumların düzenlendiği uç noktadır.</br>
•	Rotalama Uç Noktası:  Konumlar gönderilen konuma göre en yakın noktadan başlayarak rotalanacaktır. (Kuş bakışı olarak uzaklık ölçülebilir. Bina ve yollar yoksayılabilir.)</br>

</br></br>
Proje WebService katmanında Validation katmanına sahiptir.</br></br>
![photo1706261801 (1)](https://github.com/aliustunelin/golang-route-app/assets/40759486/a90d7a6a-5acc-45a9-84ed-5fcf0d1f4583)
</br>


![mainDiagram](https://github.com/aliustunelin/golang-route-app/assets/40759486/751691a4-06d6-47e2-8701-d0577d696418)


</br></br>
** Uygulama çalışma adımları: <br>
Go Paketleri
```bash
go.mongodb.org/mongo-driver/mongo/options
go get github.com/joho/godotenv
go get github.com/stretchr/testify/assert
go install github.com/golang/mock/mockgen@v1.6.0
go get  github.com/golang/mock/mockgen/model
go get golang.org/x/sync/errgroup
go get go.mongodb.org/mongo-driver/x/mongo/driver/ocsp@v1.9.1
go get -u github.com/gofiber/fiber/v2
go get golang.org/x/time/rate
go get golang.org/x/net/http
```

Go Run veya Go Build
```bash
go mod tidy
go run main.go
go build
```

Docker İmage ve Run
```bash
docker build -t my-go-app:1.0 .
docker run -p 8080:8080 my-go-app:1.0
```



--->>>> Run Project
#depends eklenecek</br>
go run main.go

</br></br></br></br></br>
Then then then; we are looking the code...


