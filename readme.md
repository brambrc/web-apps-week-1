# Web Application Mentoring Week-1

## Challenging Individuals Quote App with Database

- Concepts:
  - Same as Quote App
  - Database CRUD
    - Create database
    - Create Table
    - Insert Data into Table
    - Select Data from Table
- Input:
  - Same as Quote App
- Output:
  - Quote inserted to database
  - Select random qoutes from the databases
- Endpoint Docs :
    - `https://premium.zenquotes.io/zenquotes-documentation/`
    - `https://animechan.vercel.app/docs`
- Directions:
   - Pull branch challange ini dan buat branch baru dengan nama kalian pada repository ini.
        - clone spesific branch 

            `git clone -b challange https://github.com/brambrc/web-apps-week-1`
        - creating new branch
        
            `git checkout -b nama_kaliaannn`

            WARNING : JANGAN OTAK ATIK KODINGAN KALAU BELUM BUAT BRANCH BARU


  - Buat go mod init dan inisiasi / install library yang kalian butuhkan
  - Buat 1 table untuk menampung data dari qoutes yang akan di fetch dari api
  - Pilih salah satu endpoint atau source qoutes yang ingin dipakai, sesuaikan struktur kolom tablenya dengan response dari api tersebut
  - Pilih api yg bisa mendapatkan banyak qoutes dalam 1x hit
  - Boleh menggunakan RAW Query, ataupun ORM seperti GORM
  - Buat file migration tablenya menggunakan golang, dan file koneksi ke databasenya
  - Buat 4 endpoint, yang terdiri dari :
    - Endpoint fetch : pada endpoint ini kalian silahkan lakukan http request ke api qoutes yang kalian pilih, dan lakukan query INSERT untuk memasukan data qoutes yg didapatkan dari response, kedalam database yang kita buat. 
    - Endpoint get/select : pada endpoint ini kalian silahkan tampilkan qoutes random yang sudah tersimpan di database
    - Endpoint count : pada endpoint ini kalian silahkan tampilkan jumlah qoutes yang tersimpan pada database
    - Endpoint post / add : pada endpoint ini kalian silahkan bisa menambahkan endpoint sendiri kedalam database.

- Bila Sudah:

    - Silahkan push branch yang kalian buat, dan buatlah pull request, pilih target branch nya ke branch challange. Kemudian assign mentor, dan asisten mentor kalian sebagai reviewer.
        
        WARNING: JANGAN DI MERGE !!!!

    - Kabarin via channel discord BE02 kalau kalian sudah mengumpulkan
    - Deadline 8 Desember 2022.
    - Gudluk !

![hadeh](https://i.insider.com/5abb9e6a3216741c008b462d?width=600)