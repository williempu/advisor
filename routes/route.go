package routes

import (
	"advisor/app/lecturer"
	"advisor/app/student"
	"advisor/app/transaction"
	"database/sql"
	"embed"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRoutes(db *sql.DB, fs embed.FS) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// initiate here
	studentRepo := student.NewRepo(db)

	lecturerRepo := lecturer.NewRepo(db)
	// lecturerController := lecturer.NewController(lecturerRepo)

	transactionRepo := transaction.NewRepo(db)

	transactionController := transaction.NewController(studentRepo, lecturerRepo, transactionRepo)

	// TODO FIXME use transaction instead of lecturer
	r.Route("/lecturer", func(lecturerRoute chi.Router) {
		lecturerRoute.Get("/", transactionController.GetLecturerQuota)
	})

	r.Route("/transaction", func(trxRoute chi.Router) {
		trxRoute.Get("/all", transactionController.GetAll)
		trxRoute.Post("/create", transactionController.Create)
		trxRoute.Post("/checkStudent", transactionController.GetByStudentId)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var templates = template.Must(template.ParseFS(fs, "templates/index.html"))
		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Fatalln(err.Error())
			return
		}
		// response.RenderHTML(w, r, "", nil)
	})

	r.Get("/transactions", func(w http.ResponseWriter, r *http.Request) {
		var templates = template.Must(template.ParseFS(fs, "templates/statistics.html"))
		if err := templates.ExecuteTemplate(w, "statistics.html", nil); err != nil {
			log.Fatalln(err.Error())
			return
		}
		// response.RenderHTML(w, r, "", nil)
	})

	return r
}
