
package inf
import (
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
)



type AccessAuthProfileRepo interface {
	//查询一条AccessAuthProfile记录
	QueryAccessAuthProfileOne(query map[string]interface{}) (*SELFENTITY.AccessAuthProfile,error) 	

	//查询AccessAuthProfile指定字段
	QueryAccessAuthProfileField(query map[string]interface{},field string) ([]interface{},error)

	//查询所有AccessAuthProfile记录
	QueryAccessAuthProfileAll(query map[string]interface{}) (*[]*SELFENTITY.AccessAuthProfile,error) 

	//查询AccessAuthProfile分页排序记录
	QueryAccessAuthProfilePage(query map[string]interface{}, limit int, sorts ...string) (*[]*SELFENTITY.AccessAuthProfile,error) 

	//查询AccessAuthProfile数量
	QueryAccessAuthProfileCount(query map[string]interface{}) (int64,error) 

	//更新AccessAuthProfile记录
	UpdateAccessAuthProfile(selector , values map[string]interface{}) error

	//删除AccessAuthProfile记录
	DeleteAccessAuthProfile(selector map[string]interface{}) error

	//插入AccessAuthProfile记录
	InsertAccessAuthProfile(entities ...*SELFENTITY.AccessAuthProfile) error

	//关闭repo
	Close() error

}