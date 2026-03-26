package domain

type SettingModel struct {
	Key       string `gorm:"type:varchar(100);primary_key;"`
	Value     string `gorm:"type:text;"`
	ValueType string `gorm:"type:varchar(20);not null;"` // string, boolean, number
}

func (setting SettingModel) TableName() string {
	return "setting"
}

const (
	SettingKeySiteTitle = "site_title"
	SettingKeyDarkMode  = "dark_mode"
)

func GetPredefinedSettings() []SettingModel {
	return []SettingModel{
		{
			Key:       SettingKeySiteTitle,
			Value:     "LibrePM",
			ValueType: "string",
		},
		{
			Key:       SettingKeyDarkMode,
			Value:     "auto",
			ValueType: "string", // or boolean if we use "true"/"false"
		},
	}
}
