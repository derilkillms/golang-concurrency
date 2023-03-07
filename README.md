# golang-concurrency & parallel-programming
- Parallel programming sederhananya adalah memecahkan suatu masalah dengan cara membaginya menjadi yang lebih kecil, dan dijalankan secara bersamaan pada waktu yang bersamaan pula
- Concurrency adalah menjalankan beberapa pekerjaan secara bergantian

![image](https://user-images.githubusercontent.com/49135753/223047402-475d8358-802c-4b2d-ba6f-2db867e4945b.png)
![image](https://user-images.githubusercontent.com/49135753/223047500-6b52db64-72cd-4f8c-be99-dae68748f6e0.png)


# goroutines
- Goroutine adalah sebuah thread ringan yang dikelola oleh Go Runtime
- Ukuran Goroutine sangat kecil, sekitar 2kb, jauh lebih kecil dibandingkan Thread yang bisa sampai 1mb / 1000kb
- Namun tidak seperti thread yang berhalan parallel, goroutine berjalan secara concurrent
- Sebenarnya, goroutine dijalankan oleh Go Scheduler dalam thread, dimana jumlah thread nya sebanyak GOMAXPROCS (biasanya sejumlah core CPU)
- Goroutin bukan thread karena goroutine sendiri berjalan diatas thread
- - G:Goroutine
- - M:Thread (Machine)
- - P:Processor

![image](https://user-images.githubusercontent.com/49135753/223052561-d8a6426e-26ea-43a9-a7f5-3cb8fd8ae850.png)

# channel
- Secara default chanel hanya bisa menampung satu data, jika ingin menambah lagi harus menunggu data yang ada di channel diambil.
- Chanel hanya bisa menerima satu jenis data.
- Chanel bisa diambil dari lebih dari satu goroutine.
- Chanel harus di close jika tidak digunakan (memory leak).
- Channel bisa digunakan untuk mengirim dan menerima data
- Untuk mengirim data : channel <- data
- Untuk menerima data : data <- channel
- Jika selesai jangan lupa tutup chanel dengan close()

![image](https://user-images.githubusercontent.com/49135753/223309495-f7f4a7fb-2d6c-49b3-a565-7bd3b7e015b8.png)

# buffered channel
- Buffered channel digunakan untuk menampung data di dalam channel.
- Proses transfer data pada channel secara default dilakukan dengan cara un-buffered, atau tidak di-buffer di memori.
- Ketika terjadi proses kirim data via channel dari sebuah goroutine, maka harus ada goroutine lain yang menerima data dari channel yang sama, dengan proses serah-terima yang bersifat blocking.
- Selama jumlah data yang dikirim tidak melebihi jumlah buffer, maka pengiriman akan berjalan asynchronous (tidak blocking).
- Untuk membuat buffer kita bisa menggunakan make seperti sebelumnya hanya saja akan sedikit berbeda seperti ini make(chan string, JumLahBuffer)

# range channel
- Terkadang terdapat kasus dimana channel dikirim data secara terus menerus oleh pengirim nya, dan kadang tidak jelas kapan pengirim tersebut berhenti menerima data.
- Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel, ketika sebuah channel di close maka secara otomatis perulangan tersebut akan berhenti
- Ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual.

# select channel
- ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine untuk tiap channel nya, lalu kita ingin mendapatkan data dari channel tersebut, nah jika kita menggunakan perulangan for range maka akan ribet karena itu hanya bisa digunakan untuk satu channel.
- Untuk melakukan hal tersebut kita bisa menggunakan select channel yang ada di golang, dengan menggunakan select channel kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan maka akan diambil terlebih dahulu secara random salah satunya, lalu saat select selanjutnya barulah akan diambil data dari channel lainnya.

# mutex
- saat kita menggunakan goroutine dia tidak hanya berjalan secara concurrent tetapi bisa pararell juga, karena bisa pararell maka akan ada beberapa thread yang berjalan secara pararell
- ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama atau istilah nya sharing variable, jadi ada satu variable yang diakses oleh beberapa goroutine di waktu yang sama.
- ini bisa menyebabkan masalah yang namanya race condition
- untuk mengatasi masalah race condition yaitu dengan Mutex
- sebelum kita merubah variable yang di sharing ke beberapa goroutine, kita melakukan locking terhadap mutex, artinya setiap gorutine yang ingin mengakses harus melakukan locking lagi, dan saat kita melakukan locking maka mutex hanya mengizinkan satu goroutine untuk melakukan locking. Setelah kita melakukan locking dan selesai merubah variable nya maka kita melakukan unlock lagi, barulah goroutine selanjutnya diperbolehkan melakukan lock lagi.

# waitgroup
- sebuah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
- diperlukan saat kita ingin menjalankan beberapa proses goroutine, tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai.
- Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setelah proses goroutine selesai kita bisa menggunaklan method Done(), dan untuk menunggu semua proses selesai kita bisa menggunakan method Wait()

# atomic
-  Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proseas concurrent tanpa khawatir terkena race condition.

# ticker

- Ticker adalah representasi kejadian yang berulang, ketika waktu ticker expired maka event akan dikirim ke channel. Untuk membuat ticker kita bisa menggunakan timeNewTicker(duration). Dan untuk menghentikannya kita bisa menggunakan TickerStop().