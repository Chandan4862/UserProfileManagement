package api

type User struct{
	
	Id   string `json:"id" bson:"id,omitempty"`
	First_name string `json:"fname"`
	Last_name string `json:"lname"`
	Email string `json:"email"`
	Mobile_number string `json:"mno"`
	//Address []string `json:"add"`

}

var userList []User