
package inf
import (
	SELFENTITY "han-networks.com/csp/config_grpc/entity"
)



type AAAProfileRepo interface {
	QueryAAAProfileOne(query interface{}) (*SELFENTITY.AAAProfile,error) 	

	QueryAAAProfileAll(query map[string]interface{}) (*[]*SELFENTITY.AAAProfile,error) 

	QueryAAAProfilePage(query map[string]interface{}, limit int, sorts ...string) (*[]*SELFENTITY.AAAProfile,error) 

	UpdateAAAProfile(selector , values map[string]interface{}) error

	DeleteAAAProfile(selector map[string]interface{}) error

	InsertAAAProfile(entities ...*SELFENTITY.AAAProfile) error

	Close() error

}