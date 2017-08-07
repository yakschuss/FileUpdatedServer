package main

import (
  "time"
  "github.com/spf13/afero"
)

type ModificationInfo struct {
  modified bool
  timeModified string
}

func main() {
}

func IsModified(fileName string, checker afero.Fs) ModificationInfo {

  diff := 24 * time.Hour
  twentyFourHoursAgo := time.Now().Add(-diff)

  info, err := checker.Stat(fileName)

  if err != nil {
    return ModificationInfo{
      false,
      "not modified",
    }
  }

  return ModificationInfo{
    info.ModTime().After(twentyFourHoursAgo),
    info.ModTime().Format("2006-01-02T15:04:05.999999-07:00"),
  }
}
