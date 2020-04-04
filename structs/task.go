package structs

type Task struct {
	ID          int    `json:"id"`
	TaskKey     string `json:"taskKey"`
	Reporter    string `json:"reporter"`
	Time        string `json:"time"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Owner       string `json:"owner"`
}

type ListReponse struct {
	Mine      []Task `json:"myClaimedTasks"`
	Others    []Task `json:"otherClaimedTasks"`
	Unclaimed []Task `json:"unclaimedTasks"`
}
