package response

import "time"

type OperationRecord struct {
	Ip           string        `json:"ip" `
	Method       string        `json:"method"`
	Path         string        `json:"path"`
	Status       int           `json:"status" `
	Latency      time.Duration `json:"latency"`
	Agent        string        `json:"agent" `
	ErrorMessage string        `json:"error_message" `
	Body         string        `json:"body"`
	Resp         string        `json:"resp"`
	ManagerID    int64         `json:"manager_id"`
	RuleName     string        `json:"rule_name"`
	CreateTime   int64         `json:"create_time"`
	UpdateTime   int64         `json:"update_time"`
	DeleteTime   int64         `json:"delete_time"`
}

type OperationRecordList struct {
	Id           int64         `json:"id" gorm:"primaryKey;column:id;"`
	Ip           string        `json:"ip"`
	Method       string        `json:"method"`
	Path         string        `json:"path"`
	Status       int           `json:"status" `
	Latency      time.Duration `json:"latency"`
	Agent        string        `json:"agent" `
	ErrorMessage string        `json:"error_message" `
	Body         string        `json:"body"`
	Resp         string        `json:"resp"`
	ManagerID    int64         `json:"manager_id"`
	RuleName     string        `json:"rule_name"`
	CreateTime   string        `json:"create_time"`
	UpdateTime   string        `json:"update_time"`
	DeleteTime   int64         `json:"delete_time"`
	Username     string        `json:"username"`
}
