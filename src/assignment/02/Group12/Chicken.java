//65070501038 Puntharee Roongprasert
public class Chicken extends Animal{

    public Chicken(int age, String gender, int weightInLbs)
    {
        super(age, gender, weightInLbs);
    }

    @Override
    public void eat(){
        System.out.println("The chicken is eating");
    }
    @Override
    public void sleep(){
        System.out.println("The chicken is sleeping");
    }
    @Override
    public void move(){
        System.out.println("The chicken is moving");
    }

}


