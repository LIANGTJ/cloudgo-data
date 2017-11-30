package entities
// Save .
func Save(u * UserInfo)  {
	_, err := engine.Insert(u)
	checkErr(err)

}

// FindAll .
func  FindAll() []UserInfo {

	ulist := make([]UserInfo, 0, 0)
	err := engine.Find(&ulist)
	checkErr(err)
	return ulist
}

// FindByID . 
func FindByID(id int) *UserInfo {
	user := new(UserInfo)
	_,err := engine.Id(id).Get(user)
	checkErr(err)
	return user
}
