package constant

type Response struct {
	CODE        string
	DESCRIPTION string
}

var SUCCESS = Response{
	CODE:        "0000",
	DESCRIPTION: "SUCCESS",
}

var REQUEST_ERROR = Response{
	CODE:        "4001",
	DESCRIPTION: "Invalid Request",
}

var GENERIC_ERROR = Response{
	CODE:        "4002",
	DESCRIPTION: "Some Problem in Service",
}
