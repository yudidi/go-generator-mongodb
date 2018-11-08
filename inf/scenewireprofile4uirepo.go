package inf
import (
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
)



type SceneWireProfile4UIRepo interface {
	//查询一条SceneWireProfile4UI记录
	QuerySceneWireProfile4UIOne(query map[string]interface{}) (*SELFENTITY.SceneWireProfile4UI,error) 	

	//查询SceneWireProfile4UI指定字段
	QuerySceneWireProfile4UIDistinct(query map[string]interface{},field string,result interface{}) (error)

	//查询所有SceneWireProfile4UI记录
	QuerySceneWireProfile4UIAll(query map[string]interface{}) ([]*SELFENTITY.SceneWireProfile4UI,error) 

	//查询SceneWireProfile4UI分页排序记录
	QuerySceneWireProfile4UIPage(query map[string]interface{}, limit int, sorts ...string) ([]*SELFENTITY.SceneWireProfile4UI,error) 

	//查询SceneWireProfile4UI数量
	QuerySceneWireProfile4UICount(query map[string]interface{}) (int64,error) 

	//更新SceneWireProfile4UI记录
	UpdateSceneWireProfile4UI(selector , values map[string]interface{}) error

	//删除SceneWireProfile4UI记录
	DeleteSceneWireProfile4UI(selector map[string]interface{}) error

	//插入SceneWireProfile4UI记录
	InsertSceneWireProfile4UI(entities ...*SELFENTITY.SceneWireProfile4UI) error

	//关闭repo
	Close() error

}