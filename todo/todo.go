package todo

import (
  "net/http"

  "github.com/go-chi/chi"
  "github.com/go-chi/render"
)

type Todo struct {
  Slug  string `json:"slug"`
  Title string `json:"title"`
  Body  string `json:"body"`
}

func Routes() *chi.Mux {
  router := chi.NewRouter()
  router.Get("/{todoID}", GetATodo)
  router.Delete("/{todoID}", DeleteTodo)
  router.Post("/", CreateTodo)
  router.Get("/", GetAllTodos)
  return router
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
  todoID := chi.URLParam(r, "todoID")
  todos := Todo{
    Slug: todoID,
    Title: "Suh world!",
    Body: "okay i was gone for a minute but i'm back now, could you please sit down?",
  }
  render.JSON(w, r, todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
  todoID := chi.URLParam(r, "todoID")
  response := make(map[string]string)
  response["message"] = "Delete TODO " + todoID + " successfully"
  render.JSON(w, r, response)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
  response := make(map[string]string)
  response["message"] = "Created TODO successfully"
  render.JSON(w, r, response)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
  todos := []Todo{
    {
      Slug:  "slug",
      Title: "hwrbgrbr",
      Body:  "efgwrhrhen tne netyetyhe",
    },
  }
  render.JSON(w, r, todos)
}

