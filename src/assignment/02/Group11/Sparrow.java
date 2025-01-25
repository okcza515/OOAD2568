//65070503408
public class Sparrow extends Animal implements Flyable {
    public Sparrow(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    public void move() {
        System.out.println("The sparrow is moving...");
    }
    public void fly() {
        System.out.println("The sparrow is flying...");
    }

}
