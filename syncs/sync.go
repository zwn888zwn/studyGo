package syncs

import "context"

func generateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}

		}
	}()
	return ch
}
func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}
