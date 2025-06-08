package iam

import "github.com/aws/aws-sdk-go-v2/service/iam/types"

type IamUser struct {
	UserName         string
	UserId           string
	Arn              string
	CreateDate       string
	PasswordLastUsed string
}

func NewIamUser(user types.User) IamUser {
	u := IamUser{
		UserName:         *user.UserName,
		UserId:           *user.UserId,
		Arn:              *user.Arn,
		CreateDate:       user.CreateDate.String(),
		PasswordLastUsed: user.PasswordLastUsed.String(),
	}
	return u
}
