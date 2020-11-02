package engine

import (
	"fmt"
	"spider/fetcher"
)

/**
 * 启动器
 */
type Engine struct {
	Scheduler ScheduleInterface
}

/**
 * 接口
 */
type ScheduleInterface interface {
	WorkerReady(chan Request)
	WorkerChan() chan Request
	Submit(Request)
	Run()
}

/**
 * 总控
 */
func (e *Engine) Run(seeds ...Request) {
	// 开worker和request
	e.Scheduler.Run()
	in := e.Scheduler.WorkerChan()
	out := make(chan ParserResult)
	// 所有worker共用这些channel, 数量可自定义
	for i := 0; i < 10; i++ {
		e.createWorker(in, out)
	}
	// 将初始化的所有请求传给request队列
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 遍历所有请求，根据对应的解析器进行解析
	emptyDetail := Details{}
	num := 0
	for {
		// 拿第一个元素并将其从切片中删除
		parserResult := <- out
		//parserResult := worker(r)
		for _, item := range parserResult.Items {
			if item == emptyDetail { // 没内容不显示
				continue
			}
			fmt.Printf("Parser Item #%d: %+v\n", num, item)
			num++
		}

		// 把获取到的url加到requests中继续走流程
		for _, request := range parserResult.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

/**
 * 并发调worker
 */
func (e *Engine) createWorker(in chan Request, out chan ParserResult) {
	// 这里做两件事, 这两件事不是顺序的, 必要被代码顺序锁误导
	go func() {
		for {
			// 第一件事: 把in传给workerChan加入到worker的消息队列中
			e.Scheduler.WorkerReady(in)
			// 第二件事: 从worker中提取调度器中成功拿到的request并交给worker处理
			request := <- in
			parserResult := worker(request)
			out <- parserResult
		}
	}()
}

/**
 * worker: 请求url获取页面和解析页面
 */
func worker(r Request) ParserResult {
	// 获取页面
	content, err := fetcher.Fetcher(r.Url)
	if err != nil {
		fmt.Printf("Fetching error: %v\n", err)
	}

	// 根据对应的解析器解析页面内容
	parserResult := r.ParserFunc(content)
	if err != nil {
		fmt.Printf("Parse Error: %v\n", err)
	}
	return parserResult
}
