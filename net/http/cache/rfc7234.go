package cache

const (
	WarnStale                 = 110
	WarnRevalidationFailed    = 111
	WarnDisconnected          = 112
	WarnHeuristicExpiration   = 113
	WarnMisc                  = 199
	WarnTransformationApplied = 214
	WarnMiscPersist           = 299
)

var warningText = map[int]string{
	WarnStale:                 "Response is Stale",
	WarnRevalidationFailed:    "Revalidation Failed",
	WarnDisconnected:          "Disconnected Operation",
	WarnHeuristicExpiration:   "Heuristic Expiration",
	WarnMisc:                  "Miscellaneous Warning",
	WarnTransformationApplied: "Transformation Applied",
	WarnMiscPersist:           "Miscellaneous Persist Warning",
}

type deltaSeconds int64
