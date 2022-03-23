package flags

import (
        "Libraries/libspark-msil/constants"
	flag "github.com/spf13/pflag"
)


var (
        baseConfigPath = flag.String(constants.BaseConfigPathKey, constants.BaseConfigPathDefaultValue,
                constants.BaseConfigPathUsage)
)


func BaseConfigPath() string {
        return *baseConfigPath
}
