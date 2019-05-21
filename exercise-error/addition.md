#### 为什么fmt.Sprint(e)会使程序陷入死循环的
```
func (e ErrNegativeSqrt) Error() string {
return fmt.Sprintln("cannot Sqrt negative number", e)
}
```
或者
```
func (e ErrNegativeSqrt) Error() string {
return fmt.Sprint(e)
}
```
> 这么写会陷入死循环。
> 无限递归陷入的死循环
> 因为e变量是一个通过实现Error()的接口函数成为了error类型，那么在```fmt.Sprint(e)```时就会调用```e.Error()```来输出错误的字符串信息
于是函数相当于
```
func (e ErrNegativeSqrt) Error() string {
return fmt.Sprint(e.Error())
}
```
从而陷入了无限递归之中

