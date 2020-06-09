package main

import (
	"fmt"
	"os"
)

type person struct {
	id   uint64
	name string
}
type student struct {
	class string
	person
}
type admin struct {
	password string
	person
}

func (a *admin) newAdmin(id uint64, name string, password string) *admin {
	return &admin{
		password: password,
		person:   person{id: id, name: name},
	}
}
func (a *admin) newStudent(id uint64, name string, class string) *student {
	return &student{
		person: person{id: id, name: name},
		class:  class,
	}
}
func (a *admin) addStudent() {
	var name string
	var class string
	var id uint64
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	fmt.Print("请输入学生ID:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生班级:")
	fmt.Scanln(&class)
	newStu := a.newStudent(id, name, class)
	allStudent[id] = newStu
	fmt.Print("添加成功!,继续请回车[enter]...")
	fmt.Scanln()
}
func (a *admin) delStudent() {
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
	fmt.Println("\t学号\t姓名\t班级")
	for k, v := range allStudent {
		fmt.Printf("\t%d\t%v\t%s\n", k, v.name, v.class)
	}
	fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
}

var (
	allStudent map[uint64]*student
	allAdmin   map[uint64]*admin
)

func publicMenu() {
	for {
		fmt.Println("=========================欢迎回来【公共版】=========================")
		fmt.Println("\t0.返回首页\t1.查询学生\t2.查询所有")
		fmt.Printf("请输入你的选择:")
		var choice int8
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			return
		case 1:
			queryStudent()
		case 2:
			queryAllStudent()
		default:
			fmt.Println("输入错误！请输入屏幕的数字")
		}
	}
}
func adminMenu(a *admin) {
	for {
		fmt.Printf("=========================欢迎回来,【%#v】管理员=========================\n", a.person.name)
		fmt.Println("\t0.注销登录\t1.增加学生\t2.删除学生\t3.查询学生\t4.查询所有")
		fmt.Printf("请输入你的选择:")
		var choice int8
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			return
		case 1:
			a.addStudent()
		case 2:
			a.delStudent()
		case 3:
			queryStudent()
		case 4:
			queryAllStudent()
		default:
			fmt.Println("输入错误！请输入屏幕的数字")
		}
	}

}
func main() {
	allAdmin = make(map[uint64]*admin, 10)
	adminObj := admin{person: person{id: 1, name: "admin"}, password: "123456"}
	allAdmin[1] = &adminObj
	allStudent = make(map[uint64]*student, 50)

	var pID uint64
	var pChoice uint8
	var pPassword string
	for {
		fmt.Println("=========================欢迎来到KNHO学生管理系统=========================")
		fmt.Println("\t0.退出系统\t1.公共版面\t2.管理版面")
		fmt.Print("请输入您的选择:")
		fmt.Scanln(&pChoice)
		switch pChoice {
		case 0:
			os.Exit(1)
		case 1:
			publicMenu()
		case 2:
			fmt.Print("请输入您的账户ID:")
			fmt.Scanln(&pID)
			//查询是否是管理员
			if value, ok := allAdmin[pID]; ok {
				fmt.Printf("请输入密码: ")
				fmt.Scanln(&pPassword)
				if allAdmin[pID].password == pPassword {
					adminMenu(value) //默认成功，跳转管理员界面
				} else {
					fmt.Print("[logging failed] 密码错误,请核实您账户的密码! 继续请回车[enter]...")
					fmt.Scanln()
				}
			} else {
				fmt.Print("[logging failed] 账户不存在,请核实您的账户! 继续请回车[enter]...")
				fmt.Scanln()
			}
		default:
			fmt.Println("输入错误！请输入屏幕的数字")
		}

	}
}
