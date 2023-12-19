public class ProductionRemoteControlCar implements RemoteControlCar, Comparable<ProductionRemoteControlCar> {
    int distance = 0;
    int numberOfVictories = 0;
    ProductionRemoteControlCar() {
        
    }
    public void setNumberOfVictories(int victories) {
        this.numberOfVictories = victories;
    }

    public int getNumberOfVictories() {
        return this.numberOfVictories;
    }
    public void drive() {
        this.distance += 10;
    }

    public int getDistanceTravelled() {
        return this.distance;
    }

    public int compareTo(ProductionRemoteControlCar car) {
        return car.getNumberOfVictories() - this.numberOfVictories;
    }
} 
