# pomodoro-report
Generates daily and weekly pomodoro totals.

This tool takes a file containing pomodoro data in the following format, and generates a
report to stdout containing the date, daily total and weekly total of pomodoros.

The expected format for the pomodoro data is:

```
20170901@1020 : p
20170901@1024 : s
20170901@1049 : p
20170901@1054 : s
20170901@1121 : p
20170901@1147 : p
20170901@1153 : s
20170901@1508 : p
20170901@1535 : p
```

The times are ignored, as are lines that don't end with "p" (in my usage, "s" represents
a short break, and "b" represents a long break).
