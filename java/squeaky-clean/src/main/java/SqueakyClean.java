class SqueakyClean {
    static String clean(String identifier) {
        identifier = replaceSpacesWithUnderscores(identifier);
        identifier = replaceControlChars(identifier);
        identifier = kebabToCamel(identifier);
        identifier = omitNonLetters(identifier);
        identifier = removeLowerCaseGreeks(identifier);
        return identifier;
    }

    static String replaceSpacesWithUnderscores(String identifier) {
        return identifier.replace(' ', '_');
    }

    static String replaceControlChars(String identifier) {
        StringBuilder str = new StringBuilder();
        for (int i = 0; i < identifier.length(); i++) {
            char c = identifier.charAt(i);
            if (Character.isISOControl(c)) {
                str.append("CTRL");
            } else {
                str.append(c);
            }
        }
        return str.toString();
    }

    static String kebabToCamel(String identifier) {
        StringBuilder str = new StringBuilder();
        String[] array = identifier.split("-");
        str.append(array[0]);
        for (int i = 1; i < array.length; i++) {
            char[] charArray = array[i].toCharArray();
            boolean camelled = false;
            for (int j = 0; j < charArray.length; j++) {
                char c = charArray[j];
                if (Character.isLetter(c)) {
                    if (camelled) {
                        str.append(c);
                    } else {
                        str.append(Character.toUpperCase(c));
                        camelled = true;
                    }
                } else {
                    continue;
                }
            } 
        }
        return str.toString();
    }

    static String omitNonLetters(String identifier) {
        StringBuilder str = new StringBuilder();
        for (int i = 0; i < identifier.length(); i++) {
            char c = identifier.charAt(i);
            if (!Character.isLetter(c) && c != '_') {
                continue;
            } else {
                str.append(c);
            }
        }
        return str.toString();
    }

    static String removeLowerCaseGreeks(String identifier) {
        return identifier.replaceAll("[\u03b1-\u03c9]", "");
    }
}
