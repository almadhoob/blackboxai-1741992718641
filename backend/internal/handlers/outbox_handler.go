package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"openfirm/internal/activitypub"
)

type OutboxHandler struct {
	activityPubService *activitypub.Service
}

func NewOutboxHandler(activityPubService *activitypub.Service) *OutboxHandler {
	return &OutboxHandler{
		activityPubService: activityPubService,
	}
}

// Get returns the contents of a user's outbox
func (h *OutboxHandler) Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	// Verify Accept header
	accept := r.Header.Get("Accept")
	if !strings.Contains(accept, "application/activity+json") &&
		!strings.Contains(accept, "application/ld+json") {
		http.Error(w, "Not Acceptable", http.StatusNotAcceptable)
		return
	}

	// Get page number from query params
	page := 1
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	// Get outbox contents
	outbox, err := h.activityPubService.GetOutbox(r.Context(), username, page)
	if err != nil {
		http.Error(w, "Failed to get outbox", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	json.NewEncoder(w).Encode(outbox)
}

// Post handles outgoing ActivityPub activities
func (h *OutboxHandler) Post(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	// Verify Content-Type header
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/activity+json") &&
		!strings.Contains(contentType, "application/ld+json") {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parse the activity
	var activity map[string]interface{}
	if err := json.Unmarshal(body, &activity); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate the activity
	if err := h.validateOutboxActivity(activity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process the activity
	result, err := h.activityPubService.HandleOutbox(r.Context(), username, activity)
	if err != nil {
		http.Error(w, "Failed to process activity", http.StatusInternalServerError)
		return
	}

	// Return the created activity
	w.Header().Set("Content-Type", "application/activity+json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// validateOutboxActivity validates an outgoing activity
func (h *OutboxHandler) validateOutboxActivity(activity map[string]interface{}) error {
	// Verify required fields
	required := []string{"@context", "type"}
	for _, field := range required {
		if _, ok := activity[field]; !ok {
			return fmt.Errorf("missing required field: %s", field)
		}
	}

	// Verify activity type is supported
	activityType, _ := activity["type"].(string)
	supportedTypes := map[string]bool{
		"Create":   true,
		"Update":   true,
		"Delete":   true,
		"Follow":   true,
		"Unfollow": true,
		"Like":     true,
		"Unlike":   true,
		"Share":    true,
	}

	if !supportedTypes[activityType] {
		return fmt.Errorf("unsupported activity type: %s", activityType)
	}

	// Validate object field for Create and Update activities
	if activityType == "Create" || activityType == "Update" {
		obj, ok := activity["object"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid object field for %s activity", activityType)
		}

		// Validate object type
		objType, _ := obj["type"].(string)
		supportedObjectTypes := map[string]bool{
			"Note":        true,
			"Article":     true,
			"JobPosting": true,
		}

		if !supportedObjectTypes[objType] {
			return fmt.Errorf("unsupported object type: %s", objType)
		}
	}

	return nil
}

// deliverActivity delivers an activity to recipients
func (h *OutboxHandler) deliverActivity(ctx context.Context, activity map[string]interface{}, recipients []string) error {
	// TODO: Implement activity delivery
	// This should:
	// 1. Sign the activity with the sender's private key
	// 2. Send the activity to each recipient's inbox
	// 3. Handle delivery failures and retries
	return nil
}

// handleCreateActivity processes Create activities in the outbox
func (h *OutboxHandler) handleCreateActivity(ctx context.Context, activity map[string]interface{}) error {
	// Extract the object being created
	obj, ok := activity["object"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid object in Create activity")
	}

	// Process based on object type
	switch objType := obj["type"].(string); objType {
	case "Note":
		return h.handleCreateNote(ctx, obj)
	case "Article":
		return h.handleCreateArticle(ctx, obj)
	case "JobPosting":
		return h.handleCreateJobPosting(ctx, obj)
	default:
		return fmt.Errorf("unsupported object type: %s", objType)
	}
}

// handleCreateNote processes the creation of a Note object
func (h *OutboxHandler) handleCreateNote(ctx context.Context, note map[string]interface{}) error {
	// TODO: Implement Note creation
	// This should:
	// 1. Validate the note content
	// 2. Store the note in the database
	// 3. Determine recipients
	// 4. Deliver the activity to recipients
	return nil
}

// handleCreateArticle processes the creation of an Article object
func (h *OutboxHandler) handleCreateArticle(ctx context.Context, article map[string]interface{}) error {
	// TODO: Implement Article creation
	// This should:
	// 1. Validate the article content
	// 2. Store the article in the database
	// 3. Determine recipients
	// 4. Deliver the activity to recipients
	return nil
}

// handleCreateJobPosting processes the creation of a JobPosting object
func (h *OutboxHandler) handleCreateJobPosting(ctx context.Context, jobPosting map[string]interface{}) error {
	// TODO: Implement JobPosting creation
	// This should:
	// 1. Validate the job posting content
	// 2. Store the job posting in the database
	// 3. Determine recipients
	// 4. Deliver the activity to recipients
	return nil
}

// Additional helper methods for handling other activity types would be implemented here
