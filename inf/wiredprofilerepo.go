package inf
import (
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
)



type WiredProfileRepo interface {
	//查询一条WiredProfile记录
	QueryWiredProfileOne(query map[string]interface{}) (*SELFENTITY.WiredProfile,error) 	

	//查询WiredProfile指定字段
	QueryWiredProfileDistinct(query map[string]interface{},field string,result interface{}) (error)

	//查询所有WiredProfile记录
	QueryWiredProfileAll(query map[string]interface{}) ([]*SELFENTITY.WiredProfile,error) 

	//查询WiredProfile分页排序记录
	QueryWiredProfilePage(query map[string]interface{}, limit int, sorts ...string) ([]*SELFENTITY.WiredProfile,error) 

	//查询WiredProfile数量
	QueryWiredProfileCount(query map[string]interface{}) (int64,error) 

	//更新WiredProfile记录
	UpdateWiredProfile(selector , values map[string]interface{}) error

	//删除WiredProfile记录
	DeleteWiredProfile(selector map[string]interface{}) error

	//插入WiredProfile记录
	InsertWiredProfile(entities ...*SELFENTITY.WiredProfile) error

	//关闭repo
	Close() error

}