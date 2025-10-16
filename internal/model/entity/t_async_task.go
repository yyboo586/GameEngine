package entity

type AsyncTask struct {
	ID            int64  `orm:"id"`
	CustomID      string `orm:"custom_id"`
	TaskType      int    `orm:"task_type"`
	Status        int    `orm:"status"`
	RetryCount    int    `orm:"retry_count"`
	Content       string `orm:"content"`
	Version       int    `orm:"version"`
	NextRetryTime int64  `orm:"next_retry_time"`
	CreateTime    int64  `orm:"create_time"`
	UpdateTime    int64  `orm:"update_time"`
}
