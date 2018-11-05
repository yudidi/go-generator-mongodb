
package impl
import (
	"gopkg.in/mgo.v2"
	MONGO "han-networks.com/csp/common_grpc/mongo"
	CONFIG "han-networks.com/csp/config_grpc/config"
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
	REPO "mongorepo/inf"
)





//获取AccessAuthProfile4UIMongoRepo对象
func NewAccessAuthProfile4UIMongoRepo() REPO.AccessAuthProfile4UIRepo {
	return &AccessAuthProfile4UIMongoRepo{session: MONGO.GetSession()}
}

//持久层mongo实现
type AccessAuthProfile4UIMongoRepo struct {
	session *mgo.Session
}



//查询一条AccessAuthProfile4UI记录
func (this *AccessAuthProfile4UIMongoRepo)QueryAccessAuthProfile4UIOne(query map[string]interface{}) (*SELFENTITY.AccessAuthProfile4UI,error) {
	entity := SELFENTITY.AccessAuthProfile4UI{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfile4UICol).Find(query).One(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}	

//查询AccessAuthProfile4UI指定字段
func (this *AccessAuthProfile4UIMongoRepo)QueryAccessAuthProfile4UIField(query map[string]interface{},field string) ([]interface{},error) {
	entityMap := []map[string]interface{}{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfile4UICol).Find(query).All(&entityMap)
	if err != nil {
		return nil, err
	}
	fields:=[]interface{}{}
	for _,entity:=range entityMap{
		for k,v:=range entity{
			if k == field{
				fields =append(fields,v)
			}
		}
	}
	return fields, nil
}	

//查询所有AccessAuthProfile4UI记录
func (this *AccessAuthProfile4UIMongoRepo)QueryAccessAuthProfile4UIAll(query map[string]interface{}) (*[]*SELFENTITY.AccessAuthProfile4UI,error) {
	entities := []*SELFENTITY.AccessAuthProfile4UI{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfile4UICol).Find(query).All(entities)
	if err != nil {
		return nil, err
	}
	return &entities, nil
}	

//查询AccessAuthProfile4UI分页排序记录
func (this *AccessAuthProfile4UIMongoRepo)QueryAccessAuthProfile4UIPage(query map[string]interface{}, limit int, sorts ...string) (*[]*SELFENTITY.AccessAuthProfile4UI,error) {
	q := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfile4UICol).Find(query)
	if sorts != nil && len(sorts) > 0 {
		for _, s := range sorts {
			q.Sort(s)
		}
	}
	entities := []*SELFENTITY.AccessAuthProfile4UI{}
	q.Limit(limit).All(&entities)

	return &entities, nil

}	

//查询AccessAuthProfile4UI数量
func (this *AccessAuthProfile4UIMongoRepo)QueryAccessAuthProfile4UICount(query map[string]interface{}) (int64,error) {
	count,err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfile4UICol).Find(query).Count()
	if err != nil {
		return -1,err
	}
	return int64(count), nil
}	

//更新AccessAuthProfile4UI记录
func (this *AccessAuthProfile4UIMongoRepo) UpdateAccessAuthProfile4UI(selector , values map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfile4UICol).UpdateAll(selector, values)
	if err != nil {
		return err
	}
	return nil

}	

//删除AccessAuthProfile4UI记录
func (this *AccessAuthProfile4UIMongoRepo) DeleteAccessAuthProfile4UI(selector map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfile4UICol).RemoveAll(selector)
	if err != nil {
		return err
	}
	return nil

}	

//插入AccessAuthProfile4UI记录
func (this *AccessAuthProfile4UIMongoRepo) InsertAccessAuthProfile4UI(entities ...*SELFENTITY.AccessAuthProfile4UI) error {
	entitiesInterface:= []interface{}{}
	for _,entity:=range entities{
		entitiesInterface=append(entitiesInterface,entity)
	}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfile4UICol).Insert(entitiesInterface...)
	if err != nil {
		return err
	}
	return nil
}


//关闭repo
func (this *AccessAuthProfile4UIMongoRepo) Close() error {
	this.session.Close()
	return nil
}
