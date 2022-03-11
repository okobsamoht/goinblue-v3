[![GoDoc](https://godoc.org/github.com/JesusIslam/goinblue?status.svg)](https://godoc.org/github.com/JesusIslam/goinblue)

# Goinblue
---------------------------------
### A work in progress package for golang to send email and SMS through sendinblue.com

#### Usage
Example:

```
package main

import (
	"github.com/JesusIslam/goinblue"
	"fmt"
)

func main() {
	myApiKey := "thisisyourapikey"

	email := &goinblue.Email{
		To: []map[string]string{
			{"email":"to@example.com": "name":"Mr. To"},
		},
		Subject: "Test",
		Sender: map[string]string{
			"email":sender@example.com", "name":Sender",
		},
		Text: "This is just a test.",
	}

	client := goinblue.NewClient(myApiKey)
	res, err := client.SendEmail(email)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```
