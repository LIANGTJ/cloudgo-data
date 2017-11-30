package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo)  {
	 Insert(u)
	
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	 return FindAllUser()
	
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	return FindUserByID(id)
}
