package datatype

type BagsList struct {
	Type string `json:"@type"`
	Bags []Bag  `json:"torrents"`
}

type Bag struct {
	Type           string `json:"@type"`
	Hash           string `json:"hash"`
	Flags          int    `json:"flags"`
	TotalSize      string `json:"total_size"`
	Description    string `json:"description"`
	FilesCount     string `json:"files_count"`
	IncludedSize   string `json:"included_size"`
	DirName        string `json:"dir_name"`
	DownloadedSize string `json:"downloaded_size"`
	// RootDir		  string `json:"root_dir"`			// No need to parse this
	ActiveDownload bool   `json:"active_download"`
	ActiveUpload   bool   `json:"active_upload"`
	Completed      bool   `json:"completed"`
	DownloadSpeed  string `json:"download_speed"`
	UploadSpeed    string `json:"upload_speed"`
	FatalError     string `json:"fatal_error"`
}
