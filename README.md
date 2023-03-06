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

# channel
# buffered channel
# mutex
# waitgroup
# atomic
# ticker

