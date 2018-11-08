package inf
import (
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
)



type AccessAuthProfile4UIRepo interface {
	//查询一条AccessAuthProfile4UI记录
	QueryAccessAuthProfile4UIOne(query map[string]interface{}) (*SELFENTITY.AccessAuthProfile4UI,error) 	

	//查询AccessAuthProfile4UI指定字段
	QueryAccessAuthProfile4UIDistinct(query map[string]interface{},field string,result interface{}) (error)

	//查询所有AccessAuthProfile4UI记录
	QueryAccessAuthProfile4UIAll(query map[string]interface{}) ([]*SELFENTITY.AccessAuthProfile4UI,error) 

	//查询AccessAuthProfile4UI分页排序记录
	QueryAccessAuthProfile4UIPage(query map[string]interface{}, limit int, sorts ...string) ([]*SELFENTITY.AccessAuthProfile4UI,error) 

	//查询AccessAuthProfile4UI数量
	QueryAccessAuthProfile4UICount(query map[string]interface{}) (int64,error) 

	//更新AccessAuthProfile4UI记录
	UpdateAccessAuthProfile4UI(selector , values map[string]interface{}) error

	//删除AccessAuthProfile4UI记录
	DeleteAccessAuthProfile4UI(selector map[string]interface{}) error

	//插入AccessAuthProfile4UI记录
	InsertAccessAuthProfile4UI(entities ...*SELFENTITY.AccessAuthProfile4UI) error

	//关闭repo
	Close() error

}