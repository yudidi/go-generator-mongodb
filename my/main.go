package main

import (
	Entity_In "han-networks.com/csp/config_grpc/entity/blacklist"
	"os"
	"reflect"
	"strings"
	"fmt"
)

/*
 * repo层生成器。根据指定结构体生成golang语言repo接口和mongo实现。
 * 注意:
 *	项目需要放置在%GOPATH%/src下
 *  TODO表示需要配置的地方
 *
 */
//需修改
const (
	//TODO
	MONGO = "han-networks.com/csp/common_grpc/mongo"
	//TODO
	CONFIG = "han-networks.com/csp/config_grpc/config"
	//TODO
	SelfEntity = "han-networks.com/csp/config_grpc/entity/bl"
	//TODO 需要修改当前项目所在路径
	PROJECT_PATH = "D:\\GOPATH\\src\\github.com\\gorepomaker\\"
	OutPut_Path = "D:\\GOPATH\\src\\github.com\\gorepomaker\\" +"blacklist\\"
)

//无需修改
const (
	REPO_IMPL_GO_FILE_PATH = OutPut_Path + "impl\\"
	REPO_INF_GO_FILE_PATH  = OutPut_Path + "inf\\"
	REPO_INTERFACE         = "gorepomaker/inf"
)

func init() {
	exist, err := PathExists(OutPut_Path)
	if err != nil {
		panic(err)
	}

	if !exist {
		// 创建文件夹
		err = os.Mkdir(OutPut_Path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	//TODO 需要修改entities数组
	entities := []interface{}{
		Entity_In.ClientBlacklistRecord{},
		Entity_In.GroupClientBlacklistRecord{},
		Entity_In.BlackClient{},
	}
	batchGenerate(entities...)
}

//执行批量生成
func batchGenerate(entities ...interface{}) {
	for _, entity := range entities {
		generate(entity)
	}
}

//执行生成
func generate(entity interface{}) {
	repoInfStr := generateRepoInf(entity)
	repoImplStr := generateRepoImpl(entity)
	goRepoFileName := generateGoRepoFileName(entity)
	writeInfo2GoFile(repoInfStr, REPO_INF_GO_FILE_PATH, goRepoFileName+".go")
	writeInfo2GoFile(repoImplStr, REPO_IMPL_GO_FILE_PATH, goRepoFileName+".go")
	fmt.Println("Generate " + goRepoFileName + " , mission success !")
}

//生成repo接口
func generateRepoInf(entity interface{}) string {
	entityType := reflect.TypeOf(entity)
	entityName := getEntityName(entityType)
	entityRepoName := getEntityRepoInfName(entityName)
	//生成接口
	pkg := getInfPackage()
	queryOneInf := generateQueryOneInf(entityName)
	queryDistinctInf := generateQueryDistinctInf(entityName)
	queryAllInf := generateQueryAllInf(entityName)
	queryPageInf := generateQueryPageInf(entityName)
	queryCount := generateQueryCountInf(entityName)
	updateInf := generateUpdateInf(entityName)
	deleteInf := generateDeleteInf(entityName)
	insertInf := generateInsertInf(entityName)
	closeInf := generateCloseInf()
	repoInf := queryOneInf + queryDistinctInf + queryAllInf + queryPageInf + queryCount + updateInf + deleteInf + insertInf + closeInf
	return pkg + `
type ` + entityRepoName + ` interface {` + repoInf + `
}`
}

//生成repo的mongo实现
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
	queryDistinct := generateQueryDistinct(entityName, entityRepoName, entityCollectionName)
	queryAll := generateQueryAll(entityName, entityRepoName, entityCollectionName)
	queryPage := generateQueryPage(entityName, entityRepoName, entityCollectionName)
	queryCount := generateQueryCount(entityName, entityRepoName, entityCollectionName)
	update := generateUpdate(entityName, entityRepoName, entityCollectionName)
	delete := generateDelete(entityName, entityRepoName, entityCollectionName)
	insert := generateInsert(entityName, entityRepoName, entityCollectionName)
	close := generateClose(entityRepoName)

	repoImpl := queryOne + queryDistinct + queryAll + queryPage + queryCount + update + delete + insert + close
	return pkg + repoStruct + repoImpl
}

func generateGoRepoFileName(entity interface{}) string {
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
	return `package impl
import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	MONGO "` + MONGO + `"
	CONFIG "` + CONFIG + `"
	SelfEntity "` + SelfEntity + `"
	REPO "` + REPO_INTERFACE + `"
)


`
}
func getInfPackage() string {
	return `package inf
import (
	SelfEntity "` + SelfEntity + `"
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
//查询一条` + entityName + `记录
func (this *` + entityRepoName + `)Query` + entityName + `One(query map[string]interface{}) (*SelfEntity.` + entityName + `,error) {
	entity := SelfEntity.` + entityName + `{}
	err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.` + entityCollectionName + `).Find(query).One(&entity)
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
	//查询一条` + entityName + `记录
	Query` + entityName + `One(query map[string]interface{}) (*SelfEntity.` + entityName + `,error) 	
`
	return repoStr
}

func generateQueryDistinct(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
//查询` + entityName + `指定字段
func (this *` + entityRepoName + `)Query` + entityName + `Distinct(query map[string]interface{},field string,result interface{}) (error) {
	err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.` + entityCollectionName + `).Find(query).Distinct(field,result)
	if err != nil {
		return err
	}
	return nil
}	
`
	return repoStr
}
func generateQueryDistinctInf(entityName string) string {
	repoStr := `
	//查询` + entityName + `指定字段
	Query` + entityName + `Distinct(query map[string]interface{},field string,result interface{}) (error)
`
	return repoStr

}
func generateQueryAll(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
//查询所有` + entityName + `记录
func (this *` + entityRepoName + `)Query` + entityName + `All(query map[string]interface{}) ([]*SelfEntity.` + entityName + `,error) {
	entities := []*SelfEntity.` + entityName + `{}
	err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.` + entityCollectionName + `).Find(query).All(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}	
`
	return repoStr
}
func generateQueryAllInf(entityName string) string {
	repoStr := `
	//查询所有` + entityName + `记录
	Query` + entityName + `All(query map[string]interface{}) ([]*SelfEntity.` + entityName + `,error) 
`
	return repoStr
}
func generateQueryPage(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
//查询` + entityName + `分页排序记录
func (this *` + entityRepoName + `)Query` + entityName + `Page(query map[string]interface{}, limit int, sorts ...string) ([]*SelfEntity.` + entityName + `,error) {
	q := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.` + entityCollectionName + `).Find(query)
	if sorts != nil && len(sorts) > 0 {
		for _, s := range sorts {
			q.Sort(s)
		}
	}
	entities := []*SelfEntity.` + entityName + `{}
	q.Limit(limit).All(&entities)

	return entities, nil

}	
`
	return repoStr
}
func generateQueryPageInf(entityName string) string {
	repoStr := `
	//查询` + entityName + `分页排序记录
	Query` + entityName + `Page(query map[string]interface{}, limit int, sorts ...string) ([]*SelfEntity.` + entityName + `,error) 
`
	return repoStr
}
func generateQueryCount(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
//查询` + entityName + `数量
func (this *` + entityRepoName + `)Query` + entityName + `Count(query map[string]interface{}) (int64,error) {
	count,err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.` + entityCollectionName + `).Find(query).Count()
	if err != nil {
		return -1,err
	}
	return int64(count), nil
}	
`
	return repoStr
}
func generateQueryCountInf(entityName string) string {
	repoStr := `
	//查询` + entityName + `数量
	Query` + entityName + `Count(query map[string]interface{}) (int64,error) 
`
	return repoStr
}
func generateUpdate(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
//更新` + entityName + `记录
func (this *` + entityRepoName + `) Update` + entityName + `(selector , values map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.` + entityCollectionName + `).UpdateAll(selector, bson.M{"$set": values})
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
	//更新` + entityName + `记录
	Update` + entityName + `(selector , values map[string]interface{}) error
`
	return repoStr
}
func generateDelete(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
//删除` + entityName + `记录
func (this *` + entityRepoName + `) Delete` + entityName + `(selector map[string]interface{}) error {
	_, err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.` + entityCollectionName + `).RemoveAll(selector)
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
	//删除` + entityName + `记录
	Delete` + entityName + `(selector map[string]interface{}) error
`
	return repoStr
}
func generateInsert(entityName, entityRepoName, entityCollectionName string) string {
	repoStr := `
//插入` + entityName + `记录
func (this *` + entityRepoName + `) Insert` + entityName + `(entities ...*SelfEntity.` + entityName + `) error {
	entitiesInterface:= []interface{}{}
	for _,entity:=range entities{
		entitiesInterface=append(entitiesInterface,entity)
	}
	err := this.session.DB(CONFIG.MgoDBName).C(SelfEntity.` + entityCollectionName + `).Insert(entitiesInterface...)
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
	//插入` + entityName + `记录
	Insert` + entityName + `(entities ...*SelfEntity.` + entityName + `) error
`
	return repoStr
}
func generateClose(entityRepoName string) string {
	repoStr := `
//关闭repo
func (this *` + entityRepoName + `) Close() error {
	this.session.Close()
	return nil
}
`
	return repoStr
}
func generateCloseInf() string {
	repoStr := `
	//关闭repo
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

	if !exist {
		// 创建文件夹
		err = os.Mkdir(goFilePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
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
