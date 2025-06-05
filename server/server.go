package server

import (
	"encoding/json"
	"io"
	"net/http"

	pb "github.com/proto-http/api"
	"google.golang.org/protobuf/proto"
)

// UserServer implements the user service
type UserServer struct {
	users map[string]*pb.User
}

// NewUserServer creates a new UserServer
func NewUserServer() *UserServer {
	// Initialize with some dummy data
	users := map[string]*pb.User{
		"1": {
			Id:    "1",
			Name:  "John Doe",
			Email: "john@example.com",
		},
	}
	return &UserServer{users: users}
}

// handleGetUserRequest handles GET requests for user data
func (s *UserServer) handleGetUserRequest(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	user, exists := s.users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := &pb.UserResponse{
		User: user,
	}

	s.writeProtobufResponse(w, response)
}

// handlePostUserRequest handles POST requests for creating users
func (s *UserServer) handlePostUserRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON request
	var jsonUser struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.Unmarshal(body, &jsonUser); err != nil {
		http.Error(w, "Failed to parse JSON request", http.StatusBadRequest)
		return
	}

	// Create protobuf user
	user := &pb.User{
		Id:    jsonUser.ID,
		Name:  jsonUser.Name,
		Email: jsonUser.Email,
	}

	// Store user
	s.users[user.Id] = user

	// Create response
	response := &pb.UserResponse{
		User: user,
	}

	s.writeProtobufResponse(w, response)
}

// writeProtobufResponse marshals and writes a protobuf response
func (s *UserServer) writeProtobufResponse(w http.ResponseWriter, response *pb.UserResponse) {
	data, err := proto.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(data)
}

// HandleUser is the main handler that routes requests based on HTTP method
func (s *UserServer) HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetUserRequest(w, r)
	case http.MethodPost:
		s.handlePostUserRequest(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
