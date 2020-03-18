<div align="center">
  <h1>go-sitemap</h1>
  
  [![forthebadge](https://forthebadge.com/images/badges/check-it-out.svg)](https://forthebadge.com)

generates sitemap dynamically, for user-generated URLs
</div>

# install
```bash
go get github.com/xis/go-sitemap
```

# using
```go
import "github.com/xis/go-sitemap"

func main() {
	var s sitemap.Sitemap
	err := sitemap.Create("www.example.com/path/", []string{"userid", "contentid", "something", "somethingelse"}).Result(&s)
	if err != nil {
		panic(err)
	}
}
```
