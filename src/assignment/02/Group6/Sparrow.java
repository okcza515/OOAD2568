// 65070501008 จิราพร วังคำหาญ
public class Sparrow extends Bird implements Flyable{
    public Sparrow(int age, String gender, double weightInLbs){
        super(age,gender,weightInLbs);
    }

    @Override
    public void fly() {
        System.out.println("Sparrow is flying.");
    }

    public void move() {
        System.out.println("Sparrow is moving.");
    }
}
