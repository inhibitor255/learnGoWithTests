package concurrency

type websiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWesites(wc websiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		u := url
		go func() {
			resultChannel <- result{u, wc(u)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
