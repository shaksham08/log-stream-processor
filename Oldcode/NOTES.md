# Important Notes

Some important notes while working on the project

- IF we take a look at the below `handler.go` file code

```go
package handler

import (
	"fmt"
	"sync"

	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func Init(ch chan models.Event, wg *sync.WaitGroup) {
	for event := range ch {
		fmt.Println("Event recieved : ", event)
	}
	wg.Done()
}
```

- we see we only have one go routine thats handling everything.
- If we have high throughput then we one go routine will not be able to handle that.
- One way if we create a buffer channel
- But also we can introduce more go routines

```go
package handler

import (
	"fmt"
	"sync"

	"github.com/shaksham08/log-stream-processor/config"
	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func processEvent(ch chan models.Event, wg *sync.WaitGroup, idx int) {
	for event := range ch {
		fmt.Println("Event recieved : ", event, "by handler", idx)
	}
	defer wg.Done()
}

func Init(ch chan models.Event, wg *sync.WaitGroup) {
	for i := 0; i < config.MAX_HANDLERS; i++ {
		wg.Add(1)
		go processEvent(ch, wg, i)
	}
}
```

- our `main.go`

```go
package main

import (
	"sync"
	"time"

	"github.com/shaksham08/log-stream-processor/pkg/handler"
	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func simulateIngress(ch chan models.Event, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- models.SystemLog{
			Log: models.Log{
				ID:     i,
				Source: "System",
				Body:   "System is running",
			},
			Severity: "Info",
		}
		time.Sleep(1 * time.Second)
	}
	defer wg.Done()
}

func main() {
	ch := make(chan models.Event)
	var wg sync.WaitGroup
	go handler.Init(ch, &wg)
	wg.Add(1)
	simulateIngress(ch, &wg)
	close(ch)
	defer wg.Wait()
}
```

- now there is a small issue lets say the go routine is active and somehow the main channel is closed but eventually the go routine will still be working but that is wrong. So we need to do some cleanup. For this we will be using Context in go

- We can read more on this link [Context in go ](https://pkg.go.dev/context)

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/shaksham08/log-stream-processor/pkg/models"
)


func sendToChannel(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i < 11; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go sendToChannel(ch, &wg)
	defer wg.Wait()
	for {
		select {
		case i, ok := <-ch:
			{
				fmt.Println("recieved from ch ", i)
			}
			if !ok {
				fmt.Println("Channel closed")
				return
			}
		}
	}

}

```

- Here the output is

```go
recieved from ch  1
recieved from ch  2
recieved from ch  3
recieved from ch  4
recieved from ch  5
recieved from ch  6
recieved from ch  7
recieved from ch  8
recieved from ch  9
recieved from ch  10
recieved from ch  0
Channel closed
```

- Here we can see we got all 10 values that we passed to channel but at the end we also go a 0 value that is because the channel was closed and the default value was set to 0

- So, the reason we're seeing 11 values (from 1 to 10 and then 0) is because after sending values from 1 to 10, the channel is closed. When we read from a closed channel in Go, it returns the zero value for the type of that channel, which in this case is an int, hence we're seeing the additional 0 value.

- Using a range loop simplifies your code, and it automatically stops when the channel is closed, avoiding the need for an explicit check on ok.
