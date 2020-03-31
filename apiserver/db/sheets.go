package db

type UserNameAndPassword struct {
	username string
	password string
	uid      uint64
}

type SidAndUid struct {
	sid string
	uid uint64
}

var userMap map[string]UserNameAndPassword = make(map[string]UserNameAndPassword, 10)
var utsMap map[uint64]string = make(map[uint64]string, 10)
var stuMap map[string]uint64 = make(map[string]uint64, 10)

func init() {
	userMap["test1"] = UserNameAndPassword{"test1", "pwd", 1}
	userMap["test2"] = UserNameAndPassword{"test2", "pwd2", 2}
}
