package entities
import (
	"fmt"
)



// Save .
func Insert(u * UserInfo)  {
	// _,err := engine.Exec(userInfoInsertSql,u.Username,u.Departname)
	// un := UserInfo{Username: "ltj", Departname: "609"}
	fmt.Println("insert: ",*u)
	_, err := engine.Insert(u)
	checkErr(err)
	
	

}

// FindAll .
func  FindAllUser() []UserInfo {

	ulist := make([]UserInfo, 0, 0)
	err := engine.Find(&ulist)
	checkErr(err)
	return ulist
}

// FindByID . 
func FindUserByID(id int) *UserInfo {
	user := new(UserInfo)
	_,err := engine.Id(id).Get(user)
	checkErr(err)
	return user
}
