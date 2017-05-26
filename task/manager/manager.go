package manager

import (
	"mesos-framework-sdk/include/mesos_v1"
)

// Consts for mesos states.
const (
	RUNNING          = mesos_v1.TaskState_TASK_RUNNING
	KILLED           = mesos_v1.TaskState_TASK_KILLED
	LOST             = mesos_v1.TaskState_TASK_LOST
	GONE             = mesos_v1.TaskState_TASK_GONE
	STAGING          = mesos_v1.TaskState_TASK_STAGING
	STARTING         = mesos_v1.TaskState_TASK_STARTING // Default executor never sends this, it sends RUNNING directly.
	UNKNOWN          = mesos_v1.TaskState_TASK_UNKNOWN
	UNREACHABLE      = mesos_v1.TaskState_TASK_UNREACHABLE
	FINISHED         = mesos_v1.TaskState_TASK_FINISHED
	DROPPED          = mesos_v1.TaskState_TASK_DROPPED
	FAILED           = mesos_v1.TaskState_TASK_FAILED
	ERROR            = mesos_v1.TaskState_TASK_ERROR
	GONE_BY_OPERATOR = mesos_v1.TaskState_TASK_GONE_BY_OPERATOR
	KILLING          = mesos_v1.TaskState_TASK_KILLING
)

// Task manager holds information about tasks coming into the framework from the API
// It can set the state of a task.  How the implementation holds/handles those tasks
// is up to the end user.
type TaskManager interface {
	Add(*mesos_v1.TaskInfo) error
	Delete(*mesos_v1.TaskInfo) error
	Get(*string) (*mesos_v1.TaskInfo, error)
	GetById(id *mesos_v1.TaskID) (*mesos_v1.TaskInfo, error)
	HasTask(*mesos_v1.TaskInfo) bool
	Set(mesos_v1.TaskState, *mesos_v1.TaskInfo) error
	State(*string) (*mesos_v1.TaskState, error)
	AllByState(state mesos_v1.TaskState) ([]*mesos_v1.TaskInfo, error)
	TotalTasks() int
	All() ([]Task, error)
}

// Used to hold information about task states in the task manager.
// Task and its fields should be public so that we can encode/decode this.
type Task struct {
	Info  *mesos_v1.TaskInfo
	State mesos_v1.TaskState
}
