public class Main {
    public static void main(String[] args) {
        SciTeacher S1 = new SciTeacher();
        MathTeacher M1 = new MathTeacher();
        SubTeacher Sb1 = new SubTeacher();

        S1.teach();
        S1.performOtherResponsibilities();
        M1.teach();
    }
}
