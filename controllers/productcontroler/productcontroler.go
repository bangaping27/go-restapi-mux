package productcontroler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bangaping27/go-restapi-mux/helper"
	"github.com/bangaping27/go-restapi-mux/models"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {

	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	ResponseJson(w, http.StatusOK, products)
}

func Show(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, fmt.Sprintf("kosong", err))
		return
	}
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Produk Tidak Ditemukan")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, product)
}

func Create(w http.ResponseWriter, r *http.Request) {

	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, product)
}

func Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Model(&product).Where("id = ?", id).Updates(&product).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, "Tidak dapat memperbarui data")
		return
	}

	product.ID = id
	ResponseJson(w, http.StatusOK, product)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	var product models.Product
	if models.DB.Delete(&product, input["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Produk Tidak Ditemukan")
		return
	}

	response := map[string]string{"message": "Produk Telah Dihapus"}
	ResponseJson(w, http.StatusOK, response)
}
