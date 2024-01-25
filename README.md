# -Golang Route App

Go'da bir rota çizim uygulaması planlayalım.
![image](https://user-images.githubusercontent.com/40759486/187882367-5535fa4d-74b6-41d2-aada-b3876b698b67.png)
Proje Gin Freamwork'ü kullanılarak, MongoDB'deki DB içerisindeki verileri kontrol edecek şekilde yazılmış olan bir webService uygulamasıdır.



-> Proje isterleri
•	Konum Ekleme Uç noktası: Gönderilen bilgiler içinde konumun enlem ve boylam bilgileri, konumun adı ve hexadecimal şekilde marker rengini veritabanına ekleyebilmelidir.
•	Konumları Listeleme Uç Noktası:  Kaydedilmiş konumları listeleyen uç noktadır.  Listede konumun adı, kullanıcının seçtiği renk ve  konum bilgileri olacaktır.
•	Konum Detayı Uç Noktası: Kaydedilmiş konumu tek olarak gösteren uç noktadır.
•	Konum Düzenleme Sayfası: Kaydedilmiş konumların düzenlendiği uç noktadır.
•	Rotalama Uç Noktası:  Konumlar gönderilen konuma göre en yakın noktadan başlayarak rotalanacaktır. (Kuş bakışı olarak uzaklık ölçülebilir. Bina ve yollar yoksayılabilir.)


Proje WebService katmanında Validation katmanına sahiptir.

Servisler eş katmana sahip düzeydedirler.

![mainDiagram](https://github.com/aliustunelin/golang-route-app/assets/40759486/751691a4-06d6-47e2-8701-d0577d696418)



Then then then; we are looking the code...


