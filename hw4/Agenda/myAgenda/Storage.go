package myAgenda

import (
	"container/list"
	"os"
	"fmt"
	"bufio"
	"encoding/json"
)

//Storage存储用户与会议数据
type Storage struct {
	m_userList    list.List
	m_meetingList list.List
	m_curUserList list.List
}

//从文件中读取数据
func (s *Storage) ReadFormFile() bool {
	userFin, uErr := os.OpenFile("user.json", os.O_CREATE, 0666)
	if uErr != nil {
		fmt.Fprintf(os.Stderr, "could not open user file\n")
		return false
	}
	defer userFin.Close()
	userReader := bufio.NewReader(userFin)

	for {
		line, crc := userReader.ReadString('\n')
		if crc != nil && len(line) == 0 {
			break
		}
		var user User
		var jsonData []byte
		fmt.Sscanf(line, "%s\n", &jsonData)
		json.Unmarshal(jsonData, &user)
		s.m_userList.PushBack(user)
	}

	meetingFin, mErr := os.OpenFile("meeting.json", os.O_CREATE, 0666)
	if mErr != nil {
		fmt.Fprintf(os.Stderr, "could not open meeting file\n")
		return false
	}
	defer meetingFin.Close()
	meetingReader := bufio.NewReader(meetingFin)

	for {
		line, crc := meetingReader.ReadString('\n')
		if crc != nil && len(line) == 0 {
			break
		}
		var meeting Meeting
		var jsonData []byte
		fmt.Sscanf(line, "%s\n", &jsonData)
		json.Unmarshal(jsonData, &meeting)
		s.m_meetingList.PushBack(meeting)
	}

	curFin, cErr := os.OpenFile("curUser.txt", os.O_CREATE, 0666)
	if cErr != nil {
		fmt.Fprintf(os.Stderr, "could not open curUser file\n")
		return false
	}
	defer curFin.Close()
	curReader := bufio.NewReader(curFin)

	for {
		line, crc := curReader.ReadString('\n')
		if crc != nil && len(line) == 0 {
			break
		}
		var username string
		fmt.Sscanf(line, "%s\n", &username)
		s.m_curUserList.PushBack(username)
	}
	return true
}

//将数据写入文件
func (s *Storage) WriteToFile() bool {
	userFout, uErr := os.Create("user.json")
	if uErr != nil {
		fmt.Fprintf(os.Stderr, "could not open user file\n")
		return false
	}
	defer userFout.Close()

	for {
		if s.m_userList.Len() == 0 {
			break
		}
		element := s.m_userList.Front()
		user := element.Value.(User)
		jsonData, jErr := json.Marshal(user)
		if jErr != nil {
			panic(jErr)
		}
		fmt.Fprintf(userFout, "%s\n", jsonData)
		s.m_userList.Remove(element)
	}

	meetingFout, mErr := os.Create("meeting.json")
	if mErr != nil {
		fmt.Fprintf(os.Stderr, "could not open meeting file\n")
		return false
	}
	defer meetingFout.Close()

	for {
		if s.m_meetingList.Len() == 0 {
			break
		}
		element := s.m_meetingList.Front()
		meeting := element.Value.(Meeting)
		jsonData, jErr := json.Marshal(meeting)
		if jErr != nil {
			panic(jErr)
		}
		fmt.Fprintf(meetingFout, "%s\n", jsonData)
		s.m_meetingList.Remove(element)
	}

	curFout, cErr := os.Create("curUser.txt")
	if cErr != nil {
		fmt.Fprintf(os.Stderr, "could not open curUser file\n")
		return false
	}
	defer curFout.Close()

	for {
		if s.m_curUserList.Len() == 0 {
			break
		}
		element := s.m_curUserList.Front()
		user := element.Value.(string)
		fmt.Fprintf(curFout, "%s\n", user)
		s.m_curUserList.Remove(element)
	}
	return true
}

//创建用户
func (s *Storage) CreateUser(u User) {
	s.m_userList.PushBack(u)
}

//创建登录用户
func (s *Storage) CreateCurUser(username string) {
	s.m_curUserList.PushBack(username)
}

//删除用户
func (s *Storage) DeleteUser(u User) {
	for i := s.m_userList.Front(); i != nil; i = i.Next() {
		user := i.Value.(User)
		if user.M_name == u.M_name {
			s.m_userList.Remove(i)
			break
		}
	}
}

//删除登录用户
func (s *Storage) DeleteCurUser(username string) {
	for i := s.m_curUserList.Front(); i != nil; i = i.Next() {
		user := i.Value.(string)
		if username == user {
			s.m_curUserList.Remove(i)
			break
		}
	}
}

//查询用户
func (s *Storage) QueryUser(filter func(User) bool) list.List {
	var res list.List
	for i := s.m_userList.Front(); i != nil; i = i.Next() {
		user := i.Value.(User)
		if filter(user) {
			res.PushBack(user)
		}
	}
	return res
}

//查询登录用户
func (s *Storage) QueryCurUser(filter func(string) bool) list.List {
	var res list.List
	for i := s.m_curUserList.Front(); i != nil; i = i.Next() {
		user := i.Value.(string)
		if filter(user) {
			res.PushBack(i.Value.(string))
		}
	}
	return res
}

//创建会议
func (s *Storage) CreateMeeting(m Meeting) {
	s.m_meetingList.PushBack(m)
}

//查询会议
func (s *Storage) QueryMeeting(filter func(Meeting) bool) list.List {
	var res list.List
	for i := s.m_meetingList.Front(); i != nil; i = i.Next() {
		meeting := i.Value.(Meeting)
		if filter(meeting) {
			res.PushBack(meeting)
		}
	}
	return res
}

//更新会议
func (s *Storage) UpdateMeeting(filter func(Meeting) bool, switcher func(*Meeting)) int {
	count := 0
	for i := s.m_meetingList.Front(); i != nil; i = i.Next() {
		meeting := i.Value.(Meeting)
		if filter(meeting) {
			switcher(&meeting)
			i.Value = meeting
			count++
		}
	}
	return count
}

//删除会议
func (s *Storage) DeleteMeeting(filter func(Meeting) bool) int {
	count := 0
	for i := s.m_meetingList.Front(); i != nil; {
		meeting := i.Value.(Meeting)
		if filter(meeting) {
			j := i
			i = i.Next()
			s.m_meetingList.Remove(j)
			count++
		} else {
			i = i.Next()
		}
	}
	return count
}
