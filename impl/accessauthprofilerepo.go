
package impl
import (
	"gopkg.in/mgo.v2"
	MONGO "han-networks.com/csp/common_grpc/mongo"
	CONFIG "han-networks.com/csp/config_grpc/config"
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
	REPO "mongorepo/inf"
)





//获取AccessAuthProfileMongoRepo对象
func NewAccessAuthProfileMongoRepo() REPO.AccessAuthProfileRepo {
	return &AccessAuthProfileMongoRepo{session: MONGO.GetSession()}
}

//持久层mongo实现
type AccessAuthProfileMongoRepo struct {
	session *mgo.Session
}



//查询一条AccessAuthProfile记录
func (this *AccessAuthProfileMongoRepo)QueryAccessAuthProfileOne(query map[string]interface{}) (*SELFENTITY.AccessAuthProfile,error) {
	entity := SELFENTITY.AccessAuthProfile{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfileCol).Find(query).One(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}	

//查询AccessAuthProfile指定字段
func (this *AccessAuthProfileMongoRepo)QueryAccessAuthProfileField(query map[string]interface{},field string) ([]interface{},error) {
	selector:=map[string]interface{}{
		"_id":0,
	}
	selector[field]=1
	entityMap := []map[string]interface{}{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfileCol).Find(query).Select(selector).All(&entityMap)
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

//查询所有AccessAuthProfile记录
func (this *AccessAuthProfileMongoRepo)QueryAccessAuthProfileAll(query map[string]interface{}) (*[]*SELFENTITY.AccessAuthProfile,error) {
	entities := []*SELFENTITY.AccessAuthProfile{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfileCol).Find(query).All(entities)
	if err != nil {
		return nil, err
	}
	return &entities, nil
}	

//查询AccessAuthProfile分页排序记录
func (this *AccessAuthProfileMongoRepo)QueryAccessAuthProfilePage(query map[string]interface{}, limit int, sorts ...string) (*[]*SELFENTITY.AccessAuthProfile,error) {
	q := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfileCol).Find(query)
	if sorts != nil && len(sorts) > 0 {
		for _, s := range sorts {
			q.Sort(s)
		}
	}
	entities := []*SELFENTITY.AccessAuthProfile{}
	q.Limit(limit).All(&entities)

	return &entities, nil

}	

//查询AccessAuthProfile数量
func (this *AccessAuthProfileMongoRepo)QueryAccessAuthProfileCount(query map[string]interface{}) (int64,error) {
	count,err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfileCol).Find(query).Count()
	if err != nil {
		return -1,err
	}
	return int64(count), nil
}	

//更新AccessAuthProfile记录
func (this *AccessAuthProfileMongoRepo) UpdateAccessAuthProfile(selector , values map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfileCol).UpdateAll(selector, values)
	if err != nil {
		return err
	}
	return nil

}	

//删除AccessAuthProfile记录
func (this *AccessAuthProfileMongoRepo) DeleteAccessAuthProfile(selector map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfileCol).RemoveAll(selector)
	if err != nil {
		return err
	}
	return nil

}	

//插入AccessAuthProfile记录
func (this *AccessAuthProfileMongoRepo) InsertAccessAuthProfile(entities ...*SELFENTITY.AccessAuthProfile) error {
	entitiesInterface:= []interface{}{}
	for _,entity:=range entities{
		entitiesInterface=append(entitiesInterface,entity)
	}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AccessAuthProfileCol).Insert(entitiesInterface...)
	if err != nil {
		return err
	}
	return nil
}


//关闭repo
func (this *AccessAuthProfileMongoRepo) Close() error {
	this.session.Close()
	return nil
}
