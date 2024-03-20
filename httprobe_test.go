package httprobe

import (
	"fmt"
	"testing"
)

func TestHTTProbe(t *testing.T) {
	results := HTTProbe([]string{"zoom.com", "amazon.com", "youtube.com"})
	fmt.Println(results)
}
