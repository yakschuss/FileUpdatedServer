package fileChecker

import (
  "os";
  "time"
)


func main() {
}

func IsModified(f string) (bool, string) {

  diff := 24 * time.Hour
  twentyFourHoursAgo := time.Now().Add(-diff)

  info, err := os.Stat(f)

  if err != nil {
    return false, "not modified"
  }

  return info.ModTime().After(twentyFourHoursAgo), info.ModTime().Format("2006-01-02T15:04:05.999999-07:00")
}
