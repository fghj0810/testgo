package db

import (
	"../rand"
	"crypto/rc4"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const SID_VERSION string = "V1"

func Login(userName string, password string) (uint64, bool) {
	if value, ok := userMap[userName]; ok {
		if value.password == password {
			return value.uid, true
		}
	}
	return 0, false
}

func SessionNew(uid uint64) (string, error) {
	salt := rand.String(16)
	head := salt[0:8]
	cipher, _ := rc4.NewCipher([]byte(head))
	src := []byte(fmt.Sprintf("%v#%v", uid, salt[8:]))
	dst := make([]byte, len(src))
	cipher.XORKeyStream(dst, src)
	preSid := base64.StdEncoding.EncodeToString(dst)
	sid := SID_VERSION + preSid[0:4] + head + preSid[4:]

	sessionDelete(uid)

	err := sessionSet(uid, sid)
	return sid, err
}

func SessionDecode(sid string) uint64 {
	if sid == "" {
		return 0
	}
	if sid[0:2] != SID_VERSION {
		return 0
	}
	uid := uint64(0)

	session := sid[2:]
	salt := session[4:12]
	body := session[0:4] + session[12:]
	dst, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return 0
	}
	src := make([]byte, len(dst))
	cipher, _ := rc4.NewCipher([]byte(salt))
	cipher.XORKeyStream(src, dst)

	// uid#salt#time#salt#client_type
	l := strings.Split(string(src), "#")
	if len(l) != 2 {
		return 0
	}

	uid, err = strconv.ParseUint(l[0], 10, 64)
	if err != nil {
		return 0
	}
	return uid
}

func ValidUidAndSid(uid uint64, sid string) bool {
	if sid2, ok := utsMap[uid]; ok {
		return sid == sid2
	}
	return false
}

func sessionSet(uid uint64, sid string) error {
	if _, ok := stuMap[sid]; ok {
		return errors.New("session conflict")
	}
	utsMap[uid] = sid
	stuMap[sid] = uid
	return nil
}

func sessionDelete(uid uint64) {
	if sid, ok := utsMap[uid]; ok {
		delete(stuMap, sid)
		delete(utsMap, uid)
	}
}
