package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"errors"
)

type User struct {
	ID int64 `gorm:"column:id" json:"id" form:"id"`
	Company string `gorm:"column:company" json:"company" form:"company"`
	Type int64 `gorm:"column:type" json:"type" form:"type"`
	Title string `gorm:"column:title" json:"title" form:"title"`
	Password string `gorm:"column:passwd" json:"passwd" form:"passwd"`
	Name string `gorm:"column:name" json:"name" form:"name"`
	NickName *string `gorm:"column:nick_name" json:"nick_name" form:"nick_name"`
	Email string `gorm:"column:email" json:"email" form:"email"`
	Phone string `gorm:"column:phone" json:"phone" form:"phone"`
	State int64 `gorm:"column:state" json:"state" form:"state"`
	CreateTime time.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time"`
	ModifyTime time.Time `gorm:"column:modify_time;default:null" json:"modify_time" form:"modify_time"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:null" json:"deleted_at" form:"deleted_at"`
}

func (u User) TableName() string {
	return "user"
}
//创建新纪录钩子
func (u *User) BeforeCreate(tx *gorm.DB) (err error)  {
	u.CreateTime = time.Now()
	u.ModifyTime = time.Now()
	return
}
//更新记录钩子
func (u *User) BeforeUpdate(tx *gorm.DB) (err error)  {
	u.ModifyTime = time.Now()
	return
}

func NewUser(email ,password string) *User {
	return &User{
		Email: email,
		Password: password,
	}
}
//创建新的记录
func NewRecords()  {
	user := User{
		Company: "sina",
		Type: 2,
		Email:"944677073@qq.com",
		NickName: new(string),
		Password:"129670",
		State: 2,
	}
	result := DbHandler.Select("Email,Password,State").Create(&user)
	fmt.Println(result.RowsAffected)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
}
//使用批量添加
func BashNewRcords()  {
	user := []User{
		{
			Company: "baidu",
			Password: "123003",
		},{
			Company: "baidu",
			Password: "123003",
		},
	}
	DbHandler.Create(&user)
	for _,u := range user{
		fmt.Println("插入id",u.ID)
	}
}
//使用map单个或者批量添加
func BashWithMap()  {
	DbHandler.Model(&User{}).Create(map[string]interface{}{
		"Company": "baidu_bash",
		"Password": "123003",
	})
	DbHandler.Model(&User{}).Create([]map[string]interface{}{
		{
			"Company": "baidu_bash_map",
			"Password": "123003",
		},
		{
			"Company": "baidu_bash_map",
			"Password": "123003",
		},
	})
}

func GetRecords()  {
	var users User
	//第一条
	//DbHandler.First(&user,3)
	//可能会导致注入风险
	//DbHandler.First(&user,"3")
	//获取多条
	//DbHandler.Find(&users,[]int{2,5})
	//获取全部,中间加上查询条件
	//name <> ?
	//name IN ? //后面是个切片
	//name like ? //后面加上条件
	//date >= ? //后面给上时间
	//record := DbHandler.Where("id=?",4).Find(&users)
	//使用struct查询多个条件同时满足，比上面的好用
	//包括map[string]interface{} 跟上面的查询条件一致 struct跟上面的结构体一致 slice主键数组切片
	//record := DbHandler.Where(map[string]interface{}{"ID": 3}).Find(&users)
	//不过注意下使用默认住：0，''，false等使用默认的字段零值时，会自动忽略，最好使用上面的map生成相应的条件
	//record := DbHandler.Where(&User{ID: 3}).Find(&users)
	//注意字段类型为int64
	record := DbHandler.Select([]string{"ID","Password","Email","Name"}).Find(&users)
	//获取一条
	//DbHandler.Take(&user)
	//获取最后一条
	//record	:= DbHandler.Last(&user)
	fmt.Printf("获取的记录数：%d \n",record.RowsAffected)
	fmt.Println(errors.Is(record.Error,gorm.ErrRecordNotFound))
	fmt.Printf("%#v \n",users)
}
