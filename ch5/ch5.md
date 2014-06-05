## 5.5 defer 、panic 和 recover
### defer
延迟函数，在外围函数或者方法返回前但是其返回值（如果有）计算之后执行。多个defer LIFO（后进先出）
，通常用来保证方法执行完之后释放资源，如关闭文件，通道，数据库连接等。

### panic 和 recover
panic和recover涉及到go语言的异常处理等。通常错误使用返回值error返回，抛给外围函数进行处理，而panic
在发生“不可能发生”的情况下会被调用。  
早期开发时可以使用panic中断程序执行，强制发生错误,方便及早发现错误，及时处理。后期部署的时候，应该避免
程序中断。  

我们可以结合recover()函数来一同使用panic，panic包含一个错误，通过recover来捕获错误。  
例如：eg：一个简单的hello函数

    package main

    import (
      "fmt"
    )

    func main() {
      hello()
    }

    func hello() {
      //通过延迟函数调用recover来捕获panic的异常
      defer func() {
        if err := recover(); err != nil {
          fmt.Println(err)
        }
      }()
      fmt.Println("Hello world, this func while panic")
      panic("test panic") //强制panic，并添加错误信息
    }

使用这种方式的时候，应该将error记录到日志中，不隐藏该问题。
程序中的对于panic 和 error的处理方式，可以参考regexp的Compile 和 MustCompile

让一个web服务器在遇到异常时仍能健壮的运行。应该为handle添加一个recover函数捕获异常。
例如statistics例子中，为homePage添加一个recover来处理

    func homePage(writer http.ResponseWriter, request *http.Request) {
      defer func() { //每一个handle 页面都需要一个defer recover
        if x := recover(); x != nil {
          log.Printf("[%v] caught panic: %v",request.RemoteAddr, x)
        }
      }()

      ....
    }

多个页面时，使用包装函数来处理。创建一个logPanics来包装homePage handler，

    http.HandleFunc("/",logPanics(homePage))
    ---------------------------------------
    //logPanics,给handler 添加recover ，返回一个包装的匿名handler
    func logPanics(function func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
      return func(writer http.ResponseWriter, request *http.Request) {
        defer func() {
          if x := recover(); x != nil {
            log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
          }
        }()
        function(writer, request)
      }
    }

---

---

## 5.6 自定义函数

小要点：对于一般有返回值的if  else 和 switch 语句，可以将return语句之后，不添加相应的
else语句和default分支。

### 函数参数

可变类型参数，例如

    func Sum(first int, rest ...int)

rest参数就是一个可变长度的int参数，如果我们有一个int的数组切片，只需要在slice后面
放一个省略号就行了。nums[..]...

---
可选参数的函数，(参数实力：image.jpeg.Encode())


go语言中没有想javascript支持可选参数，
假设我们有一个函数用来处理一些自定义的数据，默认就是简单的处理所有的数据，但有些时候我们希望
可以制定处理第一个或者最后一个项，还有是否记录函数的行为，或者对于非法项作处理等，
一个办法就是创建一个方法签名：

    func ProcessItems(items Items, first, last int, audit bool, errorHandler func(item Item))

在这个设计里，如果last为0的话意味着从first开始处理所有项。errorHandler为nil，则表示不做非法项的处理
所以默认行为就写为

    ProcessItems(items,0,0,false,nil)

较为优雅的方式是定义一个结构体保存Options，保存上面的参数，初始化为零值，这样大部分调用都可以别简化为
ProcessItems(items, Options{}).需要额外处理时，就添加相应字段就行了。

    type Options struct {
      First int //要处理的第一项
      Last int  //要处理的最后一项（0意味从第一项开始处理所有的项)
      Audit bool // true为记录所有动作
      ErrorHandler func(item Item) //不为nil时，对每一项调用一次
    }

    ProcessItem(items, Options{}) //默认选项

    errorHandler := func(item Item) {log.Println("Invalid:", item)}
    ProcessItem(items, Options{Audit:true, ErrorHandler:errorHandler})

---
init() 函数和main() 函数

在go程序的执行顺序是
1. 从main包中开始，引入其他的包
2. 包引入完毕后，创建包的一些常量和变量
3. 调用包的init() 函数
4. 执行main包的main()函数，程序运行

> 一个包中应该创建一个init.go, 负责包的初始化，init函数调用等。

---

---

### 闭包、递归调用
> go支持闭包

> 可以创建一个变量，对匿名函数（闭包）进行引用，如：thisfunc := func(name stirng){ return name}

递归函数就是自己调用自己，在编写递归函数的时候，应该注意
