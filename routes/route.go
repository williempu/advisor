package routes

import (
	"advisor/app/lecturer"
	"advisor/app/student"
	"advisor/app/transaction"
	"advisor/lib/response"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRoutes(db *sql.DB) http.Handler {
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
		trxRoute.Post("/create", transactionController.Create)
		trxRoute.Post("/checkStudent", transactionController.GetByStudentId)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response.RenderHTML(w, r, "", nil)
	})

	return r
}
