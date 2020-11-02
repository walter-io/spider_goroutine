package scheduler

import "spider/engine"

/**
 * 并发调度器
 * 下面有requestChan和workerChan, 目的是控制将所有request交给worker管理
 * 这样的可以有效的控制并发出合理的worker的数量
 */
type Scheduler struct {
	requestChan chan engine.Request  // request channel
	workerChan  chan chan engine.Request  // request manager 即request管理器
}

/**
 * 定义队列
 */
func (s *Scheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {

		var requestQue []engine.Request
		var workerQue []chan engine.Request
		for {
			// 获取channel中request和worker
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQue) > 0 && len(workerQue) > 0 {
				activeRequest = requestQue[0]
				activeWorker = workerQue[0]
			}
			select {
			case r := <-s.requestChan:
				// 接收request并放到request队列
				requestQue = append(requestQue, r)
			case w := <-s.workerChan:
				// 接收worker并放到worker队列
				workerQue = append(workerQue, w)
			case activeWorker <- activeRequest:
				// 把request给worker, 并将其从队列中删除
				requestQue = requestQue[1:]
				workerQue = workerQue[1:]
			}
		}
	}()
}

// 实现ScheduleInterface接口: 添加worker
func (s *Scheduler) WorkerReady(worker chan engine.Request) {
	s.workerChan <- worker
}

// 实现ScheduleInterface接口: 创建channel
func (s *Scheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// 实现ScheduleInterface接口: 往channel添加request
func (s *Scheduler) Submit(request engine.Request) {
	s.requestChan <- request
}
