//Kritsanaphong Thaworana 65070503403
public class Bird extends Animal {
    public Bird(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    public void move() {
        System.out.println("The Bird is moving...");
    }
}