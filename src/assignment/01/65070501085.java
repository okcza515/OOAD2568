class Puwadech {
    static int removeCharacter(String str) {
        int size = str.length() - 1;
        int integer = 0;
        int count = 0;

        for (int i = size; i >= 0; i--) {
            int num = 0;
            if (Character.isDigit(str.charAt(i))) {
                num = (int) ((str.charAt(i) - '0') * Math.pow(10, count));
                count++;
            }
            integer += num;
        }
        return integer;
    }

    public static void main(String[] args) {
        String n = "1a2a3", m = "45f6";
        int num = removeCharacter(n) + removeCharacter(m);
        System.out.println(removeCharacter(n) + "\n +");
        System.out.println(removeCharacter(m) + "\n =");
        System.out.println(num);
    }

}
