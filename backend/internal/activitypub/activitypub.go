package activitypub

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
	"github.com/jackc/pgx/v5/pgxpool"
	"openfirm/internal/models"
)

type Service struct {
	db         *pgxpool.Pool
	domain     string
	userSvc    *models.UserService
	postSvc    *models.PostService
}

func NewService(db *pgxpool.Pool, domain string) *Service {
	return &Service{
		db:      db,
		domain:  domain,
		userSvc: models.NewUserService(db),
		postSvc: models.NewPostService(db),
	}
}

// Actor represents an ActivityPub actor (user)
type Actor struct {
	Context           []string `json:"@context"`
	ID                string   `json:"id"`
	Type              string   `json:"type"`
	PreferredUsername string   `json:"preferredUsername"`
	Name              string   `json:"name,omitempty"`
	Summary           string   `json:"summary,omitempty"`
	Icon              *Image   `json:"icon,omitempty"`
	Inbox            string   `json:"inbox"`
	Outbox           string   `json:"outbox"`
	Following        string   `json:"following"`
	Followers        string   `json:"followers"`
	PublicKey        PublicKey `json:"publicKey,omitempty"`
}

type Image struct {
	Type      string `json:"type"`
	MediaType string `json:"mediaType"`
	URL       string `json:"url"`
}

type PublicKey struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

// GetActor returns the ActivityPub actor representation of a user
func (s *Service) GetActor(ctx context.Context, username string) (*Actor, error) {
	user, err := s.userSvc.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	actorURL := fmt.Sprintf("https://%s/users/%s", s.domain, username)
	actor := &Actor{
		Context: []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		ID:                actorURL,
		Type:             "Person",
		PreferredUsername: user.Username,
		Name:             user.DisplayName,
		Summary:          user.Bio,
		Inbox:            fmt.Sprintf("%s/inbox", actorURL),
		Outbox:           fmt.Sprintf("%s/outbox", actorURL),
		Following:        fmt.Sprintf("%s/following", actorURL),
		Followers:        fmt.Sprintf("%s/followers", actorURL),
	}

	if user.AvatarURL != "" {
		actor.Icon = &Image{
			Type:      "Image",
			MediaType: "image/jpeg", // Adjust based on actual image type
			URL:       user.AvatarURL,
		}
	}

	return actor, nil
}

// CreateNote creates an ActivityPub Note object from a post
func (s *Service) CreateNote(post *models.Post) (vocab.Type, error) {
	note := streams.NewActivityStreamsNote()
	
	// Set ID
	id := streams.NewActivityStreamsId()
	id.Set(fmt.Sprintf("https://%s/posts/%d", s.domain, post.ID))
	note.SetActivityStreamsId(id)

	// Set content
	content := streams.NewActivityStreamsContent()
	content.AppendXMLSchemaString(post.Content)
	note.SetActivityStreamsContent(content)

	// Set published time
	published := streams.NewActivityStreamsPublished()
	published.Set(post.CreatedAt)
	note.SetActivityStreamsPublished(published)

	// Set attribution
	attribution := streams.NewActivityStreamsAttributedTo()
	actorIRI := fmt.Sprintf("https://%s/users/%d", s.domain, post.UserID)
	attribution.AppendIRI(actorIRI)
	note.SetActivityStreamsAttributedTo(attribution)

	return note, nil
}

// HandleInbox processes incoming ActivityPub activities
func (s *Service) HandleInbox(ctx context.Context, body []byte) error {
	var activity map[string]interface{}
	if err := json.Unmarshal(body, &activity); err != nil {
		return fmt.Errorf("failed to unmarshal activity: %v", err)
	}

	switch activity["type"] {
	case "Follow":
		return s.handleFollow(ctx, activity)
	case "Undo":
		if nested, ok := activity["object"].(map[string]interface{}); ok {
			if nested["type"] == "Follow" {
				return s.handleUnfollow(ctx, activity)
			}
		}
	case "Create":
		if nested, ok := activity["object"].(map[string]interface{}); ok {
			if nested["type"] == "Note" {
				return s.handleCreate(ctx, activity)
			}
		}
	}

	return nil
}

// handleFollow processes Follow activities
func (s *Service) handleFollow(ctx context.Context, activity map[string]interface{}) error {
	// Implementation for handling Follow activities
	// This would typically involve:
	// 1. Validating the activity
	// 2. Creating a follower relationship in the database
	// 3. Sending an Accept activity in response
	return nil
}

// handleUnfollow processes Unfollow activities
func (s *Service) handleUnfollow(ctx context.Context, activity map[string]interface{}) error {
	// Implementation for handling Unfollow activities
	// This would typically involve:
	// 1. Validating the activity
	// 2. Removing the follower relationship from the database
	return nil
}

// handleCreate processes Create activities
func (s *Service) handleCreate(ctx context.Context, activity map[string]interface{}) error {
	// Implementation for handling Create activities
	// This would typically involve:
	// 1. Validating the activity
	// 2. Creating a new post in the database
	// 3. Federating the post to followers if necessary
	return nil
}

// GetOutbox returns a user's outbox (their posts)
func (s *Service) GetOutbox(ctx context.Context, username string, page int) (vocab.Type, error) {
	user, err := s.userSvc.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	collection := streams.NewActivityStreamsOrderedCollection()
	
	// Set ID
	id := streams.NewActivityStreamsId()
	id.Set(fmt.Sprintf("https://%s/users/%s/outbox", s.domain, username))
	collection.SetActivityStreamsId(id)

	// Get posts
	posts, err := s.postSvc.ListUserPosts(ctx, user.ID, (page-1)*20, 20)
	if err != nil {
		return nil, err
	}

	// Convert posts to activities
	items := streams.NewActivityStreamsOrderedItems()
	for _, post := range posts {
		note, err := s.CreateNote(post)
		if err != nil {
			return nil, err
		}
		items.Append(note)
	}
	collection.SetActivityStreamsOrderedItems(items)

	return collection, nil
}

// WebFinger handles .well-known/webfinger requests
func (s *Service) WebFinger(ctx context.Context, resource string) (map[string]interface{}, error) {
	// Implementation for WebFinger protocol
	// This would return account information in WebFinger format
	return map[string]interface{}{
		"subject": resource,
		"links": []map[string]interface{}{
			{
				"rel":  "self",
				"type": "application/activity+json",
				"href": fmt.Sprintf("https://%s/users/%s", s.domain, resource),
			},
		},
	}, nil
}

// NodeInfo handles .well-known/nodeinfo requests
func (s *Service) NodeInfo(ctx context.Context) (map[string]interface{}, error) {
	// Implementation for NodeInfo protocol
	// This would return instance information in NodeInfo format
	return map[string]interface{}{
		"version": "2.0",
		"software": map[string]interface{}{
			"name":    "openfirm",
			"version": "1.0.0",
		},
		"protocols": []string{"activitypub"},
		"services": map[string]interface{}{
			"inbound":  []string{},
			"outbound": []string{},
		},
		"usage": map[string]interface{}{
			"users": map[string]interface{}{
				"total": 0,
			},
		},
		"openRegistrations": true,
	}, nil
}
