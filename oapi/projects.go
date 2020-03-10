package oapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	ogit "github.com/danzim/opro/ogit"
	"github.com/gorilla/mux"
)

// CreateProject - Creates a project in OpenShift with JSON request
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var newProject ogit.Namespace
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter data with CI, Name and Description of a OpenShift")
	}
	projects := ogit.GetProjects()
	json.Unmarshal(reqBody, &newProject)
	projects = append(projects, newProject)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newProject)
}

// GetOneProject - Get dedicated project in API
func GetOneProject(w http.ResponseWriter, r *http.Request) {
	projectCI := mux.Vars(r)["ci"]
	projects := ogit.GetProjects()

	for _, singleProject := range projects {
		if singleProject.CI == projectCI {
			json.NewEncoder(w).Encode(singleProject)
		}
	}
}

// GetAllProjects - List all Projects in API
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := ogit.GetProjects()
	json.NewEncoder(w).Encode(projects)
}

// UpdateProject - Update dedicated project in API
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	projectCI := mux.Vars(r)["ci"]
	var updatedProject ogit.Namespace

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter data with the project name and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedProject)

	projects := ogit.GetProjects()
	for i, singleProject := range projects {
		if singleProject.CI == projectCI {
			singleProject.DisplayName = updatedProject.DisplayName
			singleProject.Description = updatedProject.Description
			projects = append(projects[:i], singleProject)
			json.NewEncoder(w).Encode(singleProject)
		}
	}
}

// DeleteProject - Delete Project in API
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	projectCI := mux.Vars(r)["ci"]
	projects := ogit.GetProjects()

	for i, singleProject := range projects {
		if singleProject.CI == projectCI {
			projects = append(projects[:i], projects[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", projectCI)
		}
	}
}
