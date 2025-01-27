public class Bird extends Animal {
    public Bird(int age, String gender, double weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    public void move() {
        System.out.println("Bird is moving.");
    }
}
