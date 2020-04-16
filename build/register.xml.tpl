<?xml version="1.0" encoding="UTF-8"?>
<package name="huaweicloud" version="{{PLUGIN_VERSION}}">
    <!-- 1.依赖分析 - 描述运行本插件包需要的其他插件包 -->
    <packageDependencies>
    </packageDependencies>

    <!-- 2.菜单注入 - 描述运行本插件包需要注入的菜单 -->
    <menus>
    </menus>

    <!-- 3.数据模型 - 描述本插件包的数据模型,并且描述和Framework数据模型的关系 -->
    <dataModel>
    </dataModel>

    <!-- 4.系统参数 - 描述运行本插件包需要的系统参数 -->
    <systemParameters>
        <systemParameter name="HWCLOUD_MYSQL_BACKUP_TYPE" scopeType="global" defaultValue="logical"/>
        <systemParameter name="HWCLOUD_MYSQL_CHARACTER_SET" scopeType="global" defaultValue="UTF8"/>
        <systemParameter name="HWCLOUD_MYSQL_LOWER_CASE_TABLE_NAMES" scopeType="global" defaultValue="0"/>
	    <systemParameter name="HWCLOUD_API_SECRET" scopeType="global" defaultValue="SecretKey=;AccessKey=;DomainId="/>
        <systemParameter name="HWCLOUD_PERIOD_TYPE" scopeType="global" defaultValue="month"/>
        <systemParameter name="HWCLOUD_DELETE_LB_LISTENER" scopeType="global" defaultValue="Y"/>
        <systemParameter name="HWCLOUD_NOT_DELETE_LB_LISTENER" scopeType="global" defaultValue="N"/>
        <systemParameter name="HWCLOUD_IS_AUTO_RENEW" scopeType="global" defaultValue="N"/>
        <systemParameter name="HWCLOUD_PRIMARY_DNS" scopeType="global" defaultValue="127.0.0.1"/>
        <systemParameter name="HWCLOUD_SECONDARY_DNS" scopeType="global" defaultValue="127.0.0.1"/>
    </systemParameters>

    <!-- 5.权限设定 -->
    <authorities>
    </authorities>

    <!-- 6.运行资源 - 描述部署运行本插件包需要的基础资源(如主机、虚拟机、容器、数据库等) -->
    <resourceDependencies>
        <docker imageName="{{IMAGENAME}}" containerName="{{CONTAINERNAME}}" portBindings="{{PORTBINDINGS}}" volumeBindings="/etc/localtime:/etc/localtime,{{BASE_MOUNT_PATH}}/huaweicloud/logs:/home/app/huaweicloud/logs" envVariables=""/>
    </resourceDependencies>

    <!-- 7.插件列表 - 描述插件包中单个插件的输入和输出 -->
    <plugins>
        <plugin name="vpc" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/vpc/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cidr</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/vpc/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="security-group" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/security-group/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">vpc_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/security-group/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="subnet" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/subnet/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cidr</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">vpc_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_PRIMARY_DNS">primary_dns</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_SECONDARY_DNS">secondary_dns</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/subnet/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">vpc_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="vm" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/vm/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="ENCRYPT_SEED">seed</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">image_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">machine_spec</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">system_disk_type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">system_disk_size</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">vpc_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">subnet_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">private_ip</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="Y" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">labels</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">security_group</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">charge_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_PERIOD_TYPE">period_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">period_num</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_IS_AUTO_RENEW">is_auto_renew</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cpu</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">memory</parameter>
                    <parameter datatype="string" sensitiveData="Y" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">private_ip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="terminate" path="/huaweicloud/v1/vm/terminate" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="stop" path="/huaweicloud/v1/vm/stop" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="start" path="/huaweicloud/v1/vm/start" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="lb" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/lb/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">subnet_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">bandwidth_size</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">vip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/lb/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">type</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="lb-target" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/lb-target/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">lb_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">lb_port</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">listener_name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">listener_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">protocol</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">host_ids</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">host_ports</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">listener_id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/lb-target/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">lb_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">listener_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">lb_port</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">protocol</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">delete_listener</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">host_ids</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">host_ports</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="block-storage" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create-mount" path="/huaweicloud/v1/block-storage/create-mount" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">disk_type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">disk_size</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">instance_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">instance_guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="ENCRYPT_SEED">seed</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">mount_dir</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">file_system_type</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">volume_name</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="umount-delete" path="/huaweicloud/v1/block-storage/umount-delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">az</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">instance_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">instance_guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="ENCRYPT_SEED">seed</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">mount_dir</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">volume_name</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="security-group-rule" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/security-group-rule/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">security_group_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">direction</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">protocol</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">port</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">remote_ip_prefix</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/security-group-rule/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="peerings" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/peerings/create"  filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">local_vpc_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">peer_vpc_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/peerings/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="public-ip" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/public-ip/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">band_width</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">ip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/public-ip/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="nat-gateway" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/nat-gateway/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">vpc_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">subnet_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/nat-gateway/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
         <plugin name="nat-snat-rule" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="add" path="/huaweicloud/v1/nat-snat-rule/add" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">gateway_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">subnet_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">public_ip_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/nat-snat-rule/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="route" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/route/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">destination</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">nexthop</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">vpc_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/route/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="rds" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/rds/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="ENCRYPT_SEED">seed</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="Y" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">port</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">machine_spec</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">support_ha</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">ha_replication_mode</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">engine_version</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">security_group_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">vpc_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">subnet_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">volume_type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">volume_size</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">charge_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_PERIOD_TYPE">period_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">period_num</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_IS_AUTO_RENEW">is_auto_renew</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">enterprise_project_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="">character_set</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="">lower_case_table_names</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">private_ip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">private_port</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">user_name</parameter>
                    <parameter datatype="string" sensitiveData="Y" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cpu</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">memory</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/rds/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="create-backup" path="/huaweicloud/v1/rds/create-backup" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">instance_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete-backup" path="/huaweicloud/v1/rds/delete-backup" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">instance_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>

        <plugin name="redis" targetPackage="" targetEntity="" registerName="" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/dcs/create" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">seed</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">instance_type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">engine_version</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">capacity</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">vpc_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">subnet_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">security_group_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">private_ip</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">port</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">charge_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">period_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">period_num</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="">is_auto_renew</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">private_ip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">port</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/dcs/delete" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>


        <!-- 配套WeCube最佳实践-->
        <plugin name="vpc" targetPackage="wecmdb" targetEntity="network_segment" registerName="vpc" targetEntityFilterRule="{network_segment_usage eq 'VPC'}">
            <interface action="create" path="/huaweicloud/v1/vpc/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.name">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.code">cidr</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.vpc_asset_id">id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.NONE">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.vpc_asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/vpc/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.vpc_asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="security-group" targetPackage="wecmdb" targetEntity="network_segment" registerName="vpc" targetEntityFilterRule="{network_segment_usage eq 'VPC'}">
            <interface action="create" path="/huaweicloud/v1/security-group/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.name">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.vpc_asset_id">vpc_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.security_group_asset_id">id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.NONE">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.security_group_asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/security-group/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.security_group_asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="subnet" targetPackage="wecmdb" targetEntity="network_segment" registerName="subnet" targetEntityFilterRule="{network_segment_usage eq 'SUBNET'}">
            <interface action="create" path="/huaweicloud/v1/subnet/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.name">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.code">cidr</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.subnet_asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.f_network_segment>wecmdb:network_segment.vpc_asset_id">vpc_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.available_zone">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_PRIMARY_DNS">primary_dns</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_SECONDARY_DNS">secondary_dns</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.subnet_asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/subnet/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.subnet_asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.f_network_segment>wecmdb:network_segment.vpc_asset_id">vpc_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="vm" targetPackage="wecmdb" targetEntity="host_resource_instance" registerName="resource" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/vm/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="ENCRYPT_SEED">seed</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.resource_instance_system>wecmdb:resource_instance_system.code">image_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.resource_instance_spec>wecmdb:resource_instance_spec.code">machine_spec</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.storage_type>wecmdb:storage_type.code">system_disk_type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.storage">system_disk_size</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.f_network_segment>wecmdb:network_segment.vpc_asset_id">vpc_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.subnet_asset_id">subnet_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.code">private_ip</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.key_name">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="Y" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.user_password">password</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.NONE">labels</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.available_zone">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.f_network_segment>wecmdb:network_segment.security_group_asset_id">security_group</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.charge_type>wecmdb:charge_type.code">charge_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_PERIOD_TYPE">period_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.billing_cycle">period_num</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_IS_AUTO_RENEW">is_auto_renew</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.NONE">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.cpu">cpu</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.memory">memory</parameter>
                    <parameter datatype="string" sensitiveData="Y" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.user_password">password</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.code">private_ip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="terminate" path="/huaweicloud/v1/vm/terminate" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="stop" path="/huaweicloud/v1/vm/stop" filterRule="{state_code eq 'stoped'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="start" path="/huaweicloud/v1/vm/start" filterRule="{state_code eq 'startup'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:host_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="lb" targetPackage="wecmdb" targetEntity="lb_resource_instance" registerName="resource" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/lb/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.ip_address>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.key_name">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.resource_instance_type>wecmdb:resource_instance_type.code">type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.ip_address>wecmdb:ip_address.network_segment>wecmdb:network_segment.subnet_asset_id">subnet_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.bandwidth_size">bandwidth_size</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.NONE">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.ip_address>wecmdb:ip_address.code">vip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/lb/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.ip_address>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.resource_instance_type>wecmdb:resource_instance_type.code">type</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="lb-target" targetPackage="wecmdb" targetEntity="lb_instance" registerName="whole" targetEntityFilterRule="">
            <interface action="add" path="/huaweicloud/v1/lb-target/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.lb_resource_instance>wecmdb:lb_resource_instance.ip_address>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.lb_resource_instance>wecmdb:lb_resource_instance.asset_id">lb_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.port">lb_port</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.name">listener_name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.asset_id">listener_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit.protocol">protocol</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit~(invoke_unit)wecmdb:invoke.invoked_unit>wecmdb:unit~(unit)wecmdb:app_instance.host_resource_instance>wecmdb:host_resource_instance.asset_id">host_ids</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit~(invoke_unit)wecmdb:invoke.invoked_unit>wecmdb:unit~(unit)wecmdb:app_instance.port">host_ports</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.asset_id">listener_id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/lb-target/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.lb_resource_instance>wecmdb:lb_resource_instance.ip_address>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.lb_resource_instance>wecmdb:lb_resource_instance.asset_id">lb_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.asset_id">listener_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.port">lb_port</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit.protocol">protocol</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_DELETE_LB_LISTENER">delete_listener</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit~(invoke_unit)wecmdb:invoke.invoked_unit>wecmdb:unit~(unit)wecmdb:app_instance.host_resource_instance>wecmdb:host_resource_instance.asset_id">host_ids</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit~(invoke_unit)wecmdb:invoke.invoked_unit>wecmdb:unit~(unit)wecmdb:app_instance.port">host_ports</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="lb-target" targetPackage="wecmdb" targetEntity="lb_instance" registerName="target" targetEntityFilterRule="">
            <interface action="add" path="/huaweicloud/v1/lb-target/create" filterRule="{state_code eq 'created'}{fixed_date isnot  NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.lb_resource_instance>wecmdb:lb_resource_instance.ip_address>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.lb_resource_instance>wecmdb:lb_resource_instance.asset_id">lb_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.port">lb_port</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.name">listener_name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.asset_id">listener_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit.protocol">protocol</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit~(invoke_unit)wecmdb:invoke.invoked_unit>wecmdb:unit~(unit)wecmdb:app_instance{state_code eq 'created'}{fixed_date is NULL}.host_resource_instance>wecmdb:host_resource_instance.asset_id">host_ids</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit~(invoke_unit)wecmdb:invoke.invoked_unit>wecmdb:unit~(unit)wecmdb:app_instance{state_code eq 'created'}{fixed_date is NULL}.port">host_ports</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.asset_id">listener_id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/lb-target/delete" filterRule="{state_code eq 'created'}{fixed_date isnot  NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.lb_resource_instance>wecmdb:lb_resource_instance.ip_address>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.lb_resource_instance>wecmdb:lb_resource_instance.asset_id">lb_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.asset_id">listener_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.port">lb_port</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit.protocol">protocol</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_NOT_DELETE_LB_LISTENER">delete_listener</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit~(invoke_unit)wecmdb:invoke.invoked_unit>wecmdb:unit~(unit)wecmdb:app_instance{state_code eq 'destroyed'}{fixed_date is NULL}.host_resource_instance>wecmdb:host_resource_instance.asset_id">host_ids</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.unit>wecmdb:unit~(invoke_unit)wecmdb:invoke.invoked_unit>wecmdb:unit~(unit)wecmdb:app_instance{state_code eq 'destroyed'}{fixed_date is NULL}.port">host_ports</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:lb_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="block-storage" targetPackage="wecmdb" targetEntity="block_storage" registerName="data_disk" targetEntityFilterRule="">
            <interface action="create-mount" path="/huaweicloud/v1/block-storage/create-mount" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.available_zone">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.name">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.storage_type>wecmdb:storage_type.code">disk_type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.disk_size">disk_size</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.asset_id">instance_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.guid">instance_guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="ENCRYPT_SEED">seed</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.user_password">password</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.mount_point">mount_dir</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.file_system">file_system_type</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.code">volume_name</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="umount-delete" path="/huaweicloud/v1/block-storage/umount-delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.available_zone">az</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.asset_id">instance_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.guid">instance_guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="ENCRYPT_SEED">seed</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.host_resource_instance>wecmdb:host_resource_instance.user_password">password</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.mount_point">mount_dir</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.code">volume_name</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:block_storage.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="security-group-rule" targetPackage="wecmdb" targetEntity="default_security_policy" registerName="resource_init" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/security-group-rule/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.owner_network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.owner_network_segment>wecmdb:network_segment.security_group_asset_id">security_group_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.security_policy_asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.security_policy_type">direction</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.protocol">protocol</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.port">port</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.policy_network_segment>wecmdb:network_segment.code">remote_ip_prefix</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.NONE">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.security_policy_asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/security-group-rule/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.owner_network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.security_policy_asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:default_security_policy.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="peerings" targetPackage="wecmdb" targetEntity="network_link" registerName="network_link" targetEntityFilterRule="{network_link_type eq '@@0054_0000000001@@华为云对等连接'}">
            <interface action="create" path="/huaweicloud/v1/peerings/create"  filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.network_segment_2>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.network_segment_2>wecmdb:network_segment.vpc_asset_id">local_vpc_id</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.name">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.network_segment_1>wecmdb:network_segment.vpc_asset_id">peer_vpc_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/peerings/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.network_segment_2>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="public-ip" targetPackage="wecmdb" targetEntity="ip_address" registerName="nat_ip" targetEntityFilterRule="{ip_address_usage eq '外网NAT'}">
            <interface action="create" path="/huaweicloud/v1/public-ip/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.network_segment>wecmdb:network_segment.name">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address~(internet_ip)wecmdb:network_link.netband_width">band_width</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.code">ip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/public-ip/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:ip_address.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="nat-gateway" targetPackage="wecmdb" targetEntity="network_link" registerName="network_link" targetEntityFilterRule="{network_link_type eq '@@0054_0000000002@@华为云NAT网关'}">
            <interface action="create" path="/huaweicloud/v1/nat-gateway/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.network_segment_2>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.network_segment_2>wecmdb:network_segment.vpc_asset_id">vpc_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.ip_address>wecmdb:ip_address.network_segment>wecmdb:network_segment.subnet_asset_id">subnet_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.name">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/nat-gateway/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.network_segment_2>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_link.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
         <plugin name="nat-snat-rule" targetPackage="wecmdb" targetEntity="network_segment" registerName="subnet" targetEntityFilterRule="{network_segment_usage eq 'SUBNET'}">
            <interface action="create-add" path="/huaweicloud/v1/nat-snat-rule/add" filterRule="{state_code eq 'created'}{fixed_date is NULL}{private_nat eq 'Y'}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.nat_rule_asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.f_network_segment>wecmdb:network_segment~(network_segment_2)wecmdb:network_link{network_link_type in ['@@0054_0000000002@@华为云NAT网关']}.asset_id">gateway_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.subnet_asset_id">subnet_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.f_network_segment>wecmdb:network_segment~(network_segment_2)wecmdb:network_link{network_link_type in ['@@0054_0000000002@@华为云NAT网关']}.internet_ip>wecmdb:ip_address.asset_id">public_ip_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.nat_rule_asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="change-add" path="/huaweicloud/v1/nat-snat-rule/add" filterRule="{state_code eq 'changed'}{fixed_date is NULL}{private_nat eq 'Y'}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.nat_rule_asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.f_network_segment>wecmdb:network_segment~(network_segment_2)wecmdb:network_link{network_link_type in ['@@0054_0000000002@@华为云NAT网关']}.asset_id">gateway_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.subnet_asset_id">subnet_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.f_network_segment>wecmdb:network_segment~(network_segment_2)wecmdb:network_link{network_link_type in ['@@0054_0000000002@@华为云NAT网关']}.internet_ip>wecmdb:ip_address.asset_id">public_ip_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.nat_rule_asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="destroy-delete" path="/huaweicloud/v1/nat-snat-rule/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}{private_nat eq 'Y'}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.nat_rule_asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="change-delete" path="/huaweicloud/v1/nat-snat-rule/delete" filterRule="{state_code eq 'changed'}{fixed_date is NULL}{private_nat eq 'N'}{nat_rule_asset_id neq ''}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.nat_rule_asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:network_segment.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="route" targetPackage="wecmdb" targetEntity="route" registerName="route" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/route/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.owner_network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.dest_network_segment>wecmdb:network_segment.code">destination</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.network_link>wecmdb:network_link.asset_id">nexthop</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.network_link>wecmdb:network_link.network_link_type>wecmdb:network_link_type.code">type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.owner_network_segment>wecmdb:network_segment.vpc_asset_id">vpc_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/route/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.owner_network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:route.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="rds" targetPackage="wecmdb" targetEntity="rdb_resource_instance" registerName="resource" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/rds/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="ENCRYPT_SEED">seed</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.key_name">name</parameter>
                    <parameter datatype="string" required="N" sensitiveData="Y" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.user_password">password</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.login_port">port</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.resource_instance_spec>wecmdb:resource_instance_spec.code">machine_spec</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.cluster_node_type>wecmdb:cluster_node_type.cluster_type>wecmdb:cluster_type.code">support_ha</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.resource_instance_type>wecmdb:resource_instance_type.code">ha_replication_mode</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.resource_instance_system>wecmdb:resource_instance_system.code">engine_version</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.f_network_segment>wecmdb:network_segment.security_group_asset_id">security_group_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.f_network_segment>wecmdb:network_segment.vpc_asset_id">vpc_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.subnet_asset_id">subnet_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.resource_set>wecmdb:resource_set.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.available_zone">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.storage_type>wecmdb:storage_type.code">volume_type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.storage">volume_size</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.charge_type>wecmdb:charge_type.code">charge_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_PERIOD_TYPE">period_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.billing_cycle">period_num</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_IS_AUTO_RENEW">is_auto_renew</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.NONE">enterprise_project_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_MYSQL_CHARACTER_SET">character_set</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_MYSQL_LOWER_CASE_TABLE_NAMES">lower_case_table_names</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.intranet_ip>wecmdb:ip_address.code">private_ip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.login_port">private_port</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.user_name">user_name</parameter>
                    <parameter datatype="string" sensitiveData="Y" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.user_password">password</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.cpu">cpu</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.memory">memory</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/rds/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="create-backup" path="/huaweicloud/v1/rds/create-backup" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.backup_asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.asset_id">instance_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.key_name">name</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.backup_asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete-backup" path="/huaweicloud/v1/rds/delete-backup" filterRule="">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="Y" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.asset_id">instance_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.backup_asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:rdb_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="redis" targetPackage="wecmdb" targetEntity="cache_resource_instance" registerName="resource" targetEntityFilterRule="">
            <interface action="create" path="/huaweicloud/v1/dcs/create" filterRule="{state_code eq 'created'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="ENCRYPT_SEED">seed</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.asset_id">id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.key_name">name</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.resource_instance_type>wecmdb:resource_instance_type.code">instance_type</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.resource_instance_system>wecmdb:resource_instance_system.code">engine_version</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.resource_instance_spec>wecmdb:resource_instance_spec.code">capacity</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.user_password">password</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.f_network_segment>wecmdb:network_segment.vpc_asset_id">vpc_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.subnet_asset_id">subnet_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.f_network_segment>wecmdb:network_segment.security_group_asset_id">security_group_id</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.resource_set>wecmdb:resource_set.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.available_zone">az</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.intranet_ip>wecmdb:ip_address.code">private_ip</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.login_port">port</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.charge_type>wecmdb:charge_type.code">charge_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_PERIOD_TYPE">period_type</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.billing_cycle">period_num</parameter>
                    <parameter datatype="string" required="N" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_IS_AUTO_RENEW">is_auto_renew</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.asset_id">id</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.user_password">password</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.intranet_ip>wecmdb:ip_address.code">private_ip</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.login_port">port</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/dcs/delete" filterRule="{state_code eq 'destroyed'}{fixed_date is NULL}">
                <inputParameters>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="system_variable" mappingSystemVariableName="HWCLOUD_API_SECRET">identity_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.intranet_ip>wecmdb:ip_address.network_segment>wecmdb:network_segment.data_center>wecmdb:data_center.location">cloud_params</parameter>
                    <parameter datatype="string" required="Y" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.asset_id">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" sensitiveData="N" mappingType="entity" mappingEntityExpression="wecmdb:cache_resource_instance.guid">guid</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" sensitiveData="N" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
    </plugins>
</package>
