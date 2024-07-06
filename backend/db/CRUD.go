package db

type InterfaceOfGoodsCRUD interface {
	CreateByObject(g *Goods) error
	FindById(id uint) (*Goods, error)
	FindAll() ([]Goods, error)
	UpdateByObject(g *Goods) error
	FindAllOrdered() ([]Goods, error)
	DeleteById(GoodsId string) error
	FuzzyGetUserByName(name string) ([]User, error)
	GetUserByName(name string) (*User, error)
	GetAllUser() ([]User, error)
	GetAllUserOrdered() ([]User, error)
	DeleteUserbyName(name string) error
}

type InterfaceOfCommentCRUD interface {
	CreateByObject(c *Comment) error
	UpdateByObject(c *Comment) error
	DeleteById(id uint) error
	FindAllByGoodsId(GoodsId uint) ([]Comment, error)
}

type InterfaceOfUserCRUD interface {
	CreateByObject(u *User) error
	GetUserByName(name string) ([]User, error)
	UpdateByObject(u *User) error
}
