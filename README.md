# go-optipng
Golang wrapper of optipng / Advanced PNG Optimizer
`optipng` command **is required**.

<http://optipng.sourceforge.net/>

## Usage

To install, use `go get`:

```
$ go get github.com/AKovalevich/go-optipng
```

### import

```go
import (
    optipng "github.com/AKovalevich/go-optipng"
)
```

#### func Compress

`func Compress(i image.Image, args []string) image.Image, error`

#### func CompressBytes

`func CompressBytes(b []byte, args []string) []bytes, error`

## License

MIT

## Author

Kovalevich Artem <https://github.com/AKovalevich>