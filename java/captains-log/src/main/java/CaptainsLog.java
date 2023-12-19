import java.util.Random;

class CaptainsLog {

    private static final char[] PLANET_CLASSES = new char[]{'D', 'H', 'J', 'K', 'L', 'M', 'N', 'R', 'T', 'Y'};

    private Random random;

    CaptainsLog(Random random) {
        this.random = random;
    }

    char randomPlanetClass() {
        int randomPos = this.random.nextInt(this.PLANET_CLASSES.length);
        return this.PLANET_CLASSES[randomPos];
    }

    String randomShipRegistryNumber() {
        String prefix = "NCC-";
        int randomNum = 1000 + this.random.nextInt(9000);
        return prefix + randomNum;
    }

    double randomStardate() {
        return 41000 + 1000 * this.random.nextDouble();
    }
}
