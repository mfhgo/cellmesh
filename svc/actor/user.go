package actor

type Account struct {
	RoleList []int64 // 已经拥有的角色ID列表
}

// 角色基本信息
type Role struct {
	RoleID   int64  // 角色ID
	Account  string // 账号, 方便快速获取
	NickName string //  玩家昵称
	ServerID int32  // 玩家所在服
}

type Item struct {
	ID       int64
	ItemID   int32
	Count    int32
	ExpireTS int64 // 过期时间 0: 永久 >0:到期时间戳
}

// 背包
type Bag struct {
	ItemByID map[int64]*Item
}