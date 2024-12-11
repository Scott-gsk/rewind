package hooks

import (
	"github.com/WeOps-Lab/rewind/lib/pkgs/database"
	"node-manager/global"
)

func (e ExampleAppHooks) InstallDataSource() {
	global.DBClient = database.NewDataBaseInstanceFromEnv(true)
}
