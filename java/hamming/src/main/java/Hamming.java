import java.util.stream.IntStream;
public class Hamming {
    private final String leftStrand;
    private final String rightStrand;
    public Hamming(String leftStrand, String rightStrand) {
        if (leftStrand.length() != rightStrand.length()) {
            throw new IllegalArgumentException("strands must be of equal length");
        }
        this.leftStrand = leftStrand;
        this.rightStrand = rightStrand;
    }

    public int getHammingDistance() {
        long counter = IntStream
            .range(0, this.leftStrand.length())
            .filter(i -> this.leftStrand.charAt(i) != this.rightStrand.charAt(i))
            .count();
        return (int) counter;
    }
}
