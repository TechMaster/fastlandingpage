# Tạo landing page siêu nhanh , nhẹ với [valyala/fasthttp](https://github.com/valyala/fasthttp)

Hiện nay Techmaster sử dụng iris framework cũng code rất dài dòng và thừa chỉ để phục vụ langding page tính. Việc này khiến cho Docker image size vừa lớn, tốc độ khởi động chậm và tốn bộ nhớ của VPS. Tôi đề xuất sử dụng thư viện [valyala/fasthttp](https://github.com/valyala/fasthttp) để tạo ra một web server siêu đơn giản, chuyên phục vụ nội dụng web tĩnh

**Cách tiến hành như sau:**

1. Tạo go module có tên là techmaster.vn/fastlandingpage
```
go mod init techmaster.vn/fastlandingpage
```

2. Tạo mới file [fastlandingpage.go](fastlandingpage.go)
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

3. Tiếp tục gõ lệnh ở termimal để biên dịch go module
```
go build
```

4. Chạy file binary sau biên dịch
```
./fastlandingpage
```

5. Mở trình duyệt vào http://localhost:8080 sẽ thấy web site đơn giản với ảnh con mèo
![](images/Hello_World.jpg)