# Take Home Test Haioo.id Answer

## 1. 
Program menghitung pecahan mata uang. Source Code: [Click here](https://github.com/yusnarsetiyadi/Take-Home-Test-Haioo.id/blob/main/moneyfraction/moneyFraction.go)

## 2. 
Program menghitung proses edit dari dua buah string. Source Code: Click here

## 3. 
script berikut adalah script dari dockerfile:
```
FROM golang
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
LISTEN 80
```
ada beberapa kesalahan yang saya temukan didalam script ini, diantaranya:
- kesalahan pertama dari script ini ada di line, seharusnya diberikan versi golang yang akan berjalan di container misal "FROM golang:1.17.7-alpine as build"
- kesalahan kedua ada di line 6, seharusnya tidak diperlukan karena source nya sudah digunakan dari line 2
- kesalahan ketiga ada di line 8, seharusnya tidak menggunakan "LISTEN" tapi menggunakan "EXPOSE"

## 4.
Tujuan Microservice:
- tujuan penggunaan microservice adalah agar sebuah project bisa lebih modular, karena desain arsitektur microservice terdiri dari berbagai unit layanan tersendiri namun masih terhubung, sehingga para developer dapat lebih cepat dalam mengembangkan fitur di masing-masing unit yang ditanganinya, resiko kegagalannya pun dapat lebih diminimalisir dan yang lebih penting lagi menurut saya adalah aplikasi nya dapat berjalan secara independen sehingga kita dapat memnuhi kebutuhan bisnis yang bisa saja berubah seiring berjalannya waktu.

## 5. 
Cara kerja Index pada database:
index adalah objek dalam database untuk mempercepat query. index dapat digunakan pada database berskala besar dengan cara membuat index pada kolom yang ingin di query, contoh: 
```
CREATE INDEX data_siswa_idx ON data_siswa(siswa_id) NOLOGGING COMPUTE STATISTICS;
```
setelah query index dibuat seperti contoh, selanjutnya kita dapat menjalankan query untuk mencari data siswa berdasarkan id dengan waktu yg lebih singkat. contoh query:
```
SELECT * FROM data_siswa WHERE id = 1
```

## 6
Backend Service Shopping Chart: [Click here](https://github.com/yusnarsetiyadi/Take-Home-Test-Haioo.id)

- clone repository
```
git clone https://github.com/yusnarsetiyadi/Take-Home-Test-Haioo.id.git
```
- change direktory
```
cd Take-Home-Test-Haioo.id
```