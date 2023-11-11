package web

import "net/http"

func(app *Application) ProccesFileTransactioHandler(w http.ResponseWriter, r *http.Request) {
	//limitar tamanho da requisicao para 32mb
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)
	//arquivo de upload de banner maximo 10MB
	r.ParseMultipartForm(10 << 20)
	file,_,err := r.FormFile("file")
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	defer file.Close()

	output,err := app.ProccessFileUseCase.Execute(r.Context(),file)
	if err != nil {
		app.errorJson(w,err)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}