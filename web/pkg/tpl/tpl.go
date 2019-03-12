package tpl

import (
	"encoding/json"
	"time"
	"net/http"
)

// Data struct
type Data struct {
	Data interface{}
}

// Render func
// func (s Data) Render(w web.ResponseWriter, r *web.Request) error {
// 	if s.TemplateFile == "" {
// 		tplErr := errors.New("tpl.Render requires non-empty TemplateFile")
// 		return tplErr
// 	}
//
// 	t := template.Must(template.ParseFiles("templates/"+s.TemplateFile, "templates/layout.html"))
// 	err := t.ExecuteTemplate(w, "base", s)
//
// 	return err
// }

// Json Function
func (d Data) RenderJson(w web.ResponseWriter, r *web.Request) error {
	data := d.Data{}
	err := json.NewDecoder(r.Body).decode(&data)
	if err != nil {
		panic(err)
	}
data.CreatedAt:=time.Now().Local()
dataJson = json.Marshal(data)
if err!=nil{
	panic(err)
}
//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(dataJson)
}
