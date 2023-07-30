# 文心千帆 GO SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/anhao/go-ernie.svg)](https://pkg.go.dev/github.com/anhao/go-ernie)

本库为文心千帆GO语言SDK，非官方库,目前官方还没有GO语言的SDK [文心千帆](https://cloud.baidu.com/product/wenxinworkshop)
目前支持：

* ERNIE-Bot
* ERNIE-Bot-turbo
* BLOOMZ-7B
* Llama-2
* Embeddings

### 安装

```go
go get github.com/anhao/go-ernie
```

需要 go 版本为 1.18+

### 使用示例
```go
package main

import (
	"context"
	"fmt"
	ernie "github.com/anhao/go-ernie"
)

func main() {

	client := ernie.NewDefaultClient("API Key", "Secret Key")
	completion, err := client.CreateErnieBotChatCompletion(context.Background(), ernie.ErnieBotRequest{
		Messages: []ernie.ChatCompletionMessage{
			{
				Role:    ernie.MessageRoleUser,
				Content: "你好呀",
			},
		},
	})
	if err != nil {
		fmt.Printf("ernie bot error: %v\n", err)
		return
	}
	fmt.Println(completion)
}
```

### 获取文心千帆 APIKEY
1. 在百度云官网进行申请：https://cloud.baidu.com/product/wenxinworkshop
2. 申请通过后创建应用：https://console.bce.baidu.com/qianfan/ais/console/applicationConsole/application
3. 获取 apikey 和 api secret

### 其他示例
<details>
<summary>ERNIE-Bot stream流 对话 </summary>

```go
import (
	"context"
	"errors"
	"fmt"
	ernie "github.com/anhao/go-ernie"
	"io"
)

func main() {

	client := ernie.NewDefaultClient("API Key", "Secret Key")
	request := ernie.ErnieBotRequest{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: true,
	}

	stream, err := client.CreateErnieBotChatCompletionStream(context.Background(), request)
	if err != nil {
		fmt.Printf("ernie bot stream error: %v\n", err)
		return
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("ernie bot Stream finished")
			return
		}
		if err != nil {
			fmt.Printf("ernie bot stream error: %v\n", err)
			return
		}
		fmt.Println(response.Result)
	}
}
```


</details>

<details>
<summary>ERNIE-Bot-Turbo  对话 </summary>

```go
package main

import (
	"context"
	"fmt"
	ernie "github.com/anhao/go-ernie"
)

func main() {

	client := ernie.NewDefaultClient("API Key", "Secret Key")
	completion, err := client.CreateErnieBotTurboChatCompletion(context.Background(), ernie.ErnieBotTurboRequest{
		Messages: []ernie.ChatCompletionMessage{
			{
				Role:    ernie.MessageRoleUser,
				Content: "你好呀",
			},
		},
	})
	if err != nil {
		fmt.Printf("ernie bot turbo error: %v\n", err)
		return
	}
	fmt.Println(completion)
}
```
</details>

<details>
<summary>ERNIE-Bot Turbo stream流 对话 </summary>

```go
import (
	"context"
	"errors"
	"fmt"
	ernie "github.com/anhao/go-ernie"
	"io"
)

func main() {

	client := ernie.NewDefaultClient("API Key", "Secret Key")
	request := ernie.ErnieBotTurboRequest{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: true,
	}

	stream, err := client.CreateErnieBotTurboChatCompletionStream(context.Background(), request)
	if err != nil {
		fmt.Printf("ernie bot stream error: %v\n", err)
		return
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("ernie bot turbo Stream finished")
			return
		}
		if err != nil {
			fmt.Printf("ernie bot turbo stream error: %v\n", err)
			return
		}
		fmt.Println(response.Result)
	}
}
```


</details>

<details>
<summary>BLOOMZ-7B 对话 </summary>

```go
package main

import (
	"context"
	"fmt"
	ernie "github.com/anhao/go-ernie"
)

func main() {

	client := ernie.NewDefaultClient("API Key", "Secret Key")
	completion, err := client.CreateBloomz7b1ChatCompletion(context.Background(), ernie.Bloomz7b1Request{
		Messages: []ernie.ChatCompletionMessage{
			{
				Role:    ernie.MessageRoleUser,
				Content: "你好呀",
			},
		},
	})
	if err != nil {
		fmt.Printf("BLOOMZ-7B error: %v\n", err)
		return
	}
	fmt.Println(completion)
}
```
</details>

<details>
<summary>BLOOMZ-7B stream流 对话 </summary>

```go
import (
	"context"
	"errors"
	"fmt"
	ernie "github.com/anhao/go-ernie"
	"io"
)

func main() {

	client := ernie.NewDefaultClient("API Key", "Secret Key")
	request := ernie.Bloomz7b1Request{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: true,
	}

	stream, err := client.CreateBloomz7b1ChatCompletionStream(context.Background(), request)
	if err != nil {
		fmt.Printf("BLOOMZ-7B error: %v\n", err)
		return
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("BLOOMZ-7B  Stream finished")
			return
		}
		if err != nil {
			fmt.Printf("BLOOMZ-7B stream error: %v\n", err)
			return
		}
		fmt.Println(response.Result)
	}
}
```

</details>


<details>
<summary>Llama2 对话 </summary>

```go
package main

import (
	"context"
	"fmt"
	ernie "github.com/anhao/go-ernie"
)

func main() {

	client := ernie.NewDefaultClient("API Key", "Secret Key")
	completion, err := client.CreateLlamaChatCompletion(context.Background(), ernie.LlamaChatRequest{
		Messages: []ernie.ChatCompletionMessage{
			{
				Role:    ernie.MessageRoleUser,
				Content: "你好呀",
			},
		},
		Model: "", //申请发布时填写的API名称
	})
	if err != nil {
		fmt.Printf("llama2 error: %v\n", err)
		return
	}
	fmt.Println(completion)
}
```
</details>

<details>
<summary>Llama2 stream流 对话 </summary>

```go
import (
	"context"
	"errors"
	"fmt"
	ernie "github.com/anhao/go-ernie"
	"io"
)

func main() {

	client := ernie.NewDefaultClient("API Key", "Secret Key")
	request := ernie.LlamaChatRequest{
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: true,
		Model: "", //申请发布时填写的API名称
	}

	stream, err := client.CreateLlamaChatCompletionStream(context.Background(), request)
	if err != nil {
		fmt.Printf("llama2 error: %v\n", err)
		return
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("llama2  Stream finished")
			return
		}
		if err != nil {
			fmt.Printf("llama2 stream error: %v\n", err)
			return
		}
		fmt.Println(response.Result)
	}
}
```

</details>

<details>
<summary>Embedding向量</summary>

```go
package main

import (
	"context"
	"fmt"
	ernie "github.com/anhao/go-ernie"
)

func main() {

	client := ernie.NewDefaultClient("API Key", "Secret Key")
	request := ernie.EmbeddingRequest{
		Input: []string{
			"Hello",
		},
	}
	embeddings, err := client.CreateEmbeddings(context.Background(), request)
	if err != nil {
		fmt.Sprintf("embeddings err: %v", err)
		return
	}
	fmt.Println(embeddings)
}

```

</details>

<details>
<summary>自定义 accessToken</summary>

```go
client :=ernie.NewClient("accessToken")
```

</details>