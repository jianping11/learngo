## 5.5 defer 、panic 和 recover
###defer
延迟函数，在外围函数或者方法返回前但是其返回值（如果有）计算之后执行。多个defer LIFO（后进先出）
，通常用来保证方法执行完之后释放资源，如关闭文件，通道，数据库连接等。

###panic 和 recover
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
