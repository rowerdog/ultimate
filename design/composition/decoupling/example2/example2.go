package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Data是我们pull和store的数据
type Data struct {
	Line string
}

type Puller interface {
	Pull(d *Data) error
}

type Storeer interface {
	Store(d *Data) error
}

type Xenia struct {
	Name string
}

// 从Xenia中pull数据
func (x *Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("error reading data from Xenia")

	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

type S3 struct {
	Name string
}

func (s *S3) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// 组成一个知道如何提取和存储的系统
type System struct {
	Xenia
	S3
}

// 转换成多态函数
func Pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Store(s Storeer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(s *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := Pull(&s.Xenia, data)
		if i > 0 {
			if _, err := Store(&s.S3, data); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
}
func main() {
	s := System{
		Xenia{
			"tesla",
		},
		S3{
			"AWS",
		},
	}
	if err := Copy(&s, 3); err != nil {
		fmt.Println(err)
	}
}
