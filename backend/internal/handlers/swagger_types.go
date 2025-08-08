package handlers

type EnqueueRequest struct {
	Queue string `json:"queue"`
	Item  string `json:"item"`
}

type DequeueRequest struct {
	Queue string `json:"queue"`
}
