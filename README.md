# -Golang Route App

Go'da bir rota çizim uygulaması planlayalım. </br> </br> </br>
![image](https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg)  </br>
Proje Fiber Freamwork'ü kullanılarak, MongoDB'deki DB içerisindeki verileri kontrol edecek şekilde yazılmış olan bir webService uygulamasıdır. </br></br></br>


</br>
-> Proje isterleri</br>
•	Konum Ekleme Uç noktası: Gönderilen bilgiler içinde konumun enlem ve boylam bilgileri, konumun adı ve hexadecimal şekilde marker rengini veritabanına ekleyebilmelidir.</br>
•	Konumları Listeleme Uç Noktası:  Kaydedilmiş konumları listeleyen uç noktadır.  Listede konumun adı, kullanıcının seçtiği renk ve  konum bilgileri olacaktır.</br>
•	Konum Detayı Uç Noktası: Kaydedilmiş konumu tek olarak gösteren uç noktadır.</br>
•	Konum Düzenleme Sayfası: Kaydedilmiş konumların düzenlendiği uç noktadır.</br>
•	Rotalama Uç Noktası:  Konumlar gönderilen konuma göre en yakın noktadan başlayarak rotalanacaktır. (Kuş bakışı olarak uzaklık ölçülebilir. Bina ve yollar yoksayılabilir.)</br>

</br></br>
Proje WebService katmanında Validation katmanına sahiptir.</br>

Servisler eş katmana sahip düzeydedirler.</br>

![mainDiagram](https://github.com/aliustunelin/golang-route-app/assets/40759486/751691a4-06d6-47e2-8701-d0577d696418)


--->>>> Run Project
#depends eklenecek
go run main.go

</br></br></br></br></br>
Then then then; we are looking the code...


