package media

type SiteConfig struct {
	Key   string `db:"key" json:"key"`
	Cover string `db:"cover" json:"cover"`
}
