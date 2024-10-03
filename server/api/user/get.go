package user

import (
	"errors"
	. "savvy_life/common"
	. "savvy_life/proto/api/v1"
	. "savvy_life/proto/module"
	"xorm.io/xorm"
)

func GetUser(req *GetUserReq) (resp *GetUserResp, err error) {
	Logger.Debug("Func: GetUser start.")
	resp = new(GetUserResp)
	resp.Code = SuccessCode
	resp.User = new(User)

	session := Engine.NewSession()
	defer func(session *xorm.Session) {
		var err error
		err = session.Close()
		if err != nil {
			Logger.Error("session.Close failed. error:", err)
		}
	}(session)

	user := new(User)
	session = session.Where("uuid=?", req.Uuid)
	if has, err := session.Get(user); err != nil {
		Logger.Error("session.Get user failed. error:", err)
		return resp, nil
	} else if !has {
		Logger.Error("no user record.")
		return resp, errors.New("no user record")
	}
	resp.User = user
	Logger.Debug("Func: GetUser end.")
	return resp, nil
}
