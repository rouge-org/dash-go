package dash

const (
	TimeNone int64 = -1
)

type TimedRecord struct {
	TTL           int64 `json:"ttl"`
	TimeCreation  int64 `json:"t_c"`
	TimeUpdate    int64 `json:"t_u"`
	TimeRead      int64 `json:"t_r"`
	TimeHeartbeat int64 `json:"t_h"`
}

func NewTimedRecord() (r *TimedRecord) {
	r = new(TimedRecord)

	r.TTL = TimeNone
	r.TimeCreation = NowMilli()
	r.TimeUpdate = TimeNone
	r.TimeRead = TimeNone
	r.TimeHeartbeat = TimeNone

	return
}

func FilterTimeNone(ts int64) *Option[int64] {
	if ts == TimeNone {
		return None[int64]()
	}

	return Some(ts)
}

func (r *TimedRecord) GetTTL() *Option[int64] {
	return FilterTimeNone(r.TTL)
}

func (r *TimedRecord) SetTTL(value int64) {
	r.TTL = value
}

func (r *TimedRecord) AddTTL(value int64) {
	r.SetTTL(r.GetTTL().GetOr(0) + value)
}

func (r *TimedRecord) GetTimeCreation() int64 {
	return r.TimeCreation
}

func (r *TimedRecord) GetTimeUpdate() *Option[int64] {
	return FilterTimeNone(r.TimeUpdate)
}

func (r *TimedRecord) SetTimeUpdate(value int64) {
	r.TimeUpdate = value
}

func (r *TimedRecord) DoTimeUpdate() {
	r.SetTimeUpdate(NowMilli())
}

func (r *TimedRecord) GetTimeRead() *Option[int64] {
	return FilterTimeNone(r.TimeRead)
}

func (r *TimedRecord) SetTimeRead(value int64) {
	r.TimeRead = value
}

func (r *TimedRecord) DoTimeRead() {
	r.SetTimeRead(NowMilli())
}

func (r *TimedRecord) GetTimeHeartbeat() *Option[int64] {
	return FilterTimeNone(r.TimeHeartbeat)
}

func (r *TimedRecord) SetTimeHeartbeat(value int64) {
	r.TimeHeartbeat = value
}

func (r *TimedRecord) DoTimeHeartbeat() {
	r.SetTimeHeartbeat(NowMilli())
}
