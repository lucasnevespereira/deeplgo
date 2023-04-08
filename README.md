# Deeplgo

Go unofficial client lib for Deepl API.

## Usage

```
    go get -u github.com/lucasnevespereira/deeplgo/deepl
```

```go
package main

import (
	"fmt"
	"github.com/lucasnevespereira/deeplgo/deepl"
)

func main() {
	deeplClient := deepl.NewClient("YOUR_DEEPL_API_KEY")
	texts := []string{"I am a human", "I like salty food"}
	response, err := deeplClient.Translate(texts, "en", "fr")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response.Translations[0].Text) // Je suis un être humain
	fmt.Println(response.Translations[1].Text) // J'aime la nourriture salée

}
```
