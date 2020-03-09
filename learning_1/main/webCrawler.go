package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch mengembalikan isi dari URL dan daftar URL yang ditemukan
	// di halaman tersebut.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl menggunakan fetcher untuk secara rekursif mengambil semua halaman
// dimulai dari url, sampai kedalaman maksimum `depth`.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher adalah Fetcher yang mengembalikan hasil dari tampungan.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher adalah pengembangan dari fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// Latihan: Web Crawler

// Dalam latihan ini anda akan menggunakan fitur konkurensi Go untuk mem-paralelkan sebuah web crawler.

// Ubah fungsi Crawl untuk mengambil URL secara paralel tanpa ada duplikasi (mengambil URL yang sama dua kali).

// Petunjuk: anda bisa menyimpan cache dari URL yang telah diambil menggunakan sebuah map, tapi map tidak aman untuk digunakan secara konkuren!