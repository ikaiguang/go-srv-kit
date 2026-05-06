package jsonpkg

import (
	"sync"
	"testing"
)

// TestMarshalWithoutEscapeHTML_ConcurrentSafety 验证并发调用不会导致数据竞争
// go test -v -count 1 ./json -run TestMarshalWithoutEscapeHTML_ConcurrentSafety
func TestMarshalWithoutEscapeHTML_ConcurrentSafety(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	results := make([][]byte, goroutines)
	errs := make([]error, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			data := map[string]int{"index": idx}
			results[idx], errs[idx] = MarshalWithoutEscapeHTML(data)
		}(i)
	}
	wg.Wait()

	for i := 0; i < goroutines; i++ {
		if errs[i] != nil {
			t.Errorf("goroutine %d: unexpected error: %v", i, errs[i])
			continue
		}
		if len(results[i]) == 0 {
			t.Errorf("goroutine %d: result is empty", i)
		}
	}
}

// TestMarshalIndentWithoutEscapeHTML_ConcurrentSafety 验证并发调用不会导致数据竞争
// go test -v -count 1 ./json -run TestMarshalIndentWithoutEscapeHTML_ConcurrentSafety
func TestMarshalIndentWithoutEscapeHTML_ConcurrentSafety(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	results := make([][]byte, goroutines)
	errs := make([]error, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			data := map[string]int{"index": idx}
			results[idx], errs[idx] = MarshalIndentWithoutEscapeHTML(data, "", "  ")
		}(i)
	}
	wg.Wait()

	for i := 0; i < goroutines; i++ {
		if errs[i] != nil {
			t.Errorf("goroutine %d: unexpected error: %v", i, errs[i])
			continue
		}
		if len(results[i]) == 0 {
			t.Errorf("goroutine %d: result is empty", i)
		}
	}
}

// TestMarshalWithoutEscapeHTML_NoHTMLEscape 验证 HTML 字符不被转义
// go test -v -count 1 ./json -run TestMarshalWithoutEscapeHTML_NoHTMLEscape
func TestMarshalWithoutEscapeHTML_NoHTMLEscape(t *testing.T) {
	data := map[string]string{"url": "https://example.com?a=1&b=2"}
	result, err := MarshalWithoutEscapeHTML(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := string(result)
	if expected := `{"url":"https://example.com?a=1&b=2"}` + "\n"; got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}
