package common

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

const timeResolution = 1000000

func SnowFlakeInit(startTime string, machineID int64) error {
	var st time.Time
	st, err := time.Parse("2006-01-01", startTime)
	if err != nil {
		return err
	}

	snowflake.Epoch = st.UnixNano() / timeResolution
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		return err
	}

	return nil
}

func GenID() int64 {
	return node.Generate().Int64()
}
