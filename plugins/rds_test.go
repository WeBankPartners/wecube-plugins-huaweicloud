package plugins

import (
	"fmt"
	"testing"
)

func TestIsExistRds(t *testing.T) {
	param := CloudProviderParam{
		CloudParams:    "CloudApiDomainName=myhuaweicloud.com;Region=ap-southeast-3;ProjectId=08485cf5ad80f2622f6dc00f79fb4ae0",
		IdentityParams: "SecretKey=vTaHbesCwh6R82Z6iXF7I8QPplIJGuHOiUFwQTUo;AccessKey=SNXGZX0W2WLXWHXK3ZQH;DomainId=07b045206c0025230f9ec00fc48612a0",
	}
	sc, _ := createRdsServiceClientV3(param)
	instanceId := "sssss"
	rdsInfo, exist, err := isRdsExist(sc, instanceId)
	fmt.Println("rdsInfo:", rdsInfo)
	fmt.Println("exist:", exist)
	fmt.Println("err:", err)
	fmt.Println("done")
}
