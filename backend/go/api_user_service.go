/*
 * User API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UserApiService is a service that implents the logic for the UserApiServicer
// This service should implement the business logic for every endpoint for the UserApi API.
// Include any external packages or services that will be required by this service.
type UserApiService struct {
}

// NewUserApiService creates a default api service
func NewUserApiService() UserApiServicer {
	return &UserApiService{}
}

// CreateUser - Creates a User
func (s *UserApiService) CreateUser(user User) (interface{}, error) {
	// TODO - update CreateUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return user, nil
}

// DeleteUser - Deletes a User
func (s *UserApiService) DeleteUser(id string) (interface{}, error) {
	// TODO - update DeleteUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return User{Id: id, Nickname: "Deleted-Dummy"}, nil
}

// GetUser - Get a User
func (s *UserApiService) GetUser(id string) (interface{}, error) {
	// TODO - update GetUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return User{Id: id, Nickname: "Get-Dummy"}, nil
}

// UpdateUser - Updates a User
func (s *UserApiService) UpdateUser(user User) (interface{}, error) {
	// TODO - update UpdateUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return user, nil
}
