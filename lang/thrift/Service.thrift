include "User.thrift"
// ./thrift-0.16.0 -r -gen go service.thrift
namespace go Sample

typedef map<string, string> Data

struct Response {
    1:required i32 errCode; //错误码
    2:required string errMsg; //错误信息
    3:required Data data;
}

//定义服务
service Greeter {
    Response SayHello(
        1:required User.User user
    )

    Response GetUser(
        1:required i32 uid
    )
}