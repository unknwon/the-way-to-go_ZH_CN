package hello

import (
	"fmt"
	"net/http"
	"template"
)

const guestbookForm = `
<html>
  <body>
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`
const signTemplateHTML = `
<html>
  <body>
    <p>You wrote:</p>
    <pre>{{html .}}</pre>
  </body>
</html>
`

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, guestbookForm)
}

func sign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}
}
