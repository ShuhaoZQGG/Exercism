public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        int rate = 221;
        double successRate = 1.0;
        if (speed >= 5 && speed <= 8) {
            successRate = 0.9;
        } else if (speed >= 9 && speed < 10) {
            successRate = 0.8;
        } else if (speed >= 10) {
            successRate = 0.77;
        }

        return rate * successRate * speed;
    }

    public int workingItemsPerMinute(int speed) {
        int rate = 221;
        int minutesPerHour = 60;
        double successRate = 1.0;
        if (speed >= 5 && speed <= 8) {
            successRate = 0.9;
        } else if (speed >= 9 && speed < 10) {
            successRate = 0.8;
        } else if (speed >= 10) {
            successRate = 0.77;
        }
        int ratePerMinute = (int)(rate * speed * successRate / minutesPerHour);
        return ratePerMinute;
    }
}
