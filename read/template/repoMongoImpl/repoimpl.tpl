package mongorepo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	MONGO "han-networks.com/csp/common_grpc/mongo"
	CONFIG "han-networks.com/csp/config_grpc/config"
	SelfEntity "han-networks.com/csp/config_grpc/entity"
	REPO "han-networks.com/csp/config_grpc/server"
)

//获取$entityNameMongoRepo对象
func New$entityNameMongoRepo() REPO.$entityNameRepo {
	return &$entityNameMongoRepo{session: MONGO.GetSession()}
}

//持久层mongo实现
type $entityNameMongoRepo struct {
	session *mgo.Session
}

//查询一条$entityName记录
func (this *$entityNameMongoRepo) Query$entityNameOne(query map[string]interface{}) (*SelfEntity.$entityName, error) {
	entity := SelfEntity.$entityName{}
	err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.$entityNameCol).Find(query).One(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

//查询$entityName指定字段
func (this *$entityNameMongoRepo) Query$entityNameDistinct(query map[string]interface{}, field string, result interface{}) error {
	err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.$entityNameCol).Find(query).Distinct(field, result)
	if err != nil {
		return err
	}
	return nil
}

//查询所有$entityName记录
func (this *$entityNameMongoRepo) Query$entityNameAll(query map[string]interface{}) ([]*SelfEntity.$entityName, error) {
	entities := []*SelfEntity.$entityName{}
	err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.$entityNameCol).Find(query).All(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

//查询$entityName分页排序记录
func (this *$entityNameMongoRepo) Query$entityNamePage(query map[string]interface{}, limit int, sorts ...string) ([]*SelfEntity.$entityName, error) {
	q := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.$entityNameCol).Find(query)
	if sorts != nil && len(sorts) > 0 {
		for _, s := range sorts {
			q.Sort(s)
		}
	}
	entities := []*SelfEntity.$entityName{}
	q.Limit(limit).All(&entities)

	return entities, nil

}

//查询$entityName数量
func (this *$entityNameMongoRepo) Query$entityNameCount(query map[string]interface{}) (int64, error) {
	count, err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.$entityNameCol).Find(query).Count()
	if err != nil {
		return -1, err
	}
	return int64(count), nil
}

//更新$entityName记录
func (this *$entityNameMongoRepo) Update$entityName(selector, values map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.$entityNameCol).UpdateAll(selector, bson.M{"$set": values})
	if err != nil {
		return err
	}
	return nil

}

//删除$entityName记录
func (this *$entityNameMongoRepo) Delete$entityName(selector map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.$entityNameCol).RemoveAll(selector)
	if err != nil {
		return err
	}
	return nil

}

//插入$entityName记录
func (this *$entityNameMongoRepo) Insert$entityName(entities ...*SelfEntity.$entityName) error {
	entitiesInterface := []interface{}{}
	for _, entity := range entities {
		entitiesInterface = append(entitiesInterface, entity)
	}
	err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.$entityNameCol).Insert(entitiesInterface...)
	if err != nil {
		return err
	}
	return nil
}

//关闭repo
func (this *$entityNameMongoRepo) Close() error {
	this.session.Close()
	return nil
}
