## Day6

### Package

定义 --> `package` 关键字 报名通常和目录名一致

+ 一个文件夹就是一个包
  + 存放`*.go`文件

导入 --> import

* 起别名

  ~~~go
  import xxx "github.com/xxx/xxx"
  ~~~

* 匿名导入

  ~~~go
  import _ "不使用  但是需要执行func init(){}方法"
  //包导入会自动执行
  //一般用于做一些初始化的操作
  //一个包只能有一个init() 方法
  //外部应用包 套娃 先执行里面的
  ~~~

修饰符 `标识符`

+ 可见性

  ~~~go
  //使用首字母大小写来区别
  var Axxx int //其他包可见
  var axxx int //其他包不可见
  ~~~

   

### Interface

接口

+ 接口是一种类型`抽象的类型`
  + 需要实现方法的清单
+ 定义

~~~go
type mover interface{
  func run() //只写方法签名 参数返回值等 方法的样子
}
~~~

+ 接口变量

  + 实现了接口就可以当做 `当前接口`的变量

+ 接口的实现

  + 实现接口的所有方法 `自动向上转型`

+ 空接口

  + 所有的`struct`都实现了空接口
  + 任何类型都可以是空接口

+ 类型断言

  + 把一个interface{}转换成原本的类型

  ~~~go
  type aA interface{
    run()
  }
  
  type a struct{
    name string `json:name`
  }
  
  func (aaa *a) run(){
    fmt.Println("run")
  }
  
  var aA interface{} = new(a)
  
  var nn a = aA.(a)
  
  ~~~

  

  

### FileOption



