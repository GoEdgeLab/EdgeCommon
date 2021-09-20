package schedulingconfigs

// CandidateInterface 候选对象接口
type CandidateInterface interface {
	// CandidateWeight 权重
	CandidateWeight() uint

	// CandidateCodes 代号
	CandidateCodes() []string
}
