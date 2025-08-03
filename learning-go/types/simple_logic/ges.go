package 

import "fmt"

func LogOutput(message string) {
	fmt.Println(msg)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userdID string) (string, bool) {
	name, ok = sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore {
		userData: map[string]string {
			"1": "Mike",
			"2": "Jane",
			"3", "Paul",
		}
	}
}

type DataStore Interface {
	UserNameForID(string) (string, bool)
}

type Logger InterFace {
	Log(string)
}

type LoggerAdapter func(string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}


type SimpleLogic Struct {
	l Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("In sayhello for: " + userId)
	name, ok := sl.ds.UserNameforId(userId)
	if !ok {
		return "", errors.New("Unknown user")
	}
	return "Hello, " + name
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("In saygoodbye for: " + userId)
	name, ok := sl.ds.UserNameforId(userId)
	if !ok {
		return "", errors.New("Unknown user")
	}
	return "GoodBye,  " + name
}

