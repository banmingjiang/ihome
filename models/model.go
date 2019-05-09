package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Model Struct
type User struct {
	Id            int           `json:"user_id"` //id默认主键自增  若想设置其他：user_id int `orm:pk,auto`
	Name          string        `orm:"size(32)" json:"name"`
	Password_hash string        `orm:"size(128)" json :"password"` //varchar 100
	Mobile        string        `orm:"size(11)" json:"mobile"`
	Real_name     string        `orm:"size(32)" json:"real_name"`
	Id_card       string        `orm:"size(20)" json:"id_card"`
	Avatar_url    string        `orm:"size(256)" json:"avatar_url"`
	Houses        []*House      `orm:"reverse(many)" json:"houses"`
	Orders        []*OrderHouse `orm:"reverse(many)" json:"orders"` //数据库一对多
}

//房屋信息
type House struct {
	Id              int           `json:"house_id"`
	User            *User         `orm:"rel(fk)" json:"user_id"`
	Area            *Area         `orm:"rel(fk)" json:"area_id"`
	Title           string        `orm:"size(64)" json:"title"`
	Price           int           `orm:"default(0)" json:"price"`
	Address         string        `orm:"size(512)" orm:"default("")" json:"address"`
	Room_count      int           `orm:"default(1)" json:"room_count"`
	Acreage         int           `orm:"default(0)" json:"acreage"`              //房屋面积
	Unit            string        `orm:"size(32)" orm:"default("")" json:"unit"` //房屋单元号
	Capacity        int           `orm:"default(1)" json:"capacity"`
	Beds            string        `orm:"size(64)" orm:"default:("")" json:"beds"`            //床配置
	Deposit         int           `orm:"default(0)" json:"deposit"`                          //押金
	Min_days        int           `orm:"default(1) json:"min_days"`                          //最少入住天数
	Max_days        int           `orm:"default(0)" json:"max_days"`                         //最大入住 0不限制
	Order_count     int           `orm:"default(0)" json:"order_count"`                      //完成的订单数
	Index_image_url string        `orm:"size(256)" orm:"default("")" json:"index_image_url"` //房屋主图
	Factilities     []*Facility   `orm:"reverse(many)" json:"facilities"`                    //房屋设施
	Images          []*HouseImage `orm:"reverse(many)" json:"img_urls"`                      //房屋图片
	Orders          []*OrderHouse `orm:"reverse(many)" json:"orders"`                        //订单
	Ctime           time.Time     `orm:"auto_now_add;type:(datetime)" json:"ctime"`
}

//首页最高展示的房屋数量
var HOME_PAGE_MAX_HOUSES int = 5

//房屋列表页每页显示条目数
var HOUSE_LIST_PAGE_CAPACITY int = 2

//区域信息
type Area struct {
	Id     int      `json:"aid"`
	Name   string   `orm:"size(32)" json:"aname"`
	Houses []*House `orm:"reverse(many)" json:"houses"`
}

//设施信息
type Facility struct {
	Id     int      `json:"fid"`
	Name   string   `orm:"size(32)"`
	Housee []*House `orm:"rel(m2m)"` //有该设施的房
}

//房屋图片
type HouseImage struct {
	Id    int    `json:"house_image_id"`
	Url   string `orm:"size(256)" json:"url"`
	House *House `orm:"rel(fk)" json:"house_id"` //图片所属房

}

const (
	ORDER_STATUS_WAIT_ACCEPT  = "WAIT_ACCEPT"  //待接单
	ORDER_STATUS_WAIT_PAYMENT = "WAIT_PAYMENT" //待支付
	ORDER_STATUS_PAID         = "PAID"         //支付
	ORDER_STATUS_WAIT_COMMENT = "WAIT_COMMENT" //待评价
	ORDER_STATUS_COMPLETE     = "COMPLETE"     //完成
	ORDER_STATUS_CANCELED     = "CANCELED"     //取消
	ORDER_STATUS_REJECTED     = "REJECTED"     //拒单
)

type OrderHouse struct {
	Id          int       `json:"order_id"`
	User        *User     `orm:"rel(fk)" json:"user_id"`
	House       *House    `orm:"rel(fk)" json:house_id`
	Begin_date  time.Time `orm:"type(datetime)"`
	End_date    time.Time `orm:"type(datetime)"`
	Days        int       //预定天数
	House_price int       //单价
	Amount      int       //订单金额
	Status      string    `orm:"default(WAIT_ACCEPT)"`                     //订单状态
	Coment      string    `orm:"size(512)"`                                //评论
	Ctime       time.Time `orm:"auto_now_add;type(datetime)" json:"ctime"` //时间

}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/ihome?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User), new(OrderHouse), new(Facility), new(HouseImage), new(Area), new(House))

	// create table 第二个参数 true:数据表存在也添加覆盖  第三个：true不存在就创建  否则不执行
	orm.RunSyncdb("default", false, true)
}
