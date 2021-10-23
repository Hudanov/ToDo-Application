package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ToDo serves map of userIds to their ToDoStorages
type ToDo struct {
	todos map[int]UserToDo
}

type UserToDo struct {
	lists map[int]ToDoList // Map listId -> List
}

type ToDoList struct {
	Id    int          `json:"id"`
	Name  string       `json:"name"`
	tasks map[int]Task // Map taskId -> Task
}

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"task_name"`
	Description string `json:"task_description"`
	Status      string `json:"status"`
}

type CreateTaskParametrs struct {
	List_id          int
	Task_name        string `json:"task_name"`
	Task_description string `json:"task_description"`
}

type UpdateTaskParametrs struct {
	Task_name string `json:"task_name"`
	Status    string `json:"status"`
}

type CreateListParametrs struct {
	List_name string `json:"list_name"`
}

var globalTodos ToDo
var userCounter int
var listsCounter int
var tasksCounter int

func createNewList(user User, params CreateListParametrs) ToDoList {

	listsCounter++
	list := ToDoList{
		Id:    listsCounter,
		Name:  params.List_name,
		tasks: make(map[int]Task),
	}

	globalTodos.todos[user.id].lists[listsCounter] = list

	return list

}

func deleteList(user User, list_id int) {
	userTodos := globalTodos.todos[user.id]
	delete(userTodos.lists, list_id)
}

func createNewTask(user User, params CreateTaskParametrs) Task {
	userTodos := globalTodos.todos[user.id]
	list := userTodos.lists[params.List_id]
	tasksCounter++
	taskId := tasksCounter
	task := Task{
		Id:          taskId,
		Name:        params.Task_name,
		Description: params.Task_description,
		Status:      "open",
	}

	if userTodos.lists == nil {
		list.tasks = make(map[int]Task)
	}
	list.tasks[taskId] = task

	return task
}

func deleteTask(user User, list_id int, task_id int) {
	userTodos := globalTodos.todos[user.id]
	list := userTodos.lists[list_id]
	delete(list.tasks, task_id)
}

func updateTask(user User, list_id int, task_id int, params UpdateTaskParametrs) {

	if entry, ok := globalTodos.todos[user.id].lists[list_id].tasks[task_id]; ok {
		entry.Name = params.Task_name
		entry.Status = params.Status

		globalTodos.todos[user.id].lists[list_id].tasks[task_id] = entry

	}

}

func updateList(user User, list_id int, params CreateListParametrs) {

	if entry, ok := globalTodos.todos[user.id].lists[list_id]; ok {

		// Then we modify the copy
		entry.Name = params.List_name

		// Then we reassign map entry
		globalTodos.todos[user.id].lists[list_id] = entry
	}
}

func (u *UserService) createNewTaskHandler(w http.ResponseWriter, r *http.Request, user User) {
	task := &CreateTaskParametrs{}

	params := mux.Vars(r)
	list_id, _ := strconv.Atoi(params["list_id"])

	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		fmt.Println(err)
	}

	jData, err := json.Marshal(createNewTask(user, CreateTaskParametrs{
		List_id:          list_id,
		Task_name:        task.Task_name,
		Task_description: task.Task_description,
	}))
	if err != nil {
		// handle error
		fmt.Printf("went wrong")
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jData))
}

func (u *UserService) createNewListHandler(w http.ResponseWriter, r *http.Request, user User) {
	list := &CreateListParametrs{}

	err := json.NewDecoder(r.Body).Decode(list)
	if err != nil {
		fmt.Println(err)
	}

	jData, err := json.Marshal(createNewList(user, *list))
	if err != nil {
		// handle error
		fmt.Printf("went wrong")
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jData))
}

func (u *UserService) deleteListHandler(w http.ResponseWriter, r *http.Request, user User) {
	params := mux.Vars(r)
	list_id, _ := strconv.Atoi(params["list_id"])

	deleteList(user, list_id)

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Deleted"))
}

func (u *UserService) deleteTaskHandler(w http.ResponseWriter, r *http.Request, user User) {
	params := mux.Vars(r)
	list_id, _ := strconv.Atoi(params["list_id"])
	task_id, _ := strconv.Atoi(params["task_id"])

	deleteTask(user, list_id, task_id)

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Deleted"))
}

func (u *UserService) getListsHandler(w http.ResponseWriter, r *http.Request, user User) {
	userTodos := globalTodos.todos[user.id]

	jData := make([]byte, len(userTodos.lists))

	for _, v := range userTodos.lists {
		tmp, err := json.Marshal(v)
		if err != nil {
			// handle error
			fmt.Printf("went wrong")
		}

		jData = append(jData, tmp...)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func (u *UserService) getTasksHandler(w http.ResponseWriter, r *http.Request, user User) {
	params := mux.Vars(r)
	list_id, _ := strconv.Atoi(params["list_id"])

	tasks := globalTodos.todos[user.id].lists[list_id].tasks

	jData, err := json.Marshal(tasks)
	if err != nil {
		// handle error
		fmt.Printf("went wrong")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func (u *UserService) updateTaskHandler(w http.ResponseWriter, r *http.Request, user User) {
	params := mux.Vars(r)
	list_id, _ := strconv.Atoi(params["list_id"])
	task_id, _ := strconv.Atoi(params["task_id"])

	task := &UpdateTaskParametrs{}

	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		fmt.Println(err)
	}

	updateTask(user, list_id, task_id, *task)

	response := globalTodos.todos[user.id].lists[list_id].tasks[task_id]

	jData, err := json.Marshal(response)
	if err != nil {
		// handle error
		fmt.Printf("went wrong")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jData)
}

func (u *UserService) updateListHandler(w http.ResponseWriter, r *http.Request, user User) {
	params := mux.Vars(r)
	list_id, _ := strconv.Atoi(params["list_id"])

	list := &CreateListParametrs{}

	err := json.NewDecoder(r.Body).Decode(list)
	if err != nil {
		fmt.Println(err)
	}

	updateList(user, list_id, *list)

	response := globalTodos.todos[user.id].lists[list_id]

	jData, err := json.Marshal(response)
	if err != nil {
		// handle error
		fmt.Printf("went wrong")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
