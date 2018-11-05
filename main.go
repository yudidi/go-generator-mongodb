package main

import (
	"han-networks.com/csp/config_grpc/entity"
	"os"
	"reflect"
	"strings"
)

//导包路径
const (
	MONGO          = "han-networks.com/csp/common_grpc/mongo"
	CONFIG         = "han-networks.com/csp/config_grpc/config"
	SELFENTITY     = "han-networks.com/csp/config_grpc/entity"
	REPO_INTERFACE = "mongorepo/inf"
)

//生成go文件路径
const (
	//TODO
	REPO_IMPL_GO_FILE_PATH = "D:\\project\\go\\src\\mongorepo\\impl\\"
	//TODO
	REPO_INF_GO_FILE_PATH = "D:\\project\\go\\src\\mongorepo\\inf\\"
)

func main() {
	entity := entity.AAAProfile{}
	repoInfStr := generateRepoInf(entity)
	repoImplStr := generateRepoImpl(entity)
	goRepoFileName := generateGoRepoFileName(entity)
	writeInfo2GoFile(repoInfStr, REPO_INF_GO_FILE_PATH, goRepoFileName+".go")
	writeInfo2GoFile(repoImplStr, REPO_IMPL_GO_FILE_PATH, goRepoFileName+".go")
}
func generateRepoInf(entity entity.AAAProfile) string {
	entityType := reflect.TypeOf(entity)
	entityName := getEntityName(entityType)
	entityRepoName := getEntityRepoInfName(entityName)
	//生成接口
	pkg := getInfPackage()
	queryOneInf := generateQueryOneInf(entityName)
	queryAllInf := generateQueryAllInf(entityName)
	queryPageInf := generateQueryPageInf(entityName)
	updateInf := generateUpdateInf(entityName)
	deleteInf := generateDeleteInf(entityName)
	insertInf := generateInsertInf(entityName)
	closeInf := generateCloseInf(entityRepoName)
	repoInf := queryOneInf + queryAllInf + queryPageInf + updateInf + deleteInf + insertInf + closeInf
	return pkg + `
type ` + entityRepoName + ` interface {` + repoInf + `
}`
}
func generateRepoImpl(entity interface{}) string {
	entityType := reflect.TypeOf(entity)
	entityName := getEntityName(entityType)
	entityCollectionName := getEntityCollectionName(entityName)
	entityRepoName := getEntityRepoImplName(entityName)
	entityRepoInterfaceName := getEntityRepoInterfaceName(entityName)

	//生成实现
	pkg := getImplPackage()
	repoStruct := getRepoStruct(entityRepoInterfaceName, entityRepoName)
	queryOne := generateQueryOne(entityName, entityRepoName, entityCollectionName)
	queryAll := generateQueryAll(entityName, entityRepoName, entityCollectionName)
	queryPage := generateQueryPage(entityName, entityRepoName, entityCollectionName)
	update := generateUpdate(entityName, entityRepoName, entityCollectionName)
	delete := generateDelete(entityName, entityRepoName, entityCollectionName)
	insert := generateInsert(entityName, entityRepoName, entityCollectionName)
	close := generateClose(entityRepoName)

	repoImpl := queryOne + queryAll + queryPage + update + delete + insert + close
	return pkg + repoStruct + repoImpl
}

func generateGoRepoFileName(entity entity.AAAProfile) string {
	entityType := reflect.TypeOf(entity)
	return strings.ToLower(entityType.Name()) + "repo"
}
func getEntityRepoImplName(entityName string) string {
	return entityName + "MongoRepo"
}
func getEntityRepoInfName(entityName string) string {
	return entityName + "Repo"
}
func getEntityRepoInterfaceName(entityName string) string {
	return entityName + "Repo"
}
func getImplPackage() string {
	return `
package impl
import (
	"gopkg.in/mgo.v2"
	MONGO "` + MONGO + `"
	CONFIG "` + CONFIG + `"
	SELFENTITY "` + SELFENTITY + `"
	REPO "` + REPO_INTERFACE + `"
)


`
}
func getInfPackage() string {
	return `
package inf
import (
	SELFENTITY "` + SELFENTITY + `"
)


`
}
func getRepoStruct(entityRepoInterfaceName, entityRepoName string) string {
	return `


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
func (this *` + entityRepoName + `)Query` + entityName + `One(query interface{}) (*SELFENTITY.` + entityName + `,error) {
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
func generateQueryOneInf(entityName string) string {
	repoStr := `
	Query` + entityName + `One(query interface{}) (*SELFENTITY.` + entityName + `,error) 	
`
	return repoStr
}
func generateQueryAll(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `)Query` + entityName + `All(query map[string]interface{}) (*[]*SELFENTITY.` + entityName + `,error) {
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
func generateQueryAllInf(entityName string) string {
	repoStr := `
	Query` + entityName + `All(query map[string]interface{}) (*[]*SELFENTITY.` + entityName + `,error) 
`
	return repoStr
}
func generateQueryPage(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `)Query` + entityName + `Page(query map[string]interface{}, limit int, sorts ...string) (*[]*SELFENTITY.` + entityName + `,error) {
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
func generateQueryPageInf(entityName string) string {
	repoStr := `
	Query` + entityName + `Page(query map[string]interface{}, limit int, sorts ...string) (*[]*SELFENTITY.` + entityName + `,error) 
`
	return repoStr
}
func generateUpdate(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `) Update` + entityName + `(selector , values map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.` + entityCollectionName + `).UpdateAll(selector, values)
	if err != nil {
		return err
	}
	return nil

}	
`
	return repoStr
}
func generateUpdateInf(entityName string) string {
	repoStr := `
	Update` + entityName + `(selector , values map[string]interface{}) error
`
	return repoStr
}
func generateDelete(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `) Delete` + entityName + `(selector map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.` + entityCollectionName + `).RemoveAll(selector)
	if err != nil {
		return err
	}
	return nil

}	
`
	return repoStr

}
func generateDeleteInf(entityName string) string {
	repoStr := `
	Delete` + entityName + `(selector map[string]interface{}) error
`
	return repoStr
}
func generateInsert(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
func (this *` + entityRepoName + `) Insert` + entityName + `(entities ...*SELFENTITY.` + entityName + `) error {
	entitiesInterface:= []interface{}{}
	for _,entity:=range entities{
		entitiesInterface=append(entitiesInterface,entity)
	}
	err := this.session.DB(CONFIG.MgoDBName).C(SELFENTITY.` + entityCollectionName + `).Insert(entitiesInterface...)
	if err != nil {
		return err
	}
	return nil
}

`
	return repoStr

}
func generateInsertInf(entityName string) string {
	repoStr := `
	Insert` + entityName + `(entities ...*SELFENTITY.` + entityName + `) error
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
func generateCloseInf(entityRepoName string) string {
	repoStr := `
	Close() error
`
	return repoStr
}
func getEntityCollectionName(entityName string) string {
	return entityName + "Col"
}
func getEntityName(entityType reflect.Type) string {
	return entityType.Name()
}

func writeInfo2GoFile(repoStr, goFilePath, goFileName string) {
	exist, err := PathExists(goFilePath)
	if err != nil {
		panic(err)
	}
	if exist {
		os.RemoveAll(goFilePath)

	}
	// 创建文件夹
	err = os.Mkdir(goFilePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	goFile, err := os.OpenFile(goFilePath+goFileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		panic(err)
	}
	defer goFile.Close()
	goFile.WriteString(repoStr)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}