package fileChecker

import (
  "time"
  "github.com/spf13/afero"
)

type ModificationInfo struct {
  Modified bool
  TimeModified string
}

type FileChecker struct {
  Checker afero.Fs
}

func (c FileChecker) IsModified(fileName string) ModificationInfo {
  diff := 24 * time.Hour
  twentyFourHoursAgo := time.Now().Add(-diff)

  info, err := c.Checker.Stat(fileName)

  if err != nil {
    return ModificationInfo{
      Modified: false,
      TimeModified: "not modified",
    }
  }

  return ModificationInfo{
    Modified: info.ModTime().After(twentyFourHoursAgo),
    TimeModified: info.ModTime().Format("2006-01-02T15:04:05.999999-07:00"),
  }
}
