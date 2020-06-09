package main

import (
	"fmt"
	"os"
)

var allStudent map[uint64]*student

type student struct {
	id   uint64
	name string
}

func newStudent(id uint64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
func addStudent() {
	var name string
	var id uint64
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	fmt.Print("请输入学生ID:")
	fmt.Scanln(&id)
	newStu := newStudent(id, name)
	allStudent[id] = newStu
	fmt.Print("添加成功!,继续请回车[enter]...")
	fmt.Scanln()
}
func delStudent() {
	var id uint64
	fmt.Print("请输入学生id:")
	fmt.Scanln(&id)
	if _, ok := allStudent[id]; ok {
		delete(allStudent, id)
		fmt.Print("删除成功")
	} else {
		fmt.Printf("没有该学生【%d】", id)
	}
	fmt.Print(",继续请回车[enter]...")
	fmt.Scanln()

}
func queryStudent() {
	var id uint64
	fmt.Print("请输入学生id:")
	fmt.Scanln(&id)
	if stu, ok := allStudent[id]; ok {
		fmt.Printf("%v", *stu)
	} else {
		fmt.Print("没有该学生")
	}
	fmt.Print(",继续请回车[enter]...")
	fmt.Scanln()
}
func queryAllStudent() {
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	fmt.Println("学号\t姓名")
	for k, v := range allStudent {
		fmt.Printf("%d\t%v\n", k, v.name)
	}
	fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
}

func main() {
	allStudent = make(map[uint64]*student, 50)
	for {
		fmt.Println("==========================================================")
		fmt.Println("0.退出\t1.增加学生\t2.删除学生\t3.查询学生\t4.查询所有")
		fmt.Printf("请输入你的选择:")
		var choice int8
		fmt.Scanln(&choice)
		fmt.Printf("你的选择是:%d\n", choice)
		switch choice {
		case 0:
			os.Exit(1)
		case 1:
			addStudent()
		case 2:
			delStudent()
		case 3:
			queryStudent()
		case 4:
			queryAllStudent()
		default:
			fmt.Println("输入错误！请输入屏幕的数字")
		}
	}
}
