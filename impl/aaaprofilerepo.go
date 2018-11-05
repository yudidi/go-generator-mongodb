
package impl
import (
	"gopkg.in/mgo.v2"
	MONGO "han-networks.com/csp/common_grpc/mongo"
	CONFIG "han-networks.com/csp/config_grpc/config"
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
	REPO "mongorepo/inf"
)





//获取AAAProfileMongoRepo对象
func NewAAAProfileMongoRepo() REPO.AAAProfileRepo {
	return &AAAProfileMongoRepo{session: MONGO.GetSession()}
}

//持久层mongo实现
type AAAProfileMongoRepo struct {
	session *mgo.Session
}



func (this *AAAProfileMongoRepo)QueryAAAProfileOne(query interface{}) (*SELFENTITY.AAAProfile,error) {
	entity := SELFENTITY.AAAProfile{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AAAProfileCol).Find(query).One(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}	

func (this *AAAProfileMongoRepo)QueryAAAProfileAll(query map[string]interface{}) (*[]*SELFENTITY.AAAProfile,error) {
	entities := []*SELFENTITY.AAAProfile{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AAAProfileCol).Find(query).All(entities)
	if err != nil {
		return nil, err
	}
	return &entities, nil
}	

func (this *AAAProfileMongoRepo)QueryAAAProfilePage(query map[string]interface{}, limit int, sorts ...string) (*[]*SELFENTITY.AAAProfile,error) {
	q := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AAAProfileCol).Find(query)
	if sorts != nil && len(sorts) > 0 {
		for _, s := range sorts {
			q.Sort(s)
		}
	}
	entities := []*SELFENTITY.AAAProfile{}
	q.Limit(limit).All(&entities)

	return &entities, nil

}	

func (this *AAAProfileMongoRepo) UpdateAAAProfile(selector , values map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AAAProfileCol).UpdateAll(selector, values)
	if err != nil {
		return err
	}
	return nil

}	

func (this *AAAProfileMongoRepo) DeleteAAAProfile(selector map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AAAProfileCol).RemoveAll(selector)
	if err != nil {
		return err
	}
	return nil

}	

func (this *AAAProfileMongoRepo) InsertAAAProfile(entities ...*SELFENTITY.AAAProfile) error {
	entitiesInterface:= []interface{}{}
	for _,entity:=range entities{
		entitiesInterface=append(entitiesInterface,entity)
	}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.AAAProfileCol).Insert(entitiesInterface...)
	if err != nil {
		return err
	}
	return nil
}


func (this *AAAProfileMongoRepo) Close() error {
	this.session.Close()
	return nil
}
