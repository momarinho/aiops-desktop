package processes

type ProcessInfo struct {
	PID         int     `json:"pid"`
	Name        string  `json:"name"`
	User        string  `json:"user"`
	CPUPercent  float64 `json:"cpu_percent"`
	MemoryMB    float64 `json:"memory_mb"`
	CreateTime  string  `json:"create_time"`
	Command     string  `json:"command"`
	IsCritical  bool    `json:"is_critical"`
	Status      string  `json:"status"`
}

// Critical system processes that should never be killed
var CriticalProcesses = map[int]string{
	1:  "init",
	2:  "kthreadd",
	3:  "rcu_gp",
	4:  "rcu_par_gp",
	6:  "kworker/0:0H-kblockd",
	7:  "mm_percpu_wq",
	8:  "ksoftirqd/0",
	9:  "rcu_sched",
	10: "migration/0",
	11: "rcu_preempt",
	12: "rcu_sched_kthread",
	13: "rcu_exp_par_gp",
	14: "rcu_exp_kthread",
}

func IsCriticalPID(pid int) bool {
	_, exists := CriticalProcesses[pid]
	return exists
}

func GetProcessStatus(pid int) string {
	if IsCriticalPID(pid) {
		return "critical"
	}
	if pid < 100 {
		return "system"
	}
	return "user"
}
