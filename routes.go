package configstore

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ApiRouter(service *KVService) chi.Router {
	r := chi.NewRouter()

	r.Get("/kv", listKVs(service))
	r.Get("/kv/{key}", getKV(service))
	r.Post("/kv", createKV(service))
	r.Put("/kv/{key}", updateKV(service))
	r.Delete("/kv/{key}", deleteKV(service))

	return r
}

func listKVs(service *KVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		kvs, err := service.ListKVS()
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.JSON(w, r, map[string]interface{}{"data": kvs})
	}
}

func getKV(service *KVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		kv, err := service.GetKV(key)
		if err != nil {
			if err.Error() == "key not found" {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusInternalServerError)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.JSON(w, r, map[string]interface{}{"data": kv})
	}
}

func createKV(service *KVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		if err := render.DecodeJSON(r.Body, &req); err != nil || req.Key == "" || req.Value == "" {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": "both key and value fields are required"})
			return
		}
		kv, err := service.CreateKV(req.Key, req.Value)
		if err != nil {
			if err.Error() == "key already present in DB" {
				render.Status(r, http.StatusBadRequest)
			} else {
				render.Status(r, http.StatusInternalServerError)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]interface{}{"data": kv})
	}
}

func updateKV(service *KVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		var req struct {
			Value string `json:"value"`
		}
		if err := render.DecodeJSON(r.Body, &req); err != nil || req.Value == "" {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": "value field is required"})
			return
		}

		kv, err := service.UpdateKV(key, req.Value)
		if err != nil {
			if err.Error() == "key not found" {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusInternalServerError)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.JSON(w, r, map[string]interface{}{"data": kv})
	}
}

func deleteKV(service *KVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		if err := service.DeleteKV(key); err != nil {
			if err.Error() == "key not found" {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusInternalServerError)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
