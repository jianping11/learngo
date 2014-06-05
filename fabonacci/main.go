package main

import (
	"fmt"
	"time"
)

func main() {
	// begin := time.Now()
	// for i := 0; i < 45; i++ {
	// 	begin := time.Now()
	// 	fmt.Print("坐标", i, " :", Fabonacci(i), " ")
	// 	end := time.Now()
	// 	durication := end.Sub(begin).Seconds()
	// 	fmt.Printf("耗时：%f 秒 \n", durication)
	// }
	// end := time.Now()
	// durication := end.Sub(begin).Seconds()
	// fmt.Printf("总耗时：%f 秒 \n", durication)

	begin := time.Now()
	for i := 0; i < 1000; i++ {
		// begin := time.Now()
		fmt.Println("坐标", i, " :", fib(i), " ")
		// end := time.Now()
		// durication := end.Sub(begin).Seconds()
		// fmt.Printf("耗时：%f 秒 \n", durication)
	}
	// fmt.Println(fib(45))
	// fmt.Print(Fabonacci(5))
	// fmt.Println(fib2(45))
	end := time.Now()
	// count = 0
	durication := end.Sub(begin).Seconds()
	fmt.Printf("总耗时：%f 秒 \n", durication)

	begin = time.Now()
	for i := 0; i < 1000; i++ {
		// begin := time.Now()
		fmt.Println("坐标", i, " :", fib2(i), " ")
		// end := time.Now()
		// durication := end.Sub(begin).Seconds()
		// fmt.Printf("耗时：%f 秒 \n", durication)
	}
	// fmt.Println(fib(45))
	// fmt.Print(Fabonacci(5))
	// fmt.Println(fib2(45))
	end = time.Now()
	// count = 0
	durication = end.Sub(begin).Seconds()
	fmt.Printf("总耗时：%f 秒 \n", durication)

}

//*****************************************************************************
// 使用递归来计算序列 简洁
func Fabonacci(n int) int {
	if n < 2 {
		return n
	}

	return Fabonacci(n-1) + Fabonacci(n-2)
}

//*****************************************************************************
// 使用递推来计算序列 性能好
func fib2(n int) int {
	// count++
	if n < 2 {
		return n
	}

	n1 := 1
	n2 := 1
	sn := 0
	for i := 0; i < n-2; i++ {
		sn = n1 + n2
		n1 = n2
		n2 = sn
	}
	return sn
}

//*****************************************************************************
// 使用缓存来计算，如果日常计算量很大哦，应该设置一个公共缓存，保存结果，综合性能更好
type memoizeFunction func(int, ...int) interface{}

var fib memoizeFunction

func init() {
	//fib的核心还是使用递归，不过在init中，将递归的算法放入到Memoize中了
	fib = Memoize(func(x int, xs ...int) interface{} {
		if x < 2 {
			//计算中
			fmt.Println("计算中")
			return x
		}
		return fib(x-1).(int) + fib(x-2).(int)
	})
}

func Memoize(function memoizeFunction) memoizeFunction {
	//创建一个缓存结果的映射cache，映射的键，映射的值是
	cache := make(map[string]interface{})
	//返回memoizeFunction 签名函数
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}
		// fmt.Println("key:", key)
		//查找映射值
		if value, found := cache[key]; found {
			//找到值
			return value
		}
		//映射值不存在就调用function（fib传进来的递归算法）计算值
		value := function(x, xs...)
		//将计算的值放入到里面
		cache[key] = value
		return value
	}
}
