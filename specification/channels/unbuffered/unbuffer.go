package main
import "fmt"

func send (i int, ch chan<- *T) {
	t := &T{i:byte(i)}
	ch <- t
	t.b = true //UNSAFE AT ANY SPEED
	//update after send could make that 
	//the change don't make efect

}

func main ()
	vs := make([]T,5)
	ch := make(chan *T) //then pruf to change the channel to buffer
	//you well se that when you read the *i
	//buffer send all the data and you can read the true value

	for i := range vs {
		go send(i,ch)
	}

	time.Sleep(1 * time.Second) //all goroutines started

	// copy quickly!
	for i := range vs {
		vs[i] = *<-ch
	}

}
