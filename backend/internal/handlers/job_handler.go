package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"openfirm/internal/models"
)

type JobHandler struct {
	jobService *models.JobService
}

func NewJobHandler(jobService *models.JobService) *JobHandler {
	return &JobHandler{
		jobService: jobService,
	}
}

type CreateJobRequest struct {
	Title        string `json:"title"`
	Company      string `json:"company"`
	Location     string `json:"location"`
	Description  string `json:"description"`
	Requirements string `json:"requirements"`
	SalaryRange  string `json:"salary_range"`
	ContactEmail string `json:"contact_email"`
	ExpiresAt    string `json:"expires_at"`
}

type JobApplicationRequest struct {
	CoverLetter string `json:"cover_letter"`
}

// Create handles job posting creation
func (h *JobHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	var req CreateJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	job := &models.Job{
		Title:        req.Title,
		Company:      req.Company,
		Location:     req.Location,
		Description:  req.Description,
		Requirements: req.Requirements,
		SalaryRange:  req.SalaryRange,
		ContactEmail: req.ContactEmail,
		PostedBy:     userID,
	}

	if err := h.jobService.CreateJob(r.Context(), job); err != nil {
		http.Error(w, "Failed to create job posting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

// Get returns a specific job posting
func (h *JobHandler) Get(w http.ResponseWriter, r *http.Request) {
	jobID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	job, err := h.jobService.GetJob(r.Context(), jobID)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

// List returns a paginated list of job postings
func (h *JobHandler) List(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit := 20
	offset := (page - 1) * limit

	jobs, err := h.jobService.ListJobs(r.Context(), offset, limit)
	if err != nil {
		http.Error(w, "Failed to fetch jobs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"jobs": jobs,
		"page": page,
	})
}

// Update handles job posting updates
func (h *JobHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	jobID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	job, err := h.jobService.GetJob(r.Context(), jobID)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	if job.PostedBy != userID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	var req CreateJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	job.Title = req.Title
	job.Company = req.Company
	job.Location = req.Location
	job.Description = req.Description
	job.Requirements = req.Requirements
	job.SalaryRange = req.SalaryRange
	job.ContactEmail = req.ContactEmail

	if err := h.jobService.UpdateJob(r.Context(), job); err != nil {
		http.Error(w, "Failed to update job posting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

// Delete removes a job posting
func (h *JobHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	jobID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	job, err := h.jobService.GetJob(r.Context(), jobID)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	if job.PostedBy != userID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	if err := h.jobService.DeleteJob(r.Context(), jobID, userID); err != nil {
		http.Error(w, "Failed to delete job posting", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Apply handles job applications
func (h *JobHandler) Apply(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	jobID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	var req JobApplicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	application := &models.JobApplication{
		JobID:       jobID,
		UserID:      userID,
		CoverLetter: req.CoverLetter,
		Status:      "pending",
	}

	if err := h.jobService.CreateJobApplication(r.Context(), application); err != nil {
		http.Error(w, "Failed to submit application", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(application)
}

// ListApplications returns all applications for a job posting
func (h *JobHandler) ListApplications(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	jobID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	job, err := h.jobService.GetJob(r.Context(), jobID)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	if job.PostedBy != userID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	applications, err := h.jobService.GetJobApplications(r.Context(), jobID)
	if err != nil {
		http.Error(w, "Failed to fetch applications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(applications)
}

// UpdateApplicationStatus updates the status of a job application
func (h *JobHandler) UpdateApplicationStatus(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	jobID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	applicationID, err := strconv.Atoi(chi.URLParam(r, "applicationId"))
	if err != nil {
		http.Error(w, "Invalid application ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	job, err := h.jobService.GetJob(r.Context(), jobID)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	if job.PostedBy != userID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	if err := h.jobService.UpdateApplicationStatus(r.Context(), applicationID, req.Status); err != nil {
		http.Error(w, "Failed to update application status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
