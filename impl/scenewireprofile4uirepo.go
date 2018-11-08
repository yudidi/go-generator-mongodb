package impl
import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	MONGO "han-networks.com/csp/common_grpc/mongo"
	CONFIG "han-networks.com/csp/config_grpc/config"
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
	REPO "gorepomaker/inf"
)





//获取SceneWireProfile4UIMongoRepo对象
func NewSceneWireProfile4UIMongoRepo() REPO.SceneWireProfile4UIRepo {
	return &SceneWireProfile4UIMongoRepo{session: MONGO.GetSession()}
}

//持久层mongo实现
type SceneWireProfile4UIMongoRepo struct {
	session *mgo.Session
}



//查询一条SceneWireProfile4UI记录
func (this *SceneWireProfile4UIMongoRepo)QuerySceneWireProfile4UIOne(query map[string]interface{}) (*SELFENTITY.SceneWireProfile4UI,error) {
	entity := SELFENTITY.SceneWireProfile4UI{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.SceneWireProfile4UICol).Find(query).One(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}	

//查询SceneWireProfile4UI指定字段
func (this *SceneWireProfile4UIMongoRepo)QuerySceneWireProfile4UIDistinct(query map[string]interface{},field string,result interface{}) (error) {
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.SceneWireProfile4UICol).Find(query).Distinct(field,result)
	if err != nil {
		return err
	}
	return nil
}	

//查询所有SceneWireProfile4UI记录
func (this *SceneWireProfile4UIMongoRepo)QuerySceneWireProfile4UIAll(query map[string]interface{}) ([]*SELFENTITY.SceneWireProfile4UI,error) {
	entities := []*SELFENTITY.SceneWireProfile4UI{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.SceneWireProfile4UICol).Find(query).All(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}	

//查询SceneWireProfile4UI分页排序记录
func (this *SceneWireProfile4UIMongoRepo)QuerySceneWireProfile4UIPage(query map[string]interface{}, limit int, sorts ...string) ([]*SELFENTITY.SceneWireProfile4UI,error) {
	q := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.SceneWireProfile4UICol).Find(query)
	if sorts != nil && len(sorts) > 0 {
		for _, s := range sorts {
			q.Sort(s)
		}
	}
	entities := []*SELFENTITY.SceneWireProfile4UI{}
	q.Limit(limit).All(&entities)

	return entities, nil

}	

//查询SceneWireProfile4UI数量
func (this *SceneWireProfile4UIMongoRepo)QuerySceneWireProfile4UICount(query map[string]interface{}) (int64,error) {
	count,err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.SceneWireProfile4UICol).Find(query).Count()
	if err != nil {
		return -1,err
	}
	return int64(count), nil
}	

//更新SceneWireProfile4UI记录
func (this *SceneWireProfile4UIMongoRepo) UpdateSceneWireProfile4UI(selector , values map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.SceneWireProfile4UICol).UpdateAll(selector, bson.M{"$set": values})
	if err != nil {
		return err
	}
	return nil

}	

//删除SceneWireProfile4UI记录
func (this *SceneWireProfile4UIMongoRepo) DeleteSceneWireProfile4UI(selector map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.SceneWireProfile4UICol).RemoveAll(selector)
	if err != nil {
		return err
	}
	return nil

}	

//插入SceneWireProfile4UI记录
func (this *SceneWireProfile4UIMongoRepo) InsertSceneWireProfile4UI(entities ...*SELFENTITY.SceneWireProfile4UI) error {
	entitiesInterface:= []interface{}{}
	for _,entity:=range entities{
		entitiesInterface=append(entitiesInterface,entity)
	}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.SceneWireProfile4UICol).Insert(entitiesInterface...)
	if err != nil {
		return err
	}
	return nil
}


//关闭repo
func (this *SceneWireProfile4UIMongoRepo) Close() error {
	this.session.Close()
	return nil
}
