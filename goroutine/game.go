package goroutine

import (
	"fmt"
	"net/http"
)

func Request() {

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, _ := http.Get("https://www.baidu.com")
			fmt.Println(res)
		}()
	}
	wg.Wait()
}
