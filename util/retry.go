package util

import (
	"strings"
	"time"
)

// RetryOnBadGateway tries to call a method and checks if the method returns an error, in this case if the response was a Bad Gateway
func RetryOnBadGateway(f func() error) {

	var err error
	for i := 0; i < 3; i++ {
		err = f()
		if err != nil {
			if strings.HasPrefix(err.Error(), "HTTP 502") {
				time.Sleep(1 * time.Second)
				continue
			} else {
				PanicOnErr(err)
			}
		} else {
			return
		}
	}
}
