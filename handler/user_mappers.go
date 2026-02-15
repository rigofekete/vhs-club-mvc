package handler

import "github.com/rigofekete/vhs-club-mvc/model"

func (r *CreateUserRequest) ToModel() *model.User {
	return &model.User{
		Username: r.Username,
		Email:    r.Email,
	}
}

func UserSingleResponse(user *model.User) *UserResponse {
	return &UserResponse{
		PublicID: user.PublicID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func UserListResponse(user []*model.User) []*UserResponse {
	userList := make([]*UserResponse, len(user))
	for i, user := range user {
		userList[i] = UserSingleResponse(user)
	}
	return userList
}
