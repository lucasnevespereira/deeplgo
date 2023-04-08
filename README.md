# Deeplgo

A Go client library for **[Deepl](https://www.deepl.com)** Translator API.

<img src="https://img.shields.io/github/go-mod/go-version/lucasnevespereira/deeplgo">

## Usage

```
go get -u github.com/lucasnevespereira/deeplgo
```
### Initializing the Client
To get an API KEY, [sign up here](https://www.deepl.com/pro#developer) for a Deepl Developer Account.
```go
deeplClient := deepl.NewClient("YOUR_DEEPL_API_KEY")
```

### Translate
```go
// Prepare one or multiple texts to translate in an array of strings
texts := []string{"I am a human", "I like salty food"}

// Call the Translate method passing in the following arguments:
// texts, source language (optional because it can be detected), target language
response, err := deeplClient.Translate(texts, "en", "fr")
if err != nil {
fmt.Println(err)
}
```

### Results

Results are an array of Transalation objects

```go
type TranslateResponse struct {
	Translations []Translation `json:"translations"`
}

type Translation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}
```

You can print the actual translation text this way
```go
fmt.Println(response.Translations[0].Text) // Je suis un être humain
fmt.Println(response.Translations[1].Text) // J'aime la nourriture salée
```


### Full Example
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



## How to Contribute

If you want to contribute to this project please read the [Contributing](CONTRIBUTING.md)

<hr>

## License

This project is under [BSD 3-Clause License](LICENSE)


