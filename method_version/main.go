package main

import (
	"fmt"
	"os"
)

// Person 公众
type Person struct {
	id   uint64
	name string
}

// Student 学生
type Student struct {
	class string
	Person
}

// Admin 管理员
type Admin struct {
	password string
	Person
}

func (a *Admin) newAdmin(id uint64, name string, password string) *Admin {
	return &Admin{
		password: password,
		Person:   Person{id: id, name: name},
	}
}
func (a *Admin) newStudent(id uint64, name string, class string) *Student {
	return &Student{
		Person: Person{id: id, name: name},
		class:  class,
	}
}

var (
	// AllStudent X
	AllStudent map[uint64]*Student
	// AllAdmin X
	AllAdmin map[uint64]*Admin
	// GID X
	GID uint64
)

func (a *Admin) addStudent() {
	var name string
	var class string
	var id uint64
	GID++
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生班级:")
	fmt.Scanln(&class)
	newStu := a.newStudent(id, name, class)
	AllStudent[GID] = newStu
	fmt.Print("添加成功!,继续请回车[enter]...")
	fmt.Scanln()
}
func (a *Admin) delStudent() {
	var id uint64
	fmt.Print("请输入学生id:")
	fmt.Scanln(&id)
	if _, ok := AllStudent[id]; ok {
		delete(AllStudent, id)
		fmt.Print("删除成功")
	} else {
		fmt.Printf("没有该学生【%d】", id)
	}
	fmt.Print(",继续请回车[enter]...")
	fmt.Scanln()
}
func (a *Admin) updateStudent() {
	var id uint64
	fmt.Print("请输入需要修改的学生id:")
	fmt.Scanln(&id)
	if oldData, ok := AllStudent[id]; ok { //allStudent的值是一个地址,因为声明的是否是map[uint64]*student
		var uid uint64
		var uname string
		var uclass string
		fmt.Println("将要修改的学生信息:", *oldData)
		fmt.Print("请输入学生学号:")
		fmt.Scanln(&uid)
		fmt.Print("请输入学生姓名:")
		fmt.Scanln(&uname)
		fmt.Print("请输入学生班级:")
		fmt.Scanln(&uclass)
		oldData.Person.id = uid
		oldData.Person.name = uname
		oldData.class = uclass
		AllStudent[id] = oldData
		fmt.Println(oldData)
		fmt.Println(uid)
		fmt.Println(AllStudent[id])
		fmt.Print("修改成功")
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
	if stu, ok := AllStudent[id]; ok {
		fmt.Printf("%v", *stu)
	} else {
		fmt.Print("没有该学生")
	}
	fmt.Print(",继续请回车[enter]...")
	fmt.Scanln()
}

func queryAllStudent() {
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	fmt.Println("\tid\t学号\t姓名\t\t班级")
	for k, v := range AllStudent {
		fmt.Printf("\t%d\t%d\t%v\t\t%s\n", k, v.Person.id, v.name, v.class)
	}
	fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
}

//PublicMenu X
func PublicMenu() {
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

// AdminMenu X
func AdminMenu(a *Admin) {
	for {
		fmt.Printf("=========================欢迎回来,【%#v】管理员=========================\n", a.Person.name)
		fmt.Println("\t\t0.注销登录\t1.增加学生\t2.删除学生")
		fmt.Println("\t\t3.修改学生\t4.查询学生\t5.查询所有")
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
			a.updateStudent()
		case 4:
			queryStudent()
		case 5:
			queryAllStudent()
		default:
			fmt.Println("输入错误！请输入屏幕的数字")
		}
	}
}

func main() {
	AllAdmin = make(map[uint64]*Admin, 10)
	adminObj := Admin{Person: Person{id: 1, name: "admin"}, password: "123456"}
	AllAdmin[1] = &adminObj
	AllStudent = make(map[uint64]*Student, 50)
	GID = 0
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
			PublicMenu()
		case 2:
			fmt.Print("请输入您的账户ID:")
			fmt.Scanln(&pID)
			//查询是否是管理员
			if value, ok := AllAdmin[pID]; ok {
				fmt.Printf("请输入密码: ")
				fmt.Scanln(&pPassword)
				if AllAdmin[pID].password == pPassword {
					AdminMenu(value) //默认成功，跳转管理员界面
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
