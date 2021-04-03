# Tạo landing page siêu nhanh , nhẹ với [valyala/fasthttp](https://github.com/valyala/fasthttp)

```
go mod init techmaster.vn/fastlandingpage
```

Tạo mới file [fastlandingpage.go](fastlandingpage.go)
```
code fastlandingpage.go
```
Đây là nội dung file
```go
package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	fs := &fasthttp.FS{
		// Path to directory to serve.
		Root: "./",

		// Generate index pages if client requests directory contents.
		GenerateIndexPages: false,

		IndexNames: []string{"index.html", "index.htm"},

		// Enable transparent compression to save network traffic.
		Compress: true,
	}

	// Create request handler for serving static files.
	h := fs.NewRequestHandler()

	// Start the server.
	if err := fasthttp.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
```

Tiếp tục gõ lệnh ở termimal:
Biên dịch go module
```
go build
```

Chạy file binary sau biên dịch
```
./fastlandingpage
```

Mở trình duyệt vào http://localhost:8080 sẽ thấy web site đơn giản với ảnh con mèo