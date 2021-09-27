package models

type Remote struct {
	XFF string `json:"x-forwarded-for"`
}

type SuccessResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

type User struct {
	First    string `bson:"first"`
	Last     string `bson:"last"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Token    string `bson:"token"`
}

type SonarProject struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type SonarConditions struct {
	ErrorThreshold string `json:"errorThreshold"`
	Metric         string `json:"metric"`
	OnLeakPeriod   bool   `json:"onLeakPeriod"`
	Operator       string `json:"operator"`
	Status         string `json:"status"`
	Value          string `json:"value"`
}

type SonarQualityGate struct {
	Conditions []SonarConditions `json:"conditions"`
}
type SimpleHook struct {
	Status string `json:"status"`
	TaskId string `json:"taskId"`
}
type WebHook struct {
	ServerUrl  string `json:"serverUrl"`
	TaskId     string `json:"taskId"`
	Status     string `json:"status"`
	Revision   string `json:"revision"`
	AnalyzedAt string `json:"analyzedAt"`
	Project    struct {
		Key  string `json:"key"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"project"`
	QualityGate struct {
		Conditions []struct {
			ErrorThreshold string `json:"errorThreshold"`
			Metric         string `json:"metric"`
			OnLeakPeriod   bool   `json:"onLeakPeriod"`
			Operator       string `json:"operator"`
			Status         string `json:"status"`
			Value          string `json:"value"`
		} `json:"conditions"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"qualityGate"`
}

type SonarProperties struct {
	BuildNumber string `json:"sonar.analysis.buildNumber"`
}

// type SonarResult struct {
// 	Status  string       `bson:"status"`
// 	Project SonarProject `bson:"project"`
// }
type SonarResult struct {
	Status string `bson:"status"`
	// Project SonarProject `bson:"project"`
}
