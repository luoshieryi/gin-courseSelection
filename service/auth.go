package service

import (
	"project/dao"
	"project/types"
	"project/util/hash"
	"strconv"
)

//Login 登陆
func Login(request types.LoginRequest) (string, string, types.ErrNo) {
	user := dao.GetMemberByUsername(request.Username)
	if user.ID == 0 || user.Deleted {
		return "", "0", types.WrongPassword
	}

	err := hash.NewHash().Check([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", strconv.FormatInt(user.ID, 10), types.WrongPassword
	}

	auth, err := dao.CreateSession(user.ID)
	if err != nil {
		return "", strconv.FormatInt(user.ID, 10), types.WrongPassword
	}

	return auth, strconv.FormatInt(user.ID, 10), types.OK
}

func Logout(auth string) types.ErrNo {
	session := dao.GetSessionByAuth(auth)
	if session.ID == 0 {
		return types.LoginRequired
	}

	err := dao.DeleteSessionByID(session.ID)
	if err != nil {
		return types.UnknownError
	}

	return types.OK
}

func Whoami(auth string) (types.TMember, types.ErrNo) {
	tMember := types.TMember{}

	session := dao.GetSessionByAuth(auth)
	if session.ID == 0 {
		return tMember, types.LoginRequired
	}

	member := dao.GetMemberByID(session.UserID)
	if member.Deleted {
		return tMember, types.UserHasDeleted
	}

	tMember = types.TMember{
		UserID:   strconv.FormatInt(member.ID, 10),
		Nickname: member.Nickname,
		Username: member.Username,
		UserType: types.UserType(member.UserType),
	}

	return tMember, types.OK
}
