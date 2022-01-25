package models

var DictStatus = map[int]string{
	1:  "ANSTEHEND",
	2:  "STORNIERT",
	3:  "WEITERGELEITET BEI EINSATZABSCHNITT",
	4:  "WEITERGELEITET BEI HAUPTABSCHNITT",
	5:  "ABGELEHNT BEI EINSATZABSCHNITT",
	6:  "ABGELEHNT BEI HAUPTABSCHNITT",
	7:  "ABGELEHNT BEI EINSATZLEITER",
	8:  "AKZEPTIERT",
	9:  "AUF DEM WEG",
	10: "GELIEFERT",
}

var DictStatusName = map[string]int{
	"ANSTEHEND":                           1,
	"STORNIERT":                           2,
	"WEITERGELEITET BEI EINSATZABSCHNITT": 3,
	"WEITERGELEITET BEI HAUPTABSCHNITT":   4,
	"ABGELEHNT BEI EINSATZABSCHNITT":      5,
	"ABGELEHNT BEI HAUPTABSCHNITT":        6,
	"ABGELEHNT BEI EINSATZLEITER":         7,
	"AKZEPTIERT":                          8,
	"AUF DEM WEG":                         9,
	"GELIEFERT":                           10,
}

var DictPriority = map[int]string{
	1: "HOCH",
	2: "MITTEL",
	3: "NIEDRIG",
}

var DictPriorityName = map[string]int{
	"HOCH":    1,
	"MITTEL":  2,
	"NIEDRIG": 3,
}

var DictActionType = map[int]string{
	1: "CREATED",
	2: "EDITED",
	3: "COMMENTED",
	4: "ACCEPTED",
	5: "DECLINED",
	6: "ASSIGNED",
	7: "CONFIRMED DELIVERY",
}

var DictActionTypeName = map[string]int{
	"CREATED":            1,
	"EDITED":             2,
	"COMMENTED":          3,
	"ACCEPTED":           4,
	"DECLINED":           5,
	"ASSIGNED":           6,
	"CONFIRMED DELIVERY": 7,
}

var DictPermission = map[int]string{
	1: "CAN VIEW",
	2: "CAN CONFIRM DELIVERY",
	3: "CAN EDIT",
	4: "CAN DECLINE",
	5: "CAN ACCEPT",
	6: "CAN COMMENT",
	7: "CAN ASSIGN TO",
}

var DictPermissionName = map[string]int{
	"CAN VIEW":             1,
	"CAN CONFIRM DELIVERY": 2,
	"CAN EDIT":             3,
	"CAN DECLINE":          4,
	"CAN ACCEPT":           5,
	"CAN COMMENT":          6,
	"CAN ASSIGN TO":        7,
}

var DictBranch = map[int]string{
	1: "Mollnhof",
	2: "Einsatzleiter",
	3: "Hauptabschnitt",
	4: "Einsatzabschnitt",
	5: "Unterabschnitt",
}

var DictBranchName = map[string]int{
	"Mollnhof":         1,
	"Einsatzleiter":    2,
	"Hauptabschnitt":   3,
	"Einsatzabschnitt": 4,
	"Unterabschnitt":   5,
}
