package main

func main() {
	ch := make(chan int)
	// Эта горутина передает в канал числа от 0 до 9 и потом завершается
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// Этот цикл считает переданные в канал числа и заблокируется. Так как остается единственная
	// активная горутина и она заблокирована, наступит deadlock.
	for n := range ch {
		println(n)
	}
}