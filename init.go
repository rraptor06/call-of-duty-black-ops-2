package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/PretendoNetwork/plogger-go"
)

func init() {
    globals.Logger = plogger.NewLogger()
}
