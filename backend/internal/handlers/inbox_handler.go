package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"openfirm/internal/activitypub"
)

type InboxHandler struct {
	activityPubService *activitypub.Service
}

func NewInboxHandler(activityPubService *activitypub.Service) *InboxHandler {
	return &InboxHandler{
		activityPubService: activityPubService,
	}
}

// Get returns the contents of a user's inbox
func (h *InboxHandler) Get(w http.ResponseWriter, r *http.Request) {
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

	// Get inbox contents
	inbox, err := h.activityPubService.GetInbox(r.Context(), username, page)
	if err != nil {
		http.Error(w, "Failed to get inbox", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	json.NewEncoder(w).Encode(inbox)
}

// Post handles incoming ActivityPub activities
func (h *InboxHandler) Post(w http.ResponseWriter, r *http.Request) {
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

	// Verify signature (if implemented)
	if err := h.verifyHttpSignature(r); err != nil {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// Process the activity
	if err := h.activityPubService.HandleInbox(r.Context(), username, body); err != nil {
		http.Error(w, "Failed to process activity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// verifyHttpSignature verifies the HTTP Signature of the request
func (h *InboxHandler) verifyHttpSignature(r *http.Request) error {
	// TODO: Implement HTTP Signature verification
	// This should:
	// 1. Parse the Signature header
	// 2. Fetch the public key of the sender
	// 3. Verify the signature against the request
	return nil
}

// validateActivity validates an incoming activity
func (h *InboxHandler) validateActivity(activity map[string]interface{}) error {
	// Verify required fields are present
	required := []string{"@context", "type", "actor"}
	for _, field := range required {
		if _, ok := activity[field]; !ok {
			return fmt.Errorf("missing required field: %s", field)
		}
	}

	// Verify activity type is supported
	activityType, _ := activity["type"].(string)
	supportedTypes := map[string]bool{
		"Create":  true,
		"Follow":  true,
		"Accept":  true,
		"Reject":  true,
		"Delete":  true,
		"Update":  true,
		"Undo":    true,
		"Like":    true,
		"Announce": true,
	}

	if !supportedTypes[activityType] {
		return fmt.Errorf("unsupported activity type: %s", activityType)
	}

	return nil
}

// handleDeliveryErrors processes errors that occur during activity delivery
func (h *InboxHandler) handleDeliveryErrors(ctx context.Context, err error, activity map[string]interface{}) {
	// Log the error
	log.Printf("Failed to process activity: %v", err)

	// TODO: Implement retry logic for failed deliveries
	// This could include:
	// 1. Storing failed deliveries in a queue
	// 2. Implementing exponential backoff
	// 3. Setting maximum retry attempts
	// 4. Notifying admins of persistent failures
}

// processActivity processes different types of activities
func (h *InboxHandler) processActivity(ctx context.Context, activity map[string]interface{}) error {
	activityType, _ := activity["type"].(string)

	switch activityType {
	case "Create":
		return h.handleCreate(ctx, activity)
	case "Follow":
		return h.handleFollow(ctx, activity)
	case "Accept":
		return h.handleAccept(ctx, activity)
	case "Reject":
		return h.handleReject(ctx, activity)
	case "Delete":
		return h.handleDelete(ctx, activity)
	case "Update":
		return h.handleUpdate(ctx, activity)
	case "Undo":
		return h.handleUndo(ctx, activity)
	case "Like":
		return h.handleLike(ctx, activity)
	case "Announce":
		return h.handleAnnounce(ctx, activity)
	default:
		return fmt.Errorf("unsupported activity type: %s", activityType)
	}
}

// handleCreate processes Create activities
func (h *InboxHandler) handleCreate(ctx context.Context, activity map[string]interface{}) error {
	// TODO: Implement Create activity handling
	// This should:
	// 1. Validate the object being created
	// 2. Store the object in the database
	// 3. Notify relevant users
	return nil
}

// handleFollow processes Follow activities
func (h *InboxHandler) handleFollow(ctx context.Context, activity map[string]interface{}) error {
	// TODO: Implement Follow activity handling
	// This should:
	// 1. Verify the follow request is valid
	// 2. Update the followers database
	// 3. Send an Accept activity in response
	return nil
}

// Additional activity type handlers would be implemented similarly
