// 65070503412
public class Fish extends Animal {
    public Fish(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }
    public void swim() {
        System.out.println("The fish is swimming...");
    }

    @Override
    public void move() {
        swim();
    }
}
