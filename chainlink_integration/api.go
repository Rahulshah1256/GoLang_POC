package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// AdmissionService represents the admission service.
type AdmissionService struct {
	Validator *validator.Validate
	DB        *gorm.DB
}

// CreateAdmissionHandler handles the creation of a new admission form.
func (s *AdmissionService) CreateAdmissionHandler(w http.ResponseWriter, r *http.Request) {
	var admission Admission
	if err := json.NewDecoder(r.Body).Decode(&admission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the admission form
	if err := s.Validator.Struct(admission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the admission form to the database
	admission.CreatedAt = time.Now()
	admission.UpdatedAt = time.Now()
	if err := s.DB.Create(&admission).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admission)
}

// GetAdmissionHandler handles the retrieval of an admission form by ID.
func (s *AdmissionService) GetAdmissionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var admission Admission
	if err := s.DB.First(&admission, id).Error; err != nil {
		http.Error(w, "Admission not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(admission)
}

// UpdateAdmissionHandler handles the update of an admission form by ID.
func (s *AdmissionService) UpdateAdmissionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var admission Admission
	if err := json.NewDecoder(r.Body).Decode(&admission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the admission form
	if err := s.Validator.Struct(admission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the admission form in the database
	admission.ID = uint(id)
	admission.UpdatedAt = time.Now()
	if err := s.DB.Save(&admission).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(admission)
}

// DeleteAdmissionHandler handles the deletion of an admission form by ID.
func (s *AdmissionService) DeleteAdmissionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Delete the admission form from the database
	if err := s.DB.Delete(&Admission{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the admission service
	db, err := GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	validator := validator.New()
	admissionService := &AdmissionService{
		Validator: validator,
		DB:        db,
	}

	// Initialize the Chainlink integration
	chainlinkIntegration, err := NewChainlinkIntegration()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the router
	router := mux.NewRouter()

	// Register API routes
	router.HandleFunc("/admissions", admissionService.CreateAdmissionHandler).Methods("POST")
	router.HandleFunc("/admissions/{id}", admissionService.GetAdmissionHandler).Methods("GET")
	router.HandleFunc("/admissions/{id}", admissionService.UpdateAdmissionHandler).Methods("PUT")
	router.HandleFunc("/admissions/{id}", admissionService.DeleteAdmissionHandler).Methods("DELETE")

	// Start the server
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Println("Server started on port 8000")
	log.Fatal(server.ListenAndServe())
}
