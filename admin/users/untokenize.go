package admin

import (
	"github.com/dracory/userstore"
	"github.com/dracory/userstore/admin/shared"
)

func userUntokenize(config shared.Config, user userstore.UserInterface) (columnNameValueMap map[string]string, err error) {
	if len(config.TokenizedColumns) < 1 {
		return columnNameValueMap, nil
	}

	columnNameTokenMap := map[string]string{}

	for _, columnName := range config.TokenizedColumns {
		columnNameTokenMap[columnName] = user.Get(columnName)
	}

	return config.TokensRead(columnNameTokenMap)
}
