// use_urlshortener.go
package main

import (
	"fmt"
	"net/http"
	"text/template"
	urlshortener "code.google.com/p/google-api-go-client/urlshortener/v1" 
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/short", short)
	http.HandleFunc("/long", long)
	http.ListenAndServe("localhost:8080", nil)
}

// the template used to show the forms and the results web page to the user
var rootHtmlTmpl = template.Must(template.New("rootHtml").Parse(`
<html><body>
<h1>URL SHORTENER</h1>
{{if .}}{{.}}<br /><br />{{end}}
<form action="/short" type="POST">
Shorten this: <input type="text" name="longUrl" />
<input type="submit" value="Give me the short URL" />
</form>
<br />
<form action="/long" type="POST">
Expand this: http://goo.gl/<input type="text" name="shortUrl" />
<input type="submit" value="Give me the long URL" />
</form>
</body></html>
`))

func root(w http.ResponseWriter, r *http.Request) {
	rootHtmlTmpl.Execute(w, nil)
}

func short(w http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("longUrl")                          
	urlshortenerSvc, _ := urlshortener.New(http.DefaultClient) 
	url, _ := urlshortenerSvc.Url.Insert(&urlshortener.Url{LongUrl: longUrl,}).Do() 
	rootHtmlTmpl.Execute(w, fmt.Sprintf("Shortened version of %s is : %s", longUrl, url.Id))
}

func long(w http.ResponseWriter, r *http.Request) {
	shortUrl := "http://goo.gl/" + r.FormValue("shortUrl")
	urlshortenerSvc, _ := urlshortener.New(http.DefaultClient)
	url, err := urlshortenerSvc.Url.Get(shortUrl).Do()
	if err != nil {
		fmt.Println("error: %v", err)
		return
	}
	rootHtmlTmpl.Execute(w, fmt.Sprintf("Longer version of %s is : %s", shortUrl, url.LongUrl))
}