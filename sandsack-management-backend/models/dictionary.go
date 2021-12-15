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
	1: "CREATED",
	2: "EDITED",
	3: "COMMENTED",
	4: "ACCEPTED",
	5: "DECLINED",
	6: "ASSIGNED",
}

var DictActionTypeName = map[string]int {
	"CREATED": 1,
	"EDITED": 2,
	"COMMENTED": 3,
	"ACCEPTED": 4,
	"DECLINED": 5,
	"ASSIGNED": 6,
}