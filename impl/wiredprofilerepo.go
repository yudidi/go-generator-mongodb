package impl
import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	MONGO "han-networks.com/csp/common_grpc/mongo"
	CONFIG "han-networks.com/csp/config_grpc/config"
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
	REPO "gorepomaker/inf"
)





//获取WiredProfileMongoRepo对象
func NewWiredProfileMongoRepo() REPO.WiredProfileRepo {
	return &WiredProfileMongoRepo{session: MONGO.GetSession()}
}

//持久层mongo实现
type WiredProfileMongoRepo struct {
	session *mgo.Session
}



//查询一条WiredProfile记录
func (this *WiredProfileMongoRepo)QueryWiredProfileOne(query map[string]interface{}) (*SELFENTITY.WiredProfile,error) {
	entity := SELFENTITY.WiredProfile{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.WiredProfileCol).Find(query).One(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}	

//查询WiredProfile指定字段
func (this *WiredProfileMongoRepo)QueryWiredProfileDistinct(query map[string]interface{},field string,result interface{}) (error) {
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.WiredProfileCol).Find(query).Distinct(field,result)
	if err != nil {
		return err
	}
	return nil
}	

//查询所有WiredProfile记录
func (this *WiredProfileMongoRepo)QueryWiredProfileAll(query map[string]interface{}) ([]*SELFENTITY.WiredProfile,error) {
	entities := []*SELFENTITY.WiredProfile{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.WiredProfileCol).Find(query).All(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}	

//查询WiredProfile分页排序记录
func (this *WiredProfileMongoRepo)QueryWiredProfilePage(query map[string]interface{}, limit int, sorts ...string) ([]*SELFENTITY.WiredProfile,error) {
	q := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.WiredProfileCol).Find(query)
	if sorts != nil && len(sorts) > 0 {
		for _, s := range sorts {
			q.Sort(s)
		}
	}
	entities := []*SELFENTITY.WiredProfile{}
	q.Limit(limit).All(&entities)

	return entities, nil

}	

//查询WiredProfile数量
func (this *WiredProfileMongoRepo)QueryWiredProfileCount(query map[string]interface{}) (int64,error) {
	count,err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.WiredProfileCol).Find(query).Count()
	if err != nil {
		return -1,err
	}
	return int64(count), nil
}	

//更新WiredProfile记录
func (this *WiredProfileMongoRepo) UpdateWiredProfile(selector , values map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.WiredProfileCol).UpdateAll(selector, bson.M{"$set": values})
	if err != nil {
		return err
	}
	return nil

}	

//删除WiredProfile记录
func (this *WiredProfileMongoRepo) DeleteWiredProfile(selector map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.WiredProfileCol).RemoveAll(selector)
	if err != nil {
		return err
	}
	return nil

}	

//插入WiredProfile记录
func (this *WiredProfileMongoRepo) InsertWiredProfile(entities ...*SELFENTITY.WiredProfile) error {
	entitiesInterface:= []interface{}{}
	for _,entity:=range entities{
		entitiesInterface=append(entitiesInterface,entity)
	}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.WiredProfileCol).Insert(entitiesInterface...)
	if err != nil {
		return err
	}
	return nil
}


//关闭repo
func (this *WiredProfileMongoRepo) Close() error {
	this.session.Close()
	return nil
}
