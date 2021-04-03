# Tạo landing page siêu nhanh , nhẹ với [valyala/fasthttp](https://github.com/valyala/fasthttp)

Hiện nay Techmaster sử dụng iris framework cũng code rất dài dòng và thừa chỉ để phục vụ langding page tĩnh. Việc này khiến cho Docker image size vừa lớn, tốc độ khởi động chậm và tốn bộ nhớ của VPS. Tôi đề xuất sử dụng thư viện [valyala/fasthttp](https://github.com/valyala/fasthttp) để tạo ra một web server siêu đơn giản, chuyên phục vụ nội dụng web tĩnh

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

## Thử benchmark web site này

Sử dụng tool [https://github.com/codesenberg/bombardier](https://github.com/codesenberg/bombardier)

```
bombardier -c 125 -n 10000000 http://localhost:8080
```

Kết quả
```
Bombarding http://localhost:8080 with 10000000 request(s) using 125 connection(s)
 10000000 / 10000000 [==================================] 100.00% 104070/s 1m36s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec    104109.46    5062.08  121245.54
  Latency        1.20ms    90.39us    24.52ms
  HTTP codes:
    1xx - 0, 2xx - 10000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    55.90MB/s
```

Tốc độ phục vụ rất tốt