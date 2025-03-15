package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"openfirm/internal/activitypub"
)

type ActorHandler struct {
	activityPubService *activitypub.Service
}

func NewActorHandler(activityPubService *activitypub.Service) *ActorHandler {
	return &ActorHandler{
		activityPubService: activityPubService,
	}
}

// Get returns the ActivityPub actor representation of a user
func (h *ActorHandler) Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	// Check if client accepts ActivityPub format
	accept := r.Header.Get("Accept")
	if !strings.Contains(accept, "application/activity+json") &&
		!strings.Contains(accept, "application/ld+json") {
		http.Error(w, "Not Acceptable", http.StatusNotAcceptable)
		return
	}

	actor, err := h.activityPubService.GetActor(r.Context(), username)
	if err != nil {
		http.Error(w, "Actor not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	json.NewEncoder(w).Encode(actor)
}

// Webfinger handles .well-known/webfinger requests
func (h *ActorHandler) Webfinger(w http.ResponseWriter, r *http.Request) {
	resource := r.URL.Query().Get("resource")
	if resource == "" {
		http.Error(w, "Resource parameter required", http.StatusBadRequest)
		return
	}

	// Extract username from acct: URI
	parts := strings.Split(resource, ":")
	if len(parts) != 2 || !strings.HasPrefix(resource, "acct:") {
		http.Error(w, "Invalid resource format", http.StatusBadRequest)
		return
	}

	username := strings.Split(parts[1], "@")[0]

	response, err := h.activityPubService.WebFinger(r.Context(), username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/jrd+json")
	json.NewEncoder(w).Encode(response)
}

// NodeInfo handles .well-known/nodeinfo requests
func (h *ActorHandler) NodeInfo(w http.ResponseWriter, r *http.Request) {
	nodeInfo, err := h.activityPubService.NodeInfo(r.Context())
	if err != nil {
		http.Error(w, "Failed to get node info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nodeInfo)
}

// NodeInfoSchema handles .well-known/nodeinfo/2.0 requests
func (h *ActorHandler) NodeInfoSchema(w http.ResponseWriter, r *http.Request) {
	schema := map[string]interface{}{
		"links": []map[string]string{
			{
				"rel":  "http://nodeinfo.diaspora.software/ns/schema/2.0",
				"href": "https://" + r.Host + "/nodeinfo/2.0",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schema)
}

// HostMeta handles .well-known/host-meta requests
func (h *ActorHandler) HostMeta(w http.ResponseWriter, r *http.Request) {
	hostMeta := map[string]interface{}{
		"links": []map[string]string{
			{
				"rel":      "lrdd",
				"template": "https://" + r.Host + "/.well-known/webfinger?resource={uri}",
			},
		},
	}

	if strings.Contains(r.Header.Get("Accept"), "application/json") {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hostMeta)
		return
	}

	// Default to XRD format
	w.Header().Set("Content-Type", "application/xrd+xml")
	w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
  <Link rel="lrdd" template="https://` + r.Host + `/.well-known/webfinger?resource={uri}"/>
</XRD>`))
}

// Following returns a list of accounts the user follows
func (h *ActorHandler) Following(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	page := 1 // TODO: Add pagination support

	following, err := h.activityPubService.GetFollowing(r.Context(), username, page)
	if err != nil {
		http.Error(w, "Failed to get following list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	json.NewEncoder(w).Encode(following)
}

// Followers returns a list of accounts that follow the user
func (h *ActorHandler) Followers(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	page := 1 // TODO: Add pagination support

	followers, err := h.activityPubService.GetFollowers(r.Context(), username, page)
	if err != nil {
		http.Error(w, "Failed to get followers list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	json.NewEncoder(w).Encode(followers)
}

// Featured returns a collection of featured posts
func (h *ActorHandler) Featured(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	featured, err := h.activityPubService.GetFeatured(r.Context(), username)
	if err != nil {
		http.Error(w, "Failed to get featured posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	json.NewEncoder(w).Encode(featured)
}
