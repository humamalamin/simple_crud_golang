package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/crud_simple_go/helpers"
	"gitlab.com/crud_simple_go/models"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Methode initialize berguna untuk koneksi ke database
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRouter()
}

func (a *App) initializeRouter() {
	a.Router.HandleFunc("/group_businesses", a.getGroupBusinesses).Methods("GET")
	a.Router.HandleFunc("/group_businesses", a.createGroupBusiness).Methods("POST")
	a.Router.HandleFunc("/group_business/{id:[0-9+]}", a.getGroupBusiness).Methods("GET")
	a.Router.HandleFunc("/group_business/{id:[0-9+]}", a.updateGroupBusiness).Methods("PUT")
	a.Router.HandleFunc("/group_business/{id:[0-9+]}", a.deleteGroupBusiness).Methods("DELETE")
}

// MEthod run digunakan untuk menjalankan script golang
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// Menampilkan group business dengan single row
func (a *App) getGroupBusiness(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrBadParamInput), "ID group business salah")
		return
	}

	gb := models.GroupBusiness{ID: id}
	if err := gb.GetGroupBusiness(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrNotFound), "Group business tidak ditemukan")
		default:
			helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrInternalServerError), err.Error())
		}

		return
	}

	data := helpers.Response{
		Status:  int(helpers.GetStatusCode(err)),
		Message: "Success",
		Data:    gb,
	}

	helpers.RespondWithJSON(w, http.StatusOK, data)
}

// Menampilkan all data group business
func (a *App) getGroupBusinesses(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	if count > 10 || count < 1 {
		count = 10
	}

	if start < 0 {
		start = 0
	}

	groupBusinesses, err := models.GetGroupBusinesses(a.DB, start, count)
	if err != nil {
		helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrInternalServerError), err.Error())
		return
	}

	data := helpers.Response{
		Status:  int(helpers.GetStatusCode(err)),
		Message: "Success",
		Data:    groupBusinesses,
	}

	helpers.RespondWithJSON(w, http.StatusOK, data)
}

// Menyimpan data group business baru
func (a *App) createGroupBusiness(w http.ResponseWriter, r *http.Request) {
	var gb models.GroupBusiness
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&gb); err != nil {
		helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrBadParamInput), "request error")
		return
	}

	defer r.Body.Close()
	err := gb.CreateGroupBusiness(a.DB)
	if err != nil {
		helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrInternalServerError), err.Error())
		return
	}

	data := helpers.Response{
		Status:  int(helpers.GetStatusCode(err)),
		Message: "Success",
		Data:    gb,
	}

	helpers.RespondWithJSON(w, http.StatusCreated, data)
}

// Mengubah data group business
func (a *App) updateGroupBusiness(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrBadParamInput), "ID group business salah")
		return
	}

	var gb models.GroupBusiness
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&gb); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "request tidak valid")
		return
	}
	defer r.Body.Close()
	gb.ID = id
	if err := gb.UpdateGroupBusiness(a.DB); err != nil {
		helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrInternalServerError), err.Error())
		return
	}

	data := helpers.Response{
		Status:  int(helpers.GetStatusCode(err)),
		Message: "Success",
		Data:    gb,
	}

	helpers.RespondWithJSON(w, http.StatusOK, data)
}

// Menghapus datan group business
func (a *App) deleteGroupBusiness(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrBadParamInput), "ID group business salah")
		return
	}

	gb := models.GroupBusiness{ID: id}
	if err := gb.DeleteGroupBusiness(a.DB); err != nil {
		helpers.RespondWithError(w, helpers.GetStatusCode(helpers.ErrInternalServerError), err.Error())
		return
	}

	data := helpers.Response{
		Status:  int(helpers.GetStatusCode(err)),
		Message: "Success",
		Data:    nil,
	}

	helpers.RespondWithJSON(w, http.StatusOK, data)
}
