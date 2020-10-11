package main

import (
	"fmt"
	"strconv"
)
import s "strings"

/*
https://github.com/lifei6671/interview-go/blob/master/question/q006.md
问题描述

有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，
最开始，机器人位于 0 0，方向为正Y，可认为机器人在一个标准的直角坐标系中。 可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。
问最后机器人的坐标是多少？

*/

const (
	LEFT    string = "("
	RIGHT   string = ")"
	TOLEFT  int    = 0
	TOUP    int    = 1
	TORIGHT int    = 2
	TODOWN  int    = 3
)

type Stack struct {
	list []*Loc
}

func InitStack() *Stack {
	var list []*Loc
	list = []*Loc{}
	return &Stack{list: list}

}

func (s *Stack) Pop() (*Loc, error) {
	var (
		item *Loc
		err  error
	)
	if len(s.list) == 0 {
		return nil, err
	}
	item = s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return item, nil
}

func (s *Stack) Append(item *Loc) (left *Loc, right *Loc) {
	// 一个带有监控功能的添加操作
	s.list = append(s.list, item)
	if s.Size() > 1 {
		return s.matchEvent()
	}
	return nil, nil
}

func (s *Stack) matchEvent() (*Loc, *Loc) {
	size := s.Size() - 1
	prev := size - 1
	lastItem := s.list[size]
	prevItem := s.list[prev]
	if s.canMatch(prevItem, lastItem) {
		s.Pop()
		s.Pop()
		return prevItem, lastItem
	}
	return nil, nil
}

func (s *Stack) canMatch(left *Loc, right *Loc) bool {
	return left.direction == LEFT && right.direction == RIGHT
}

func (s *Stack) Size() int {
	return len(s.list)
}

func (s *Stack) isEmpty() bool {
	return s.Size() == 0
}

type Loc struct {
	signLoc   int
	repeatNum int
	numLoc    int
	direction string
}

type Parser struct {
	bytes  []byte
	curPtr int
	stack  *Stack
	endLoc int
}

func InitParser(str string) *Parser {
	bytes := []byte(str)
	size := len(bytes) - 1
	stack := &Stack{list: []*Loc{}}
	return &Parser{
		bytes:  bytes,
		curPtr: 0,
		stack:  stack,
		endLoc: size,
	}
}

func (p *Parser) reLocPtr(left *Loc, right *Loc) {
	var (
		command    string
		newBytes   []byte
		leftBytes  []byte
		rightBytes []byte
	)
	commandBytes := p.readCommands(left.signLoc+1, right.signLoc)
	repeatNum := p.getNum(left.numLoc)
	command = s.Repeat(string(commandBytes), repeatNum)
	newBytes = []byte(command)
	leftBytes, rightBytes = p.bytes[:left.numLoc], p.bytes[right.signLoc+1:]
	p.bytes = p.combine(leftBytes, newBytes, rightBytes)
	p.curPtr = len(leftBytes) + len(newBytes)
	p.endLoc = len(p.bytes) - 1
}

func (p *Parser) combine(left []byte, middle []byte, right []byte) []byte {
	var res []byte

	res = make([]byte, len(left))
	copy(res, left)
	res = append(res, middle...)

	res = append(res, right...)
	return res
}

func (p *Parser) getNum(loc int) int {
	res, _ := strconv.Atoi(string(p.bytes[loc]))
	return res

}

func (p *Parser) readCommands(start int, end int) []byte {
	return p.bytes[start:end]
}

func (p *Parser) isLeft() bool {
	return string(p.bytes[p.curPtr]) == LEFT
}

func (p *Parser) isRight() bool {
	return string(p.bytes[p.curPtr]) == RIGHT
}

func (p *Parser) current() string {
	return string(p.bytes[p.curPtr])
}

func (p *Parser) Parse() (string, error) {
	var (
		loc   *Loc
		left  *Loc
		right *Loc
		err   error
	)

	for p.curPtr <= p.endLoc {
		if p.isLeft() {
			loc = &Loc{
				p.curPtr,
				int(p.bytes[p.curPtr-1]),
				p.curPtr - 1,
				LEFT,
			}
			p.stack.Append(loc)
		} else if p.isRight() {
			loc = &Loc{
				p.curPtr,
				int(p.bytes[p.curPtr-1]),
				p.curPtr - 1,
				RIGHT,
			}
			left, right = p.stack.Append(loc)
			if left != nil && right != nil {
				p.reLocPtr(left, right)
			}
			continue //因为reLocPtr之后curPtr已经移动到未访问过的下一个char上了，不需要再在下面p.curPtr++了
		}

		p.curPtr++
	}
	//fmt.Printf("%c", p.bytes)
	fmt.Println("解析后command:", string(p.bytes))
	fmt.Println("指令数量：", len(p.bytes))
	//fmt.Println(p.stack.Size())
	if p.stack.Size() > 0 {
		return "", err
	}
	return string(p.bytes), nil
}

type RobotLoc []int

func (rl RobotLoc) getX() int {
	return rl[0]
}

func (rl RobotLoc) getY() int {
	return rl[1]
}

type Robot struct {
	direction int
	xLoc      int
	yLoc      int
	command   string
}

func InitRobot(direction int, startLoc []int, command string) (robot *Robot, err error) {
	if len(startLoc) != 2 {
		return nil, err
	}
	robot = &Robot{
		direction: direction,
		xLoc:      startLoc[0],
		yLoc:      startLoc[1],
		command:   command,
	}
	return
}

func (r *Robot) Move() {
	for _, command := range r.command {
		if string(command) == "R" {
			r.turnRight()
		} else if string(command) == "L" {
			r.turnLeft()
		} else if string(command) == "F" {
			r.moveForward()
		} else if string(command) == "B" {
			r.moveBack()
		}
	}
}

func (r *Robot) moveForward() {
	if r.direction == TOUP {
		r.yLoc++

	} else if r.direction == TODOWN {
		r.yLoc--
	} else if r.direction == TOLEFT {
		r.xLoc--
	} else if r.direction == TORIGHT {
		r.xLoc++
	}
}

func (r *Robot) moveBack() {
	if r.direction == TOUP {
		r.yLoc--

	} else if r.direction == TODOWN {
		r.yLoc++
	} else if r.direction == TOLEFT {
		r.xLoc++
	} else if r.direction == TORIGHT {
		r.xLoc--
	}
}

func (r *Robot) turnLeft() {
	if r.direction == TOUP {
		r.direction = TOLEFT

	} else if r.direction == TODOWN {
		r.direction = TORIGHT
	} else if r.direction == TOLEFT {
		r.direction = TODOWN
	} else if r.direction == TORIGHT {
		r.direction = TOUP
	}
}
func (r *Robot) turnRight() {
	if r.direction == TOUP {
		r.direction = TORIGHT

	} else if r.direction == TODOWN {
		r.direction = TOLEFT
	} else if r.direction == TOLEFT {
		r.direction = TOUP
	} else if r.direction == TORIGHT {
		r.direction = TODOWN
	}
}

func main() {
	var (
		parser        *Parser
		command       string
		parsedCommand string
		robot         *Robot
		err           error
		startLoc      []int
		direction     int
	)

	command = "FL2(F3(L2(FFFFFFFFFFFFFFFFF)RRF2(FLR))RB)"
	//command = "FL2(FF)R9(B)"
	//command = "F3(R2(FFFBBB)RR)RBB"
	//command = "FFFFFFFFFFFFF"
	fmt.Println("原始的command:", command)
	parser = InitParser(command) // TODO 问题一：数字不支持个位以上，比如 12(B)222(F);问题二：没有对输入command的正确性判断，括号可能会有错误
	if parsedCommand, err = parser.Parse(); err != nil {
		goto EXIT
	}

	direction = TORIGHT
	startLoc = []int{0, 0}
	if robot, err = InitRobot(direction, []int{0, 0}, parsedCommand); err != nil {
		goto EXIT
	}
	robot.Move()
	fmt.Println("机器人初始位置(x,y)：", startLoc)
	fmt.Println("机器人初始面向(0=向左，1=向下，2=向右，3=向下)：", direction)
	fmt.Println("机器人移动后位置：", []int{robot.xLoc, robot.yLoc})
	return
EXIT:
	fmt.Println("err:", err)
}
