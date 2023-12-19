public class LogLine {
    String logLine;
    public LogLine(String logLine) {
        this.logLine = logLine;
    }

    public LogLevel getLogLevel() {
        String levelWithBracket = this.logLine.split(":")[0];
        String level = levelWithBracket.substring(1, levelWithBracket.length() - 1);
        switch(level) {
            case "TRC":
                return LogLevel.TRACE;
            case "DBG":
                return LogLevel.DEBUG;
            case "INF":
                return LogLevel.INFO;
            case "WRN":
                return LogLevel.WARNING;
            case "ERR":
                return LogLevel.ERROR;
            case "FTL":
                return LogLevel.FATAL;
            default:
                return LogLevel.UNKNOWN;
        }
    }

    public String getOutputForShortLog() {
        LogLevel logLevel = getLogLevel();
        int level = logLevel.getLevel();
        String message = this.logLine.split(": ")[1];
        return String.format("%s:%s", level, message);
    }
}
