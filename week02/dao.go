package week02

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Site struct {
	ID     string
	Region string
	AZ     string
}

func GetSite(siteID string) (*Site, error) {
	site := &Site{}

	err := GetDB().QueryRow("select * from t_sites where id = ?", siteID).Scan(site)
	// 习惯不把 ErrNoRows 作为 error 上抛
	// 上层调用者通过判断返回值 site 是否为 nil，来做相应的业务处理
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "db.Query err")
	}

	return site, nil
}
