# html-link-parser  

To use this package, you can simply call Extract with an io.Reader that contains the HTML you want to parse. For example:  

```
package main

import (
    "fmt"
    "strings"

    "github.com/somatosensory/html-link-parser"
)

func main() {
    html := `
        <html>
        <body>
            <a href="/dog">dog</a>
            <a href="/cat">cat <b>bold</b></a>
        </body>
        </html>
    `
    r := strings.NewReader(html)
    links, err := link.Extract(r)
    if err != nil {
        panic(err)
    }
    for _, l := range links {
        fmt.Printf("Link: %v (Text: %v)\n", l.Href, l.Text)
    }
}
```

This will output:

```
Link: /dog (Text: dog)
Link: /cat (Text: cat bold)
```
