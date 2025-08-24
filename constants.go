package env

// Any of the following values are considered true: "true", "True", "TRUE", "T", "t", "1", "yes", "Yes", "YES"
// Any of the following values are considered false: "false", "False", "FALSE", "F", "f", "0", "no", "No", "NO"
// Any other value returns the default value.

const (
	// True values
	// - "true", "True", "TRUE"
	// - "T", "t"
	// - "1" (as well as any other positive number)
	// - 0.01 (as well as any other positive float)
	// - "yes", "Yes", "YES"
	// - "on", "On", "ON"
	TrueValues = "true,True,TRUE,T,t,1,yes,Yes,YES,on,On,ON"

	// False values
	// - "false", "False", "FALSE"
	// - "F", "f"
	// - "0", (as well as any other negative number)
	// - 0.00 (as well as any other negative float)
	// - "off", "Off", "OFF"
	FalseValues = "false,False,FALSE,F,f,0,no,No,NO,off,Off,OFF"
)
