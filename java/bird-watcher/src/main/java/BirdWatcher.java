
class BirdWatcher {
    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public int[] getLastWeek() {
        return new int[] { 0, 2, 5, 3, 7, 8, 4};
    }

    public int getToday() {
        return this.birdsPerDay[this.birdsPerDay.length - 1];
    }

    public void incrementTodaysCount() {
        this.birdsPerDay[this.birdsPerDay.length - 1] += 1;
    }

    public boolean hasDayWithoutBirds() {
        for(int i = 0; i < this.birdsPerDay.length; i++) {
            if (this.birdsPerDay[i] == 0) {
                return true;
            }
        }
        return false;
    }

    public int getCountForFirstDays(int numberOfDays) {
        int counter = 0;
        numberOfDays = Math.min(numberOfDays, this.birdsPerDay.length);
        for (int i = 0; i < numberOfDays; i++) {
            counter += this.birdsPerDay[i];
        }
        return counter;
    }

    public int getBusyDays() {
        int busyThreshold = 5;
        int counter = 0;
        for (int i = 0; i < this.birdsPerDay.length; i++) {
            if (this.birdsPerDay[i] >= busyThreshold) {
                counter += 1;
            }
        }
        return counter;
    }
}
