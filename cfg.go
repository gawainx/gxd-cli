/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-May-28
 *
*/

package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"fmt"
    "encoding/json"
    "github.com/docker/go-connections/nat"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/api/types"
    "strings"
    "path/filepath"
    "os"
    "github.com/docker/docker/client"
    "context"
)

//PortInt: 端口号类型
type PortInt int

//PortInt.String() : 将端口号转换为字符串
//  返回值: 字符串格式的端口号
func (pi PortInt) String() string{
    return fmt.Sprintf("%d",pi)
}

//TOMLConfig: 储存从 TOML 读取的配置信息
//Title: 配置名称
//Net: 网络字段
//Service: 服务(对应于容器)
type TOMLConfig struct {
    Title           string
    Net             NetworkConfig    `toml:"network"`
    Service         ContainerConfig
}

// NetworkConfig: To create network
// Name: 网络的名称
type NetworkConfig struct {
    Name string
}

// NetworkConfig.Name: 返回创建网络的名称
// 返回值:网络的名称
func (n *NetworkConfig) String() string{
    return n.Name
}

// Config container
//    Priority      设定启动顺序
//    Name 			设定容器名
//    Image			string,设定镜像名
//    Detached 		bool,设定是否后台运行(不输出初始化日志记录),
//                  true 表示不输出初始化记录
//                  false 表示
//    WorkDir		设定容器工作目录
//    CMD           设定容器运行的命令
//    Net           配置容器的网络信息
//    Ports 		配置容器端口供外部访问
//                  支持挂载多端口
//    Volumes		配置挂载卷信息
//    AutoRemove    设定容器运行完毕后是否删除该容器
//                  true 表示自动删除
//                  false 表示不删除
type ContainerConfig struct {
    Priority        uint //launch order
    Name 			string
    Image			string
    Detached 		bool
    WorkDir			string `toml:"work_dir"`
    CMD             string `toml:"cmd"`
    Net             string
    Ports 			[]Port
    Volumes			Vols
    AutoRemove      bool   `toml:"auto_remove"`
}

//ContainerConfig.RunContainer: 从配置运行容器
//cli:  用于访问 docker 守护进程
//ctx:  传递本次操作的上下文信息
func (c *ContainerConfig) RunContainer(cli *client.Client, ctx context.Context){
    c.Volumes.ReplacePWD() //replace pwd to current abs dir
    //set mount volumes
    vols := make([]string,len(c.Volumes))
    for index,item := range c.Volumes{
        vols[index] = item.String()
    }

    //set exposed ports for containers and publish ports
    exports := make(nat.PortSet)
    pts := make(nat.PortMap)
    //配置端口映射数据结构
    for _,p := range c.Ports{
        tmpPort, _ := nat.NewPort("tcp",p.Target.String())
        pb := make([]nat.PortBinding,0)
        pb = append(pb,nat.PortBinding{
            HostPort:p.Host.String(),
        })
        exports[tmpPort] = struct{}{}
        pts[tmpPort] = pb
    }

    //创建容器
    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image:c.Image,
        ExposedPorts:exports,
        Cmd:strings.Split(c.CMD," "),
        WorkingDir:c.WorkDir,
    },&container.HostConfig{
        Binds:vols,
        PortBindings:pts,
        NetworkMode:container.NetworkMode(c.Net),
        AutoRemove:c.AutoRemove,
    },nil,c.Name)
    if err != nil{
        panic(err)
    }
    //遇到容器创建错误时发起 panic
    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }else{
        log.Printf("Container %s is created and started.\n",resp.ID)
    }
}

//将配置信息转换为 json 数据用于输出
//返回值: JSON 格式数据
//用于排查问题
func (c *ContainerConfig) JSONStr() string{
    res,e := json.Marshal(c)
    if e != nil{
        return ""
    }else{
        return string(res)
    }
}

//Port:端口映射信息数据
//Port.Host:宿主机端口
//Port.Target: 容器内部端口
type Port struct {
    Host			PortInt
    Target			PortInt
}
//Port.String: 输出端口映射配列
func (p *Port) String()string{
	return fmt.Sprintf("%d:%d",p.Host,p.Target)
}

//Vol: 设置卷映射
//Vol.Host: 宿主机文件夹
//Vol.Target: 目标容器文件夹
type Vol struct {
    Host 			string
    Target 			string
}
//Vol.String: 输出端口映射配列
func(v *Vol) String()string{
	return fmt.Sprintf("%s:%s",v.Host,v.Target)
}
//Vols: 储存多卷映射序列
type Vols []Vol
//ReplacePWD: 替换卷映射过程中的" pwd" 为当前工作目录
func (vs *Vols) ReplacePWD(){
    curDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    for i, v := range *vs{
        if strings.ToLower(v.Host[:3]) == "pwd"{
            //fmt.Println(v.Host[:3])
            (*vs)[i].Host = strings.Replace(v.Host,v.Host[:3],curDir,-1)
            //fmt.Println(v)
        }
    }
}

// Init toml from *.toml
//filename: 文件名信息
func (t *TOMLConfig) InitFromFile(filename string){
    _,e := toml.DecodeFile(filename,t)
    if e != nil{
    	log.Fatal(e)
    }
}

//func RunContainer(t *TOMLConfig){
//    t.Service.RunContainer()
//}
