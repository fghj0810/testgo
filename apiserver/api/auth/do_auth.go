package api_auth

import (
	"../../common"
	"../../db"
	"encoding/json"
	"net/http"
	"strings"
)

type userInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleFunction(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	ret := make(map[string]interface{}, 5)
	var info userInfo = userInfo{"", ""}
	if r.Method != "POST" {
		ret["code"] = common.Code_InvalidAPI
		encoder.Encode(ret)
		return
	}
	if r.ContentLength > 0 {
		j := json.NewDecoder(r.Body)
		err := j.Decode(&info)
		if err != nil {
			ret["code"] = common.Code_InvalidAPI
			encoder.Encode(ret)
			return
		} else {
			info.Username = strings.TrimSpace(info.Username)
			info.Password = strings.TrimSpace(info.Password)
		}
	}
	if info.Username != "" && info.Password != "" {
		// 用户名密码登录
		if uid, ok := db.Login(info.Username, info.Password); ok {
			sid, err := db.SessionNew(uid)
			if err != nil {
				ret["code"] = common.Code_Auth_NewSessionError
			} else {
				ret["code"] = common.Code_Success
				ret["message"] = "账号登录成功"
				// 设置cookie
				cookie := http.Cookie{Name: "SID", Value: sid, Path: "/", MaxAge: 24 * 60 * 60}
				http.SetCookie(w, &cookie)
			}
		} else {
			ret["code"] = common.Code_Auth_UsernameOrPasswordNotMatch
		}
	} else {
		// sid 登录
		if cookie, err := r.Cookie("SID"); err != nil {
			ret["code"] = common.Code_Auth_SessionError
		} else {
			if db.ValidUidAndSid(db.SessionDecode(cookie.Value), cookie.Value) {
				ret["code"] = common.Code_Success
				ret["message"] = "session登录成功"
				// 设置cookie
				cookie := http.Cookie{Name: "SID", Value: cookie.Value, Path: "/", MaxAge: 24 * 60 * 60}
				http.SetCookie(w, &cookie)
			} else {
				ret["code"] = common.Code_Auth_SessionError
			}
		}
	}

	encoder.Encode(ret)
}
