class NeedForSpeed {
    private int distance = 0;
    private int battery = 100;
    private int speed;
    private int batteryDrain;
    public int getSpeed() {
        return this.speed;
    }

    public int getBattery() {
        return this.battery;
    }

    public int getBatteryDrain() {
        return this.batteryDrain;
    }

    public int getDistance() {
        return this.distance;
    }
    NeedForSpeed(int speed, int batteryDrain) {
        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public boolean batteryDrained() {
        return this.battery <= 0;
    }

    public int distanceDriven() {
        return this.distance;
    }

    public void drive() {
        if (!batteryDrained() && this.battery >= this.batteryDrain) {
            this.distance += this.speed;
            this.battery -= this.batteryDrain;
        }
    }

    public static NeedForSpeed nitro() {
        return new NeedForSpeed(50, 4);
    }
}

class RaceTrack {
    private int distance;
    RaceTrack(int distance) {
        this.distance = distance;
    }

    public boolean tryFinishTrack(NeedForSpeed car) {
        return (int) (car.getBattery() / car.getBatteryDrain() * car.getSpeed() + car.getDistance()) >= this.distance;
    }
}
