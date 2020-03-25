# Wecube HuaweiCloud API Guid
  
提供统一接口定义，为使用者提供清晰明了的使用方法。

## API 操作资源（Resources）: 

**私有网络**

- [创建虚拟私有云](#vpc-create)  
- [私有虚拟私有云](#vpc-delete)

**子网**

- [子网创建](#subnet-create) 
- [子网销毁](#subnet-delete) 

**云服务器**

- [云服务器创建](#vm-create)
- [云服务器销毁](#vm-terminate)
- [云服务器启动](#vm-start)
- [云服务器停机](#vm-stop)

**云硬盘管理**

- [云硬盘创建并挂载](#storage-create-mount)
- [云硬盘卸载并销毁](#storage-umount-delete)


**负载均衡**

- [弹性负载均衡器创建](#loadbalancer-create)
- [弹性负载均衡器销毁](#loadbalancer-delete)
- [添加后端主机到负载均衡器](#loadbalancer-add-target)
- [从负载均衡器删除后端主机](#loadbalancer-delete-target)

**NAT网关**

- [NAT网关创建](#nat-gateway-create)
- [NAT网关销毁](#nat-gateway-terminate)
- [NAT网关添加snat规则](#nat-gateway-add-snat-rule)
- [NAT网关删除snat规则](#nat-gateway-delete-snat-rule)

**对等连接**

- [对等连接创建](#peering-connection-create)
- [对等连接销毁](#peering-connection-terminate)

**弹性公网IP**

- [弹性公网IP创建](#public-ip-create)
- [弹性公网IP销毁](#public-ip-terminate)

**路由策略**

- [路由策略创建](#route-policy-create)
- [路由策略销毁](#route-policy-terminate)

**安全组**

- [安全组创建](#security-group-create)
- [安全组销毁](#security-group-terminate)
- [安全组规则添加](#security-group-policy-create)
- [安全组规则删除](#security-group-ploicy-delete)

**云数据库RDS**

- [rds创建](#rds-create)
- [rds销毁](#rds-delete)
- [rds创建备份](#rds-create-backup)
- [rds销毁备份](#rds-delete-backup)

## API 概览及实例：  

### 私有网络

#### <span id="vpc-create">虚拟私有云创建</span>
[POST] /huaweicloud/v1/vpc/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|VPC实例ID，若有值，则会检查该VPC是否已存在 若已存在， 则不创建
name|string|否|VPC名称
cidr_block|string|是|VPC网段
enterprise_project_id|string|否|资源所属huaweicloud project

##### 输出参数：
参数名称|类型|描述
:--|:--|:--    
guid|string|CI类型全局唯一ID
id|string|VPC实例ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/vpc/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
       "Inputs":[
        {
          "guid":"0010_000000010",
          "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
          "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1; ProjectId=07b04b0a66000f092f6ec00f79a087c6",
          "name":"test_vpc",
          "cidr":"192.x.x.x/16"
       }
    ]
 }'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010",
                "id": "bb131fb8-b9fd-4fd7-a668-7a289c996fc1"
            }
        ]
    }
}
```


#### <span id="vpc-delete">虚拟私有云销毁</span>
[POST] /huaweicloud/v1/vpc/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|VPC实例ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--    
guid|string|CI类型全局唯一ID

##### 示例：
输入：

```
  curl -X POST http://127.0.0.1:8083/huaweicloud/v1/vpc/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
        "inputs":[
        {
            "guid":"0010_000000010",
            "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
            "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
            "id":"d19d6b34-67aa-43ff-943b-3d0b35888110"
        }
    ]
 }'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010"
            }
        ]
    }
}
```


### 子网

#### <span id="subnet-create">子网创建</span>
[POST] /huaweicloud/v1/subnet/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|子网实例ID，若有值，则会检查该子网是否已存在， 若已存在， 则不创建
name|string|是|子网名称
vpc_id|string|是|VPC实例ID
cidr_block|string|是|子网网段

##### 输出参数：
参数名称|类型|描述
:--|:--|:--    
guid|string|CI类型全局唯一ID
id|string|子网实例ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/subnet/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
        "inputs":[
        {
            "guid":"0010_000000010",
            "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
            "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
            "vpc_id":"beaa9d21-d5a7-4970-84a0-2de5b0240b69",
            "name": "test_subnet",
            "cidr":"192.x.x.x/24"
        }
    ]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010",
                "id": "3d5b766c-bd5b-472e-b14b-815b3b64a472"
            }
        ]
    }
}
```


#### <span id="subnet-terminate">子网销毁</span>
[POST] /huaweicloud/v1/subnet/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|子网实例ID
vpc_id|string|是|子网所属VPC实例ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/subnet/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
        "inputs":[
        {
            "guid":"0010_000000010",
            "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
            "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
            "id":"2e0dd197-082b-4a50-85d3-1c45c58d224f",
            "vpc_id":"beaa9d21-d5a7-4970-84a0-2de5b0240b69"
        }
    ]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010"
            }
        ]
    }
}
```

### 云服务器

#### <span id="vm-create">云服务器创建</span>
[POST] /huaweicloud/v1/vm/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|云服务器实例ID，若有值，则会检查该云服务器是否已存在， 若已存在， 则不创建
seed|string|是|云服务器密码加密种子
vpc_id|string|是|VPC实例ID
subnet_id|string|是|子网实例ID
image_id|string|是|虚拟机要安装的操作系统镜像ID
machine_spec|string|是|机器规格，使用2c2g的格式，插件后端会根据输入自动查找最匹配的机型
system_disk_size|string|是|系统盘大小，单位为G
password|string|否|虚拟机密码，如果不设置，插件后端会生成随机密码
az|string|是|虚拟机所属可用区
security_groups|string|否|虚拟机关联的安全组
charge_type|string|是|付费方式，支持按量计费和包年包月,可选值为prePaid和postPaid
period_type|string|否|包年包月时需指定，可选值为month和year
period_num|string|否|当period_type为month时，表示多少个月;period_type为year表示几年
is_auto_renew|string|否|包年包月时需指定，是否自动续费
name|string|是|云服务器实例名称
labels|string|否|云服务器的标签
private_ip|string|否|如果指定该参数，创建的vm将使用该ip作为局域网ip地址

##### 输出参数：
参数名称|类型|描述
:--|:--|:--    
request_id|string|请求ID
guid|string|CI类型全局唯一ID
id|string|云服务器实例ID
cpu|string|云服务器CPU核数
memory|string|云服务器内存大小
password|string|云服务器root密码，该密码为加密后的密码
private_ip|string|是|内网IP

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/vm/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
     "inputs": [
 	    {
            "guid":"0010_000000010",
            "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
            "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
            "seed":"seed-001",
            "image_id":"7077ec61-7553-4890-8b33-364005a590b9",
            "machine_spec":"1C1G",
            "system_disk_size":"40",
            "password":"Abcd1234",
            "az":"cn-south-1c",
            "security_groups":"c92366ef-4650-4496-8843-4bd02ce3910d",
            "charge_type":"postPaid",
            "name":"tyler_test_vm",
            "labels": "aa=bb;key=v"
        }
    ]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010",
                "id": "be31d19d-2e2a-43d1-a4fc-430a07b68f14",
                "cpu": "1",
                "memory": "1",
                "password": "{cipher_a}459df6cbd84dc63dbc1270499f3812ba",
                "private_ip": "192.x.x.x"
            }
        ]
    }
} 
```

#### <span id="vm-terminate">云服务器销毁</span>
[POST] /huaweicloud/v1/vm/terminate

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|云服务器实例ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/vm/terminate \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    {
        "identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
        "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
	    "guid": "1234",
        "id":"be31d19d-2e2a-43d1-a4fc-430a07b68f14"
	}
  ]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234"
            }
        ]
    }
} 
```

#### <span id="vm-start">云服务器启动</span>
[POST] /huaweicloud/v1/vm/start

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|云服务器实例ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID


##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/vm/start\
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
        {
            "identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
            "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
            "guid": "1234",
            "id":"be31d19d-2e2a-43d1-a4fc-430a07b68f14"
        }
    ]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234"
            }
        ]
    }
}
```

#### <span id="vm-stop">云服务器停机</span>
[POST] /huaweicloud/v1/vm/stop

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|云服务器实例ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID


##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/vm/stop \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
       {
            "identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
            "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
            "guid": "1234",
            "id":"be31d19d-2e2a-43d1-a4fc-430a07b68f14"
        }
    ]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0008_0000000088"
            }
        ]
    }
} 
```


### 云硬盘

#### <span id="storage-create-mount">云硬盘创建并挂载</span>
[POST] /huaweicloud/v1/block-storage/create-mount

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|云硬盘实例ID，若有值，则会检查该云硬盘是否已存在， 若已存在， 则不创建
az|string|是|云硬盘所属可用区
disk_type|string|是|云硬盘类型，可选值为SATA,SSD和SAS
disk_size|string|是|云硬盘大小，单位为GB
instance_id|string|是|需要挂载云硬盘的云服务器实例ID
instance_guid|string|是|云服务器实例在wecmdb中的guid
seed|string|是|云服务器密码加密用的种子，解密时需要使用
password|string|是|云服务器加密后的密码
file_system_type|string|是|云盘挂载到主机上格式化的文件系统，目前支持ext3,ext4和xfs
mount_dir|string|是|云硬盘挂载到主机的目录

##### 输出参数：
参数名称|类型|描述
:--|:--|:--    
guid|string|CI类型全局唯一ID
id|string|云硬盘实例ID
volume_name|string|云硬盘在主机上的卷名，格式如/dev/vdb
attach_id|string|云硬盘挂载到主机的id

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/block-storage/create-mount \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
            "identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
            "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
            "guid": "1234",
            "az":"cn-south-1c",
            "disk_type":"SATA",
            "disk_size":"30",
            "instance_id": "2cc3eeae-c4f3-4bc6-b5ee-2b5065c9870e",
            "instance_guid": "0010_000000010",
            "seed": "seed-001",
            "password":"{cipher_a}459df6cbd84dc63dbc1270499f3812ba",
            "file_system_type":"ext4",
            "mount_dir": "/data/test"
        }
    ]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234",
                "volume_name": "/dev/vdc",
                "id": "f66fe20e-9241-4544-b1be-1ba7f5773a12",
                "attach_id": "f66fe20e-9241-4544-b1be-1ba7f5773a12"
            }
        ]
    }
}  
```


#### <span id="storage-umount-delete">云硬盘卸载并销毁</span>
[POST] /huaweicloud/v1/block-storage/umount-delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|云硬盘实例ID
attach_id|string|是|云硬盘挂载到主机的id
az|string|是|云硬盘所属可用区
instance_id|string|是|需要挂载云硬盘的云服务器实例ID
instance_guid|string|是|云服务器实例在cmdb中的guid
seed|string|是|云服务器密码加密时用的种子，解密时需要使用
password|string|是|云服务器的加密后的密码
mount_dir|string|是|云硬盘挂载到主机的目录
volume_name|string|是|云盘的卷名称

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/block-storage/umount-delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
            "identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
            "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
            "guid": "1234",
            "az":"cn-south-1c",
            "attach_id":"f66fe20e-9241-4544-b1be-1ba7f5773a12",
            "id": "f66fe20e-9241-4544-b1be-1ba7f5773a12",
            "instance_id": "2cc3eeae-c4f3-4bc6-b5ee-2b5065c9870e",
            "instance_guid": "0010_000000010",
            "seed": "seed-001",
            "password":"{cipher_a}459df6cbd84dc63dbc1270499f3812ba",
            "volume_name":"/dev/vdc",
            "mount_dir": "/data/test"
        }
    ]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234"
            }
        ]
    }
}
```


### 弹性负载均衡器

#### <span id="loadbalancer-create">弹性负载均衡器创建</span>
[POST] /huaweicloud/v1/lb/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|可选值，如果该值对应的负载均衡器还存在，将不会新建
name|string|是|负载均衡器的名称
type|string|是|负载均衡器类型，可选值为Internal和External
subnet_id|string|是|负载均衡器的所属子网ID
bandwidth_size|string|否|当时外网负载均衡器时通过该参数需指定公网带宽大小

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
vip|string|负载均衡器的VIP
guid|string|CI类型全局唯一ID


##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/lb/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
			"identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
			"cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
			"guid": "1234",
			"name":"tyler-test-lb1",
			"type":"External",
			"subnet_id": "2de6e922-d2f8-4d92-8d3f-484d81a03d16",
			"bandwidth_size": "10"
		}]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234",
                "id": "8b9f1d04-992f-42b9-a658-b556210e641a",
                "vip": "121.37.9.10"
            }
        ]
    }
}
```


#### <span id="loadbalancer-delete">弹性负载均衡器删除</span>
[POST] /huaweicloud/v1/lb/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|负载均衡器ID
type|string|是|负载均衡器类型，可选值为Internal和External


##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID


##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/lb/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
			"identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
			"cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
			"guid": "1234",
			"id":"8b9f1d04-992f-42b9-a658-b556210e641a",
			"type":"External"
		}]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234"
            }
        ]
    }
}
```

#### <span id="loadbalancer-add-target">弹性负载均衡器添加后端主机</span>

[POST] /huaweicloud/v1/lb/add-backtarget

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
lb_id|string|是|负载均衡器ID
lb_port|string|是|负载均衡器上开放的端口
protocol|string|是|负载均衡器上开放的协议，只支持tcp和udp
host_ids|string|是|需要添加到负载均衡器后端的虚机id
host_ports|string|是|需要添加到负载均衡器后端的虚机开放的主机端口

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID


##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/lb/add-backtarget \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
			"identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
			 "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
			"guid": "1234",
			"lb_id":"2e7f0bc6-ec49-41ba-ab40-37152cff814e",
			"lb_port":"8080",
			"protocol":"tcp",
			"host_ids": "[2cc3eeae-c4f3-4bc6-b5ee-2b5065c9870e,7c811c94-1053-4555-84ca-b343f101eb73]",
			"host_ports": "[8888,9999]"
		}]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234"
            }
        ]
    }
}
```



#### <span id="loadbalancer-delete-target">弹性负载均衡器删除后端主机</span>

[POST] /huaweicloud/v1/lb/del-backtarget

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
lb_id|string|是|负载均衡器ID
lb_port|string|是|负载均衡器上开放的端口
protocol|string|是|负载均衡器上开放的协议，只支持tcp和udp
host_ids|string|是|需要添加到负载均衡器后端的虚机id
host_ports|string|是|需要添加到负载均衡器后端的虚机开放的主机端口

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/lb/del-backtarget \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
			"identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
			 "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
			"guid": "1234",
			"lb_id":"2e7f0bc6-ec49-41ba-ab40-37152cff814e",
			"lb_port":"8080",
			"protocol":"tcp",
			"host_ids": "[2cc3eeae-c4f3-4bc6-b5ee-2b5065c9870e,7c811c94-1053-4555-84ca-b343f101eb73]",
			"host_ports": "[8888,9999]"
		}]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234"
            }
        ]
    }
}
```


### NAT网关

#### <span id="nat-gateway-create">NAT网关创建</span>

[POST] /huaweicloud/v1/nat-gateway/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|nat网关id，如果已经存在就不会重复新建
vpc_id|string|是|nat网关所属vpc
subnet_id|string|是|nat网关给vpc下的哪个子网用

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID
id|string|创建成功的NAT网关ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/nat-gateway/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
		"guid":"0010_000000010",
     	"identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
        "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
        "vpc_id":"209e670a-95e2-4e73-81f8-7f931e5847a1",
        "subnet_id":"2de6e922-d2f8-4d92-8d3f-484d81a03d16"
		}]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010",
                "id": "2fa43ad2-6e82-45fd-b040-8e03670f492b"
            }
        ]
    }
}
```


#### <span id="nat-gateway-terminate">NAT网关销毁</span>

[POST] /huaweicloud/v1/nat-gateway/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|nat网关id

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID


##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/nat-gateway/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
		"guid":"0010_000000010",
  	    "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
    "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
    "id":"2fa43ad2-6e82-45fd-b040-8e03670f492b"
		}]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010",
            }
        ]
    }
}
```



#### <span id="nat-gateway-add-snat-rule">NAT网关添加snat规则</span>

[POST] /huaweicloud/v1/nat-gateway/add-snat-rule

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|nat网关snat规则ID,如果已经存在将不再创建
gateway_id|string|是|NAT网关ID
subnet_id|string|是|需要使用该NAT网关的子网ID
public_ip_id|sting|是|弹性公网IP的ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID
id|string|创建成功的NAT网关规则ID


##### 示例：

输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/add-snat-rule \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
		"guid":"0010_000000010",
     	"identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
      "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
     "gateway_id":"2fa43ad2-6e82-45fd-b040-8e03670f492b",
     "subnet_id":"2de6e922-d2f8-4d92-8d3f-484d81a03d16",
     "public_ip_id": "56f76de6-ac4e-4221-8996-8b30128a98c2"
		}]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
               "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010",
                "id": "e477cdad-bc3b-4076-8609-e337b1162dbb"
            }
        ]
    }
}
```


#### <span id="nat-gateway-delete-snat-rule">NAT网关删除snat规则</span>

[POST] /huaweicloud/v1/nat-gateway/delete-snat-rule

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|nat网关snat规则ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：

输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/add-snat-rule \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
   "inputs": [
	   {
		"guid":"0010_000000010",
     	"identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
      "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
     "id":"e477cdad-bc3b-4076-8609-e337b1162dbb"
		}]
}'
```

输出：

```
{
    "resultCode": "0",
    "resultMessage": "success",
    "results": {
        "outputs": [
            {
               "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010",
            }
        ]
    }
}
```

### 对等连接

#### <span id="peering-connection-create">对等连接创建</span>
[POST] /huaweicloud/v1/peerings/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|如果对等连接存在，将不会重复创建
local_vpc_id|string|是|一端VPC ID
peer_vpc_id|string|是|另外一端VPC ID


##### 输出参数：
参数名称|类型|描述
:--|:--|:--    
request_id|string|请求ID
guid|string|CI类型全局唯一ID
id|string|对等连接实例ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/peerings/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		"guid":"0010_000000010",
  	     "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
       "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
       "id":"cc1a1fbc-baf7-4a55-9671-c9c70babf948",
       "local_vpc_id":"209e670a-95e2-4e73-81f8-7f931e5847a1",
        "peer_vpc_id": "2d470aeb-34e5-4b22-bb51-cf1a9e7d1a75"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010",
                "id": "31201b76-7942-4ed9-9c6c-9db885161af3"
            }
        ]
    }
} 
```


#### <span id="peering-connection-terminate">对等连接销毁</span>
[POST] /huaweicloud/v1/peerings/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|对等连接实例ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/peerings/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		"guid":"0010_000000010",
  	    "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
       "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
       "id":"31201b76-7942-4ed9-9c6c-9db885161af3"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
               "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010"
            }
        ]
    }
} 
```

### 弹性公网IP

#### <span id="public-ip-create">弹性公网IP创建</span>
[POST] /huaweicloud/v1/public-ip/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|如果该ID对应的弹性公网IP已经存在，将不会新建
band_width|string|是带宽大小

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID
id|string|弹性公网IP的ID
ip|string|弹性公网IP对应的ip地址

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/public-ip/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		  "guid":"0010_000000010",
  	       "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
           "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
           "band_width":"10"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010",
                "id": "82a63f93-f97c-45f0-86f3-7160e47e69bc",
                "ip": "121.37.253.201"
            }
        ]
    }
} 
```

#### <span id="public-ip-terminate">弹性公网IP销毁</span>
[POST] /huaweicloud/v1/public-ip/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|弹性公网IP对应的资源ID


##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/public-ip/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		  "guid":"0010_000000010",
  	       "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
           "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
            "id":"82a63f93-f97c-45f0-86f3-7160e47e69bc"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "0010_000000010"
            }
        ]
    }
} 
```

### 路由策略

#### <span id="route-policy-create">路由策略创建</span>
[POST] /huaweicloud/v1/route/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|如果该ID对应的路由策略已经存在，将不会新建
destination|string|是|路由策略对应的目标网段
type|string|是|路由下一跳对应的资源类型，目前只支持对等连接
nexthop|string|是|路由下一跳对应的资源id，目前只支持对等连接ID
vpc_id|string|是|路由策略对哪个VPC生效|

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID
id|string|路由策略对应的ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/route/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		  "guid":"0010_000000010",
  	       "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
           "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
           "destination":"192.168.0.0/16",
		   "nexthop":"077bcc8d-cb6b-4771-990d-31eed3e7a98a",
		   "type":"peering",
		   "vpc_id": "beaa9d21-d5a7-4970-84a0-2de5b0240b69"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234",
                "id": "66e02158-bad3-4b13-b43c-e893824fa501"
            }
        ]
    }
} 
```

#### <span id="route-policy-terminate">路由策略销毁</span>
[POST] /huaweicloud/v1/route/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|路由策略对应的ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/route/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		  "guid":"0010_000000010",
  	       "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
           "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
		   "id":"09b76b0c-e2a3-4374-ac3b-06da72b12e3d"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid":"0010_000000010"
            }
        ]
    }
} 
```

### 安全组

#### <span id="security-group-create">安全组创建</span>
[POST] /huaweicloud/v1/security-group/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|如果该ID对应的安全组已存在，将不会新建
name|string|是|安全组名称

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID
id|string|创建成功的安全组ID

##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/security-group/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		  "guid":"0010_000000010",
  	       "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
           "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
		   "name":"test"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid":"0010_000000010"，
                 "id": "2e0dd197-082b-4a50-85d3-1c45c58d224f"
            }
        ]
    }
} 
```

#### <span id="security-group-terminate">安全组销毁</span>
[POST] /huaweicloud/v1/security-group/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|安全组对应的ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID


##### 示例：
输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/security-group/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		  "guid":"0010_000000010",
  	       "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
           "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
		   "id":"2e0dd197-082b-4a50-85d3-1c45c58d224f"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid":"0010_000000010"
            }
        ]
    }
} 
```



#### <span id="security-group-policy-create">安全组规则添加</span>
[POST] /huaweicloud/v1/security-group-rule/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|安全策略对应的ID，如果已存在将不会新建
security_group_id|string|是|安全策略所属安全组ID
direction|string|是|可选值为egress和ingress表示入站和出站规则
port_range_min|string|是|开放的端口最小值
port_range_max|string|是|开放的端口最大值
remote_ip_prefix|string|安全组放通的IP网段


##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID
id|string|安全组规则ID

##### 示例：

输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/security-group-rule/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		  "guid":"0010_000000010",
  	       "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
           "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
		   "security_group_id":"43080ac1-b3d9-4cbf-92b3-402a6ca67085",
           "direction": "egress",
           "protocol": "tcp",
           "port_range_min":"8080",
           "port_range_max":"8088",
           "remote_ip_prefix":"10.4.0.0/20"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid":"0010_000000010",
                "id":"cc1a1fbc-baf7-4a55-9671-c9c70babf948"
            }
        ]
    }
} 
```

#### <span id="security-group-ploicy-delete">安全组规则删除</span>
[POST] huaweicloud/v1/security-group-rule/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|安全策略对应的ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：

输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/security-group-rule/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		  "guid":"0010_000000010",
  	       "identity_params": "SecretKey=xxx;AccessKey=xxx;DomainId=xxx",
           "cloud_params":"CloudApiDomainName=myhuaweicloud.com;Region=cn-south-1;ProjectId=07b04b0a66000f092f6ec00f79a087c6",
           "id":"cc1a1fbc-baf7-4a55-9671-c9c70babf948"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid":"0010_000000010"
            }
        ]
    }
} 
```

### 云数据库RDS

#### <span id="rds-create">rds创建</span>
[POST] /huaweicloud/v1/rds/create

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|rds实例对应的ID,如果该ID对应的实例已存在将不会重新创建
seed|string|是|加密rds实例密码时使用
name|string|是|数据库实例的名称
password|string|否|如果没指定该值，将产生随机密码
port|string|否|数据库实例的端口
machine_spec|string|是|创建数据库实例的配置，支持2C2G这样的格式
engine_type|string|是|创建的数据库类型，可选值为MYSQL,PostgreSQL和SQLServer
engine_version|string|是|创建的数据库版本，如MYSQL 5.7
security_group_id|string|否|数据库关联的安全组
vpc_id|string|是|数据库实例所在vpc
subnet_id|string|是|数据库实例所在子网
az|string|是|数据库实例所在可用区
volume_type|string|是|数据库实例的存储媒介，可以是SSD等
volume_size|string|是|数据库实例的存储大小，单位为G
charge_type|string|是|付费方式，支持 prePaid(包年包月) 和 postPaid(按量计费)
period_type|string|否|包年包月时必填，可选值为month和year
period_num |string|否|包年包月时必填，表示买多久
is_auto_renew|string|否|包年包月时必填，是否自动续费

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID
id|string|创建成功的数据库实例ID
private_ip|string|创建成功的数据库实例内网IP
user_name|string|创建成功的数据库管理员账号名
password|string|加密后的数据库密码

##### 示例：

输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/rds/create \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		 	"identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
			 "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
			"guid": "1234",
			"seed":"seed-01",
			"name":"test-mysql",
			"password":"Abcd1234",
			"port": "3306",
			"machine_spec": "1C1G",
			"engine_type": "MySQL",
			"engine_version": "5.7",
			"security_group_id":"43080ac1-b3d9-4cbf-92b3-402a6ca67085",
			"vpc_id":"209e670a-95e2-4e73-81f8-7f931e5847a1",
			"subnet_id":"5c7916bd-f84b-4a78-b6d4-5a723f3e669f",
			"az":"cn-south-1c",
			"volume_type":"ULTRAHIGH",
			"volume_size":"40",
			"charge_type":"postPaid"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234",
                "id": "ae174276008a499eaaf28799e6fa7d52in01",
                "private_ip": "192.168.2.192",
                "private_port": "3306",
                "user_name": "root",
                "password": "{cipher_a}6b9e44e2ebb38170bb0da8838cae4e0d"
            }
        ]
    }
} 
```


#### <span id="rds-delete">rds销毁</span>
[POST] /huaweicloud/v1/security-group/delete

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|rds实例对应的ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：

输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/rds/delete \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		 	"identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
			 "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
			"guid": "1234",
			"id":"ae174276008a499eaaf28799e6fa7d52in01"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234",
            }
        ]
    }
} 
```


#### <span id="rds-create-backup">rds创建备份</span>
[POST] /huaweicloud/v1/rds/create-backup

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|否|数据库实例备份实例ID，如果已存在，将不会创建备份
instance_id|string|是|要备份的数据库实例ID


##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID
id|string|数据库备份实例ID


##### 示例：

输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/rds/create-backup \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		 	"identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
			 "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
			"guid": "1234",
			"instance_id": "ae174276008a499eaaf28799e6fa7d52in01"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234",
                "id": "7eb053fc1ba446858a6f811920fa5239br01"
            }
        ]
    }
} 
```


#### <span id="rds-delete-backup">rds销毁备份</span>
[POST] /huaweicloud/v1/rds/delete-backup

##### 输入参数：
参数名称|类型|必选|描述
:--|:--|:--|:-- 
guid|string|是|CI类型全局唯一ID
identity_params|string|是|公有云用户鉴权参数， 包括access-key，secret-key和domain-id
cloud_param|string|是|云api相关参数，包括云API域名，region和project-id
id|string|是|数据库实例备份实例ID

##### 输出参数：
参数名称|类型|描述
:--|:--|:--
guid|string|CI类型全局唯一ID

##### 示例：

输入：

```
curl -X POST http://127.0.0.1:8083/huaweicloud/v1/rds/delete-backup \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "inputs": [
    	{
		 	"identity_params":"AccessKey=xxx;SecretKey=xxx;DomainId=xxx",
			 "cloud_params":"CloudApiDomainName=myhuaweicloud.com;ProjectId=07b04b0a66000f092f6ec00f79a087c6;Region=cn-south-1",
			"guid": "1234",
			"id":"ae174276008a499eaaf28799e6fa7d52in01"
		}
	]
}'
```

输出：

```
{
    "result_code": "0",
    "result_message": "success",
    "results": {
        "outputs": [
            {
                "errorCode": "0",
                "errorMessage": "",
                "guid": "1234"
            }
        ]
    }
} 
```
