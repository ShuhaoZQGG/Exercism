import java.lang.Math;
class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {
        double sum = (1 + input) * input / 2;
        double squareOfSum = Math.pow(sum, 2);
        return (int)squareOfSum;
    }

    int computeSumOfSquaresTo(int input) {
        int sumOfSquares = input * (input + 1) * (2*input + 1) / 6;
        return sumOfSquares;
    }

    int computeDifferenceOfSquares(int input) {
        return Math.abs(computeSumOfSquaresTo(input) - computeSquareOfSumTo(input));
    }

}
