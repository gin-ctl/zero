package time

//gorm中重新格式化json时间数据格式返回给前端
import (
	"time"
)

import (
	"database/sql/driver"
	"fmt"
)

const (
	timezone   = "Asia/Shanghai"
	timeFormat = "2006-01-02 15:04:05"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t Time) Local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Time(t).In(loc)
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *Time) AddDate(years int, months int, days int) Time {
	return Time(time.Time(*t).AddDate(years, months, days))
}

func (t Time) DaysUntilExpiration() int {
	currentTime := time.Now()
	expiration := time.Time(t)

	// 计算过期日期与当前日期之间的差距
	days := int(expiration.Sub(currentTime).Hours() / 24)

	return days
}
