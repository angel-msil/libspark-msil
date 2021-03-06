package flags

import (
        "github.com/angel-msil/libspark-msil/constants"
	flag "github.com/spf13/pflag"
)


var (
        baseConfigPath = flag.String(constants.BaseConfigPathKey, constants.BaseConfigPathDefaultValue,
                constants.BaseConfigPathUsage)
)


func BaseConfigPath() string {
        return *baseConfigPath
}
