public class Sparrow extends Animal implements Flyable{

    public Sparrow (int age, String gender, int weightInLbs)
    {
        super(age, gender, weightInLbs);
    }
    @Override
    public void eat()
    {
        System.out.println( "The sparrow is eating");
    }
    @Override
    public void sleep()
    {
        System.out.println( "The sparrow is sleeping");
    }
    @Override
    public void move()
    {
        System.out.println( "The sparrow is moving");
    }
    @Override
    public void fly()
    {
        System.out.println( "The sparrow is flying");
    }
}
//65070503436 Sinsorn Chaithavornkit
