package main

import "han-networks.com/csp/entity"

// 根据entity生成proto中的msg对象.
func main(){
	in := entity.GroupClientBlacklistRecord{}
	out := ""


}

func loopEntity_GeneProtoMessageDTO(){
	template := `message ClientBlacklistRecordDTO {
    int64 Id = 1;
    int64 Siteid = 2;
    string Mac = 3;
    bool Auto = 4; // true:动态加黑，false，静态加黑
    bool Delete = 5; // true : 已删除， false ：未删除
    string StartTime = 6;
    string EndTime = 7;
    string DeleteTime = 8;
}
`
	templateStart := `message ClientBlacklistRecordDTO {`
	templateEnd := `}`

	/*Output用于拼接新的message对象的字符串. 直接放入proto文件中.
	 func f(结构体,output string)
	 1. 遍历接结构体的字段
	 1.1 如果字段是基本类型string,int,bool
		{
		 1. 构造一个proto.message字段
		 2. templateStart拼接
		 3.
		}
	1.2 如果字段kind==struct.
	1.2.1 内嵌的struct(内嵌DO). xxxDO yyyDO
		{
	      1. 递归调用(子struct,output string0)
		}
	1.2.2 其他struct. time.Time. 转换为string进行处理.
	 */
}
