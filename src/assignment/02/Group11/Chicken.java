//Intouch Krajangprateep 65070503442
public class Chicken extends Animal {
    public Chicken(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }
    @Override
    public void move()
    {
        System.out.println("The chicken is moving");
    }
}