import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Locale;
class AppointmentScheduler {
    public LocalDateTime schedule(String appointmentDateDescription) {
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("MM/dd/yyyy HH:mm:ss");
        LocalDateTime parsed = LocalDateTime.parse(appointmentDateDescription, formatter);
        return parsed;
    }

    public boolean hasPassed(LocalDateTime appointmentDate) {
        LocalDateTime now = LocalDateTime.now();
        return appointmentDate.isBefore(now);
    }

    public boolean isAfternoonAppointment(LocalDateTime appointmentDate) {
        int hour = appointmentDate.getHour();
        return hour >= 12 && hour < 18;
    }

    public String getDescription(LocalDateTime appointmentDate) {
        String prefix = "You have an appointment on ";
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("EEEE, MMMM d, yyyy, 'at' h:mm a", Locale.US);
        String postfix = ".";
        String formattedString = appointmentDate.format(formatter);
        return prefix + formattedString + postfix;
    }

    public LocalDate getAnniversaryDate() {
        return LocalDate.of(2023, 9, 15);
    }}
