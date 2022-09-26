package errortool

var (
	Codes = Define()
	ErrDB = Codes.Plugin(DBErrorPackage).(dbError)
)
