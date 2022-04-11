package game

type TaskInfo struct {
	TaskId int
	State  int
}

type ModUniqueTask struct {
	MyTaskInfo map[int]*TaskInfo
} //map 不加锁 线上一定会崩溃

func (self *ModUniqueTask) IsTaskFinish(taskId int) bool {
	if taskId == 10001 || taskId == 10002 {
		return true
	}
	task, ok := self.MyTaskInfo[taskId]
	if !ok {
		return false
	}
	return task.State == TASK_STATE_FINISH
}
