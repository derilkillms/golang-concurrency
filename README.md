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
# mutex
# waitgroup
# atomic
# ticker

