package plugins

import (
	"fmt"
	"testing"
)

func TestRds(t *testing.T) {
	params := CloudProviderParam{
		IdentityParams: "SecretKey=vTaHbesCwh6R82Z6iXF7I8QPplIJGuHOiUFwQTUo;AccessKey=SNXGZX0W2WLXWHXK3ZQH;DomainId=07b045206c0025230f9ec00fc48612a0",
		CloudParams:    "CloudApiDomainName=myhuaweicloud.com;Region=ap-southeast-3;ProjectId=08485cf5ad80f2622f6dc00f79fb4ae0",
	}
	input := RdsCreateInput{
		CloudProviderParam: params,
	}
	instanceId := "7f89313baac643af9109c874d66f0b95in01"
	err := updateRdsConfiguration(&input, instanceId)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("done")
}

// func TestGetConfiguration(t *testing.T) {
// 	sc, err := createGolangSdkRdsServiceClientV3(input.CloudProviderParam)
// 	if err != nil {
// 		return err
// 	}
// 	instanceId := "7f89313baac643af9109c874d66f0b95in01"
// 	goConf
// }

// func TestGetFlavorRef(t *testing.T) {
// 	params := CloudProviderParam{
// 		IdentityParams: "SecretKey=vTaHbesCwh6R82Z6iXF7I8QPplIJGuHOiUFwQTUo;AccessKey=SNXGZX0W2WLXWHXK3ZQH;DomainId=07b045206c0025230f9ec00fc48612a0",
// 		CloudParams:    "CloudApiDomainName=myhuaweicloud.com;Region=ap-southeast-3;ProjectId=08485cf5ad80f2622f6dc00f79fb4ae0",
// 	}

// 	input := RdsCreateInput{
// 		HostType:           "1c1g",
// 		CloudProviderParam: params,
// 	}
// 	flavor, cpu, memory, err := getRdsFlavorByHostType(&input)
// 	if err != nil {
// 		fmt.Printf("error=%v\n", err)
// 	}
// 	fmt.Printf("flavor=%v, cpu=%v, memory=%v\n", flavor, cpu, memory)
// }

func TestUpdateRdsConfiguration(t *testing.T) {
	params := CloudProviderParam{
		IdentityParams: "SecretKey=vTaHbesCwh6R82Z6iXF7I8QPplIJGuHOiUFwQTUo;AccessKey=SNXGZX0W2WLXWHXK3ZQH;DomainId=07b045206c0025230f9ec00fc48612a0",
		CloudParams:    "CloudApiDomainName=myhuaweicloud.com;Region=ap-southeast-3;ProjectId=08485cf5ad80f2622f6dc00f79fb4ae0",
	}

	input := RdsCreateInput{
		HostType:           "1c1g",
		CloudProviderParam: params,
	}
	err := updateRdsConfiguration(&input, "c20da5cf536c45948f7afe058579b089in01")
	if err != nil {
		fmt.Printf("error=%v", err)
	}
}

func TestDeleteRdsConfiguration(t *testing.T) {
	params := CloudProviderParam{
		IdentityParams: "SecretKey=1dYRU1DBVaYKSixG0kcw2kqIehgEGlEdT4nb7dLN;AccessKey=Y361GSQB7SZPTRBZVAWI;DomainId=07bfdaf7d00026f90f8ac01118d2e880",
		CloudParams:    "CloudApiDomainName=myhuaweicloud.com;Region=ap-southeast-3;ProjectId=07c34b94508010ff2fb3c011b4a986e3",
	}
	id := "960d871022dd4234bae50e860fde429bpr01"
	err := deleteConfiguration(params, id)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	fmt.Printf("done\n")
}
