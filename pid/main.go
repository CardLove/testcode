package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {

	// 获取所有的pids
	pids, _ := process.Pids()
	fmt.Println(pids)

	// 返回pid是否存在
	isExist, _ := process.PidExists(4)
	fmt.Println(isExist)

	// 获取所有进程，返回进程的数组
	ps, _ := process.Processes()
	fmt.Println(ps)

	// 获取pid进程
	p, _ := process.NewProcess(87076)
	fmt.Println(p.Pid)

	// 返回进程的cpu使用率(统计一段时间内cpu的使用率)
	fmt.Println(p.Percent(1 * time.Second))
	// /proc/pid/stat字段说明 https://blog.csdn.net/zyboy2000/article/details/50456764
	// CPU使用率(计算公式 当前进程总CPU运行时间(utime+stime)/(当前时间-进程创建时间))
	fmt.Println(p.CPUPercent())
	// 返回进程状态
	fmt.Println(p.Status())
	// 返回进程连接
	fmt.Println(p.Connections())
	// 返回进程的timeStat信息
	fmt.Println(p.Times())
	// 进程名称
	fmt.Println(p.Name())
	// 进程创建时间
	fmt.Println(p.CreateTime())
	// 获取进程的子进程
	fmt.Println(p.Children())
	// 获取进程的命令行
	fmt.Println(p.Cmdline())

}
