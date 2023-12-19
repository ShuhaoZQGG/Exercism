class CalculatorConundrum {
    public String calculate(int operand1, int operand2, String operation) {
        int result = 0;
        try {
            switch (operation) {
            case "+":
                result = operand1 + operand2;
                break;
            case "*":
                result = operand1 * operand2;
                break;
            case "/":
                result = operand1 / operand2;
                break;
            case "":
                throw new IllegalArgumentException("Operation cannot be empty");
            case null:
                throw new IllegalArgumentException("Operation cannot be null");
            default:
                throw new IllegalOperationException(String.format("Operation '%s' does not exist", operation));
            }
        } catch(ArithmeticException e) {
            throw new IllegalOperationException("Division by zero is not allowed", e);
        }

        return String.format("%s %s %s = %s", operand1, operation, operand2,  result);
    }
}
