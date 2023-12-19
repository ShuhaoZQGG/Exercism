class Scrabble {
    private final String word;
    Scrabble(String word) {
        this.word = word;
    }

    int getScore() {
        int result = 0;
        for (int i = 0; i < word.length(); i++) {
            char c = word.charAt(i);
            switch(Character.toUpperCase(c)) {
                case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
                    result ++;
                    break;
                case 'D', 'G':
                    result += 2;
                    break;
                case 'B', 'C', 'M', 'P':
                    result += 3;
                    break;
                case 'F', 'H', 'V', 'W', 'Y':
                    result += 4;
                    break;
                case 'K':
                    result += 5;
                    break;
                case 'J', 'X':
                    result += 8;
                    break;
                case 'Q', 'Z':
                    result += 10;
                    break;
            }
        }
        return result;
    }

}
