package global

const (
	// mapping fields
	Mapping = "mapping"
	Id      = "id"
	Host    = "host"
	Uri     = "uri"
	Method  = "method"

	// json output fields
	RequestInfo = "requestInfo"
	Message     = "message"
	ErrMessage  = "error_message"
	Status      = "status"
	Ok          = "ok"

	// params
	UserId = "userID"

	//logger name
	LoggerName = "com.poalim.bank.hackathon.login.logger"

	//template replacements (for logger)
	OldTemplate = `"time":"${time_rfc3339_nano}"`
	NewTemplate = `"prefix":"com.poalim.bank.hackathon.login.logger","time":"${time_rfc3339_nano}"`
)
