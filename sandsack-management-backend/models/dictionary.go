package models


var DictStatus = map[int]string {
	1: "PENDING",
}

var DictStatusName = map[string]int {
	"PENDING": 1,
}

var DictPriority = map[int]string {
	1: "HIGH",
	2: "MIDDLE",
	3: "LOW",
}

var DictPriorityName = map[string]int {
	"HIGH": 1,
	"MIDDLE": 2,
	"LOW": 3,
}

var DictActionType = map[int]string {
	1: "CREATE_ORDER",
}

var DictActionTypeName = map[string]int {
	"CREATE_ORDER": 1,
}