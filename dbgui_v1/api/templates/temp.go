package api

import (
	"fmt"
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/templates/page-login.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("api/templates/page-login.html")
	if err != nil {
		fmt.Println("----9999999999999999999---", err)
		return
	}
	//	data := models.Rdsdata{}
	//	datas, err := data.FindAllData(server.DB)

	//responses.JSON(w, http.StatusOK, datas)

	fmt.Println("----9999---")

	tmpl.Execute(w, nil)
}
