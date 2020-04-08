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
        <plugin name="vpc" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/vpc/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">cidr</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/vpc/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="security-group" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/security-group/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">vpc_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/security-group/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="subnet" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/subnet/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cidr</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">vpc_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">az</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/subnet/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">vpc_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="vm" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/vm/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">seed</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">image_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">machine_spec</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">system_disk_size</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">vpc_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">subnet_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">private_ip</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">password</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">labels</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">az</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">security_group</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">charge_type</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">period_type</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">period_num</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">is_auto_renew</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">cpu</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">memory</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">private_ip</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="terminate" path="/huaweicloud/v1/vm/terminate">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="stop" path="/huaweicloud/v1/vm/stop">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="start" path="/huaweicloud/v1/vm/start">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="lb" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/lb/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">type</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">subnet_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">bandwidth_size</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">vip</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/lb/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">type</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="lb-target" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/lb-target/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">lb_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">lb_port</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">listener_name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">protocol</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">host_ids</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">host_ports</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">listener_id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/lb-target/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">lb_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">lb_port</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">protocol</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">host_ids</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">host_ports</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="block-storage" targetPackage="" targetEntity="">
            <interface action="create-mount" path="/huaweicloud/v1/block-storage/create-mount">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">az</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">disk_type</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">disk_size</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">instance_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">instance_guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">seed</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">password</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">mount_dir</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">file_system_type</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">volume_name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <!--<parameter datatype="string" mappingType="entity" mappingEntityExpression="">attach_id</parameter> -->
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="umount-delete" path="/huaweicloud/v1/block-storage/umount-delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">az</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                    <!--   <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">attach_id</parameter> -->
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">instance_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">instance_guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">seed</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">password</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">mount_dir</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">volume_name</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="security-group-rule" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/security-group-rule/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">security_group_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">direction</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">protocol</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">port_range_min</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">port_range_max</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">remote_ip_prefix</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">enterprise_project_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/security-group-rule/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="peerings" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/peerings/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">local_vpc_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">peer_vpc_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/peerings/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="public-ip" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/public-ip/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">band_width</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">ip</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/public-ip/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="nat-gateway" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/nat-gateway/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">vpc_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">subnet_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">Name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/nat-gateway/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
         <plugin name="nat-snat-rule" targetPackage="" targetEntity="">
            <interface action="add" path="/huaweicloud/v1/nat-snat-rule/add">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">gateway_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">subnet_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">public_ip_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/nat-snat-rule/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="route" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/route/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">destination</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">nexthop</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">type</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">vpc_id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/route/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
        <plugin name="rds" targetPackage="" targetEntity="">
            <interface action="create" path="/huaweicloud/v1/rds/create">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">seed</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">password</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">port</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">machine_spec</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">support_ha</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">ha_replication_mode</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">engine_version</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">security_group_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">vpc_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">subnet_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">az</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">volume_type</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">volume_size</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">charge_type</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">period_type</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">period_num</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">is_auto_renew</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">enterprise_project_id</parameter>
                    <parameter datatype="string" mappingType="system_variable" mappingSystemVariableName="" required="Y">character_set</parameter>
                    <parameter datatype="string" mappingType="system_variable" mappingSystemVariableName="" required="Y">lower_case_table_names</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">private_ip</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">private_port</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">user_name</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">password</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">cpu</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">memory</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete" path="/huaweicloud/v1/rds/delete">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="create-backup" path="/huaweicloud/v1/rds/create-backup">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="N">id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">instance_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">name</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">id</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
            <interface action="delete-backup" path="/huaweicloud/v1/rds/delete-backup">
                <inputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">guid</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">identity_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">cloud_params</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">instance_id</parameter>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="" required="Y">id</parameter>
                </inputParameters>
                <outputParameters>
                    <parameter datatype="string" mappingType="entity" mappingEntityExpression="">guid</parameter>
                    <parameter datatype="string" mappingType="context">errorCode</parameter>
                    <parameter datatype="string" mappingType="context">errorMessage</parameter>
                </outputParameters>
            </interface>
        </plugin>
    </plugins>
</package>
