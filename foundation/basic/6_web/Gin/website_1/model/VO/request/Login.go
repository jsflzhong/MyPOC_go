package request

/*
注意两个字段都是binding
 */
type Login struct {
	//针对不同的请求的header中的content-type,给出了不同的匹配的key. 下面分别指定了form, json, uri, xml的四种content-type下的key.
	User     string `form:"username" json:"user" uri:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}