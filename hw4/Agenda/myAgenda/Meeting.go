package myAgenda

//Meeting包装会议数据
type Meeting struct {
	M_sponsor string
	M_participators []string
	M_startDate Date
	M_endDate Date
	M_title string
}

//删除一个参与者
func (m *Meeting)RemoveParticipator(t_participator string)  {
	for element := 0; element< len(m.M_participators);element++{
		if m.M_participators[element] == t_participator{
			m.M_participators = append(m.M_participators[0:element], m.M_participators[element+1:len(m.M_participators)]...)
			break
		}
	}
}

//添加一个参与者
func (m *Meeting)AddParticipator(t_participator string)  {
	m.M_participators = append(m.M_participators, t_participator)
}

//判断参与者是否在会议中
func (m Meeting)IsParticipator(t_participator string) bool {
	for element := 0; element< len(m.M_participators);element++{
		if m.M_participators[element] == t_participator{
			return true
		}
	}
	return false
}




