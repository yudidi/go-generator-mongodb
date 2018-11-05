package main

import (
	"han-networks.com/csp/config_grpc/entity"
	"os"
	"reflect"
)

type m map[string]interface{}

//导包路径
const (
	MONGO         = "han-networks.com/csp/common_grpc/mongo"
	CONFIG        = "han-networks.com/csp/config_grpc/config"
	SELFENTITY    = "han-networks.com/csp/config_grpc/entity"
	REPOINTERFACE = "han-networks.com/csp/config_grpc/server"
)

//生成go文件路径
const (
	GOFILEPATH = "d:/repo.go"
)
//不需更新字段
var (
	EXCLUDEBSON=[]string{"_id","scenetype","sceneid"}
)

func main() {
	entity := entity.AAAProfile{}
	repoStr := generateRepo(&entity, entity)
	writeInfo2GoFile(repoStr, GOFILEPATH)
}
func generateRepo(entityPointer interface{}, entity interface{}) string {
	entityPointerType := reflect.TypeOf(entityPointer)
	entityType := reflect.TypeOf(entity)
	entityName := getEntityName(entityType)
	entityCollectionName := getEntityCollectionName(entityName)
	entityRepoName := getEntityRepoName(entityName)
	entityRepoInterfaceName := getEntityRepoInterfaceName(entityName)
	bsonTagMap := getBSONTagMap(entityPointerType)
	//生成
	pkg := getPackage(entityRepoInterfaceName, entityRepoName)
	queryOne := generateQueryOne(entityName, entityRepoName, entityCollectionName)
	queryAll := generateQueryAll(entityName, entityRepoName, entityCollectionName)
	queryPage := generateQueryPage(entityName, entityRepoName, entityCollectionName)
	update := generateUpdate(entityName, entityRepoName, entityCollectionName, bsonTagMap)
	delete := generateDelete(entityRepoName, entityCollectionName)
	insert:=generateInsert(entityName, entityRepoName, entityCollectionName)
	close := generateClose(entityRepoName)
	return pkg + queryOne + queryAll + queryPage + update + delete +insert+ close
}

func getEntityRepoName(entityName string) string {
	return entityName + "MongoRepo"
}
func getEntityRepoInterfaceName(entityName string) string {
	return entityName + "Repo"
}
func getPackage(entityRepoInterfaceName, entityRepoName string) string {
	return `
package repo
import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	MONGO "` + MONGO + `"
	CONFIG "` + CONFIG + `"
	SELFENTITY "` + SELFENTITY + `"
	REPO "` + REPOINTERFACE + `"
)

//获取` + entityRepoName + `对象
func New` + entityRepoName + `() REPO.` + entityRepoInterfaceName + ` {
	return &` + entityRepoName + `{session: MONGO.GetSession()}
}

//持久层mongo实现
type ` + entityRepoName + ` struct {
	session *mgo.Session
}


`
}

func generateQueryOne(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `)QueryOne(query interface{}) (*SELFENTITY.` + entityName + `,error) {
	entity := SELFENTITY.` + entityName + `{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.` + entityCollectionName + `).Find(query).One(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}	
`

	return repoStr
}

func generateQueryAll(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `)QueryAll(query map[string]interface{}) (*[]*SELFENTITY.` + entityName + `,error) {
	entities := []*SELFENTITY.` + entityName + `{}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.` + entityCollectionName + `).Find(query).All(entities)
	if err != nil {
		return nil, err
	}
	return &entities, nil
}	
`
	return repoStr
}
func generateQueryPage(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `)QueryPage(query map[string]interface{}, limit int, sorts ...string) (*[]*SELFENTITY.` + entityName + `,error) {
	q := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.` + entityCollectionName + `).Find(query)
	if sorts != nil && len(sorts) > 0 {
		for _, s := range sorts {
			q.Sort(s)
		}
	}
	entities := []*SELFENTITY.` + entityName + `{}
	q.Limit(limit).All(&entities)

	return &entities, nil

}	
`
	return repoStr
}
func generateUpdate(entityName, entityRepoName, entityCollectionName string, bsonTagMap m) string {

	updatedStr := ""
	for k, v := range bsonTagMap {
		vv := v.(string)
		if ContainStr(EXCLUDEBSON,vv){
			continue
		}
		updatedStr += `
	"` + vv + `":entity.` + k + `,`
	}

	repoStr := `
func (this *` + entityRepoName + `) Update(selector map[string]interface{}, entity *SELFENTITY.` + entityName + `) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.` + entityCollectionName + `).UpdateAll(selector, bson.M{"$set": bson.M{
		` + updatedStr + `
	}})
	if err != nil {
		return err
	}
	return nil

}	
`
	return repoStr
}

func generateDelete(entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `) Delete(selector map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.` + entityCollectionName + `).RemoveAll(selector)
	if err != nil {
		return err
	}
	return nil

}	
`
	return repoStr

}
func generateInsert(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `) Insert(entity *SELFENTITY.`+entityName+`) error {
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.`+entityCollectionName+`).Insert(entity)
	if err != nil {
		return err
	}
	return nil
}

`
	return repoStr

}
func generateClose(entityRepoName string) string {
	repoStr := `
func (this *` + entityRepoName + `) Close() error {
	this.session.Close()
	return nil
}
`
	return repoStr
}

func getEntityCollectionName(entityName string) string {
	return entityName + "Col"
}
func getEntityName(entityType reflect.Type) string {
	return entityType.Name()
}
func getBSONTagMap(entityType reflect.Type) m {
	jsonTagMap := m{}
	for i := 0; i < entityType.Elem().NumField(); i++ {
		jsonTagMap[entityType.Elem().Field(i).Name] = entityType.Elem().Field(i).Tag.Get("bson")
	}
	return jsonTagMap
}

func writeInfo2GoFile(repoStr, goFilePath string) {
	goFile, _ := os.OpenFile(goFilePath, os.O_RDWR|os.O_CREATE, 0766)
	defer goFile.Close()
	goFile.WriteString(repoStr)
}

//数组中是否包含字符串元素
func ContainStr(source []string, target string) bool {
	for _, s := range source {
		if s == target {
			return true
		}
	}
	return false
}
