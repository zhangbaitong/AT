package resource

//注册中心-资源表
type Resource struct {
	res_id        string //资源ID
	res_name      string //资源名称
	owner_acid    int    //拥有者账号ID
	operator_acid int    //授权者
	status        int    //资源状态
	create_time   int    //创建时间
}
