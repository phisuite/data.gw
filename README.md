# Phi Suite Data.gw

| **Homepage** | [https://phisuite.com][0]        |
| ------------ | -------------------------------- | 
| **GitHub**   | [https://github.com/phisuite][1] |

## Overview

This project contains the Go module to create the **Data http gateway**.  
For more details, see [Phi Suite Data][2].

## Installation

```bash
go get github.com/phisuite/data.gw
```

## Creating the gateway

```go
package main
import "github.com/phisuite/data.gw"

conn, err := grpc.DialContext(ctx, "localhost:50051")
defer conn.Close()

router := runtime.NewServeMux()
err = data.RegisterEventAPIHandler(ctx, router, conn)
err = http.ListenAndServe(":8080", router)
```
For more details, see [grpc-gateway][10].

[0]: https://phisuite.com
[1]: https://github.com/phisuite
[2]: https://github.com/phisuite/data
[10]: https://grpc-ecosystem.github.io/grpc-gateway
__