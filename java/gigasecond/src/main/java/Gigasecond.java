import java.time.LocalDate;
import java.time.LocalDateTime;

public class Gigasecond {
    private final double gigasecond = 1e9;
    private final LocalDateTime timeAfterGigaSecond;
    public Gigasecond(LocalDate moment) {
        LocalDateTime localDateTime = moment.atStartOfDay();
        this.timeAfterGigaSecond = localDateTime.plusSeconds((long)gigasecond);
    }

    public Gigasecond(LocalDateTime moment) {
        this.timeAfterGigaSecond = moment.plusSeconds((long)gigasecond);
    }

    public LocalDateTime getDateTime() {
        return this.timeAfterGigaSecond;
    }
}
