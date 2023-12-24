package auth

type ProfileResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type PhotoResponse struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
	UserId   int    `json:"userId"`
}

type CommentResponse struct {
	UserId  int    `json:"userId"`
	PhotoId int    `json:"photoId"`
	Message string `json:"message"`
}
