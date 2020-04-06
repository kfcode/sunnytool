参考 https://blog.csdn.net/weixin_38249179/article/details/105332346



//client端的代码生成形式如下，其中p=cli 表示生成client端代码，pkg 表示client 端的包名，sub=sunnycode, 表示子目录， p=cli的情况下必填，

生成的client的代码目录形式为(假如pkg=usercli， sub=sunnytest)：
 github.com/kfcode/clients
                        --sunnytest
                            --usercli
                              usercli.go
                              user.pb.go
                             
protoc --plugin=/Users/simon/go/bin/protoc-gen-sunnytool --sunnytool_out=plugins=rpcx,p=cli,pkg=user_cli,step=cli:. ./user.proto

//svr端的代码生成形式如下，其中p=svr 表示生成svr端代码, 代码生成默认放到当前工程的proto 目录中，不需要pkg 参数，有也忽略。
protoc --plugin=/Users/simon/go/bin/protoc-gen-sunnytool --sunnytool_out=plugins=rpcx,p=svr,step=svr:. ./user.proto

//在当前user目录执行生成svr代码如下
 github.com/kfcode/user
                        --proto
                           user.pb.go
                        handler.go
                        service.go
                        sunny.go
                        
                              
                             
