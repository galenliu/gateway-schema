package gateway_schema

type UserProfile struct {
	BaseDir        string `validate:"required" json:"base_dir"`
	DataDir        string `validate:"required" json:"data_dir"`
	AddonsDir      string `validate:"required" json:"addons_dir"`
	ConfigDir      string `validate:"required" json:"config_dir"`
	UploadDir      string `validate:"required" json:"upload_dir"`
	MediaDir       string `validate:"required" json:"media_dir"`
	LogDir         string `validate:"required" json:"log_dir"`
	GatewayVersion string
}

type Preferences struct {
	Language string `validate:"required" json:"language"`
	Units    Units  `validate:"required" json:"units"`
}

type Units struct {
	Temperature string `validate:"required" json:"temperature"`
}
