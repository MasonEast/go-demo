package main

/*
通道可被认为是Goroutine通信的管道

虽然Goroutine共享数据可以通过传统的同步机制锁来解决，但还是建议使用Channel通道来通信解决

通过通信来传递内存数据，使得内存数据在不同的goroutine中传递，而不是使用共享内存来通信

channel是引用类型，在作为参数传递时，传递的是内存地址

channel是同步的，意味着同一时间，只能有一条goroutine来操作

通道是goroutine之间的连接，所以通道的发送和接收必须处在不同的goroutine中

发送和接收默认是阻塞的

缓冲通道：带有一个缓冲区的通道，发送到一个缓冲通道只有在缓冲区满时才会阻塞，类似的，接收只有在缓冲区空时才阻塞
*/
func main() {
	channelFn()
	calcSquaresFn()
	rangeFn()
	bufferChFn()
	singleWayFn()
	timerFn()
	selectFn()
}