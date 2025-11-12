package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Epss The Exploit Prediction Scoring System (EPSS) object describes the estimated
// probability a vulnerability will be exploited. EPSS is a community-driven effort
// to combine descriptive information about vulnerabilities (CVEs) with evidence of
// actual exploitation in-the-wild. (https://www.first.org/epss/ EPSS).
type Epss struct {
	// The timestamp indicating when the EPSS score was calculated.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The timestamp indicating when the EPSS score was calculated.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The EPSS score's percentile representing relative importance and ranking of
	// the score in the larger EPSS dataset.
	Percentile float64 `json:"percentile,omitempty"`

	// The EPSS score representing the probability [0-1] of exploitation in the
	// wild in the next 30 days (following score publication).
	Score string `json:"score,omitempty"`

	// The version of the EPSS model used to calculate the score.
	Version string `json:"version,omitempty"`
}
