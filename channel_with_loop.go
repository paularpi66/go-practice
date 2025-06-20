func main() {
	ch := make(chan int)

	go func() {
		for i := 1; <= 5; i++{
			ch <- I
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
}