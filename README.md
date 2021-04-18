# PayJS Go SDK

## Native 扫码支付

```go
package main
import (
	"rinne.dev/payjs"
	"rinne.dev/payjs/native"
)
p := payjs.New(...)
req := native.CreateRequest(...)
resp, err := p.Native(req)
```
