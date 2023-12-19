public class Lasagna {
    // TODO: define the 'expectedMinutesInOven()' method
    int expectedMinutes = 40;
    public int expectedMinutesInOven() {
        return expectedMinutes;
    }
    // TODO: define the 'remainingMinutesInOven()' method
    public int remainingMinutesInOven(int inOvenMinutes) {
        return expectedMinutes - inOvenMinutes;
    }
    // TODO: define the 'preparationTimeInMinutes()' method
    public int preparationTimeInMinutes(int layer) {
        return 2 * layer;
    }
    // TODO: define the 'totalTimeInMinutes()' method
    public int totalTimeInMinutes(int layer, int inOvenMinutes) {
        return 2 * layer + inOvenMinutes;
    }
}
