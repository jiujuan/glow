package utils

import "time"

func DateFormat(inttime int, format string) string {
    if format == "" {
        format = "2006-01-02 15:04"
    }
    return time.Unix(int64(inttime), 0).Format(format)
}
