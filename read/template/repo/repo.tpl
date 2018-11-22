package $PackageName

import (
	SelfEntity "han-networks.com/csp/config_grpc/entity/$PackageName"
)

type $entityNameRepo interface {
	//查询一条 $entityName 记录
	QueryOne$entityName(query map[string]interface{}) (*SelfEntity.$entityName,error)

	//查询 $entityName 指定字段
	QueryDistinct$entityName(query map[string]interface{},field string,result interface{}) (error)

	//查询所有 $entityName 记录
	QueryAll$entityName(query map[string]interface{}) ([]*SelfEntity. xx ,error) 

	//查询 $entityName 分页排序记录
	QueryPage$entityName(query map[string]interface{}, limit int, sorts ...string) ([]*SelfEntity.$entityName,error)

	//查询 $entityName 数量
	QueryCount$entityName(query map[string]interface{}) (int64,error) 

	//更新 $entityName 记录
	Update$entityName(selector , values map[string]interface{}) error

	//删除 $entityName 记录
	Delete$entityName (selector map[string]interface{}) error

	//插入 $entityName 记录
	Insert$entityName (entities ...*SelfEntity.$entityName ) error

	//关闭repo
	Close() error
}