# RPC

## Golang的RPC包

### 内置RPC包介绍

服务端注册一个对象，注册的对象会作为一个服务被暴露给客户端。对象中可导出的方法可以被远程访问。  
对象的方法必须满足以下要求，否则将会被忽略:

1) 方法是可导出的
2) 方法有两个参数，都是可导出类型或内置类型
3) 方法的第二个参数是指针
4) 方法只有一个error接口类型的返回值

方法示例:

```go
func (t *T) MethodName(argType T1, replyType *T2) error
```

### 内置RPC包分析

1) 如何判断Golang的类型是否为内置类型?

```go
func isBuiltinType(t reflect.Type) bool {
return t.PkgPath() == ""
}
```
