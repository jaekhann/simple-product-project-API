package controller

import (
	"awesomeProject/mappers"
	"awesomeProject/service"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		doGet(w, r)
	case "POST":
		doPost(w, r)
	case "PUT":
		doPut(w, r)
	case "DELETE":
		doDelete(w, r)
	default:
		doUnhandled(w, r)
	}
}

func doPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/product" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	product := mappers.FromJson(string(bytes))
	isExist := service.IsExist(product.Name)
	if isExist {
		_, _ = w.Write([]byte("product's name is not unique\n"))
	}
	id := service.CreateProduct(product)

	if !isExist {
		product.Id = uint64(id)
		w.Header().Set("Content-Type", "application/json")
		json := mappers.ToJson(product)
		_, _ = w.Write(json)
		w.WriteHeader(http.StatusCreated)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func doPut(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/product" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	product := mappers.FromJson(string(bytes))
	service.IsExist(product.Name)
	isUpdated := service.UpdateProduct(product)

	w.Header().Set("Content-Type", "text/html")
	if isUpdated {
		_, _ = w.Write([]byte("<h1>Successfully updated!</h1>"))
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotModified)
		_, _ = w.Write([]byte("<h1>Not Modified!</h1>"))
	}
}

func doGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/products" {
		products := service.GetAllProducts()
		json := mappers.ToJsons(products)
		_, _ = w.Write(json)
		w.WriteHeader(http.StatusOK)
	} else if strings.HasPrefix(r.URL.Path, "/product") {
		params := mux.Vars(r)
		id := params["id"]
		name := params["name"]

		parseInt, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Println(err)
		}
		p1 := service.GetProductById(int(parseInt))
		p2 := service.GetProductByName(name)

		if p1.Id != 0 {
			w.Header().Set("Content-Type", "application/json")
			json := mappers.ToJson(p1)
			_, err = w.Write(json)
			if err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusOK)
		} else if p2.Id != 0 {
			w.Header().Set("Content-Type", "application/json")
			json := mappers.ToJson(p2)
			_, err = w.Write(json)
			if err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusOK)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func doDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	name := params["name"]

	parseInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println(err)
	}
	p1 := service.DeleteProductById(int(parseInt))
	p2 := service.DeleteProductByName(name)

	if !p1 && !p2 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte("<h1>Successfully deleted!</h1>"))
		w.WriteHeader(http.StatusOK)
	}
}

func doUnhandled(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprint(w, "<h1>Not found!</h1>")
}
