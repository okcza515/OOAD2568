abstract public class Animal {
    private int age;
    private String gender;
    private double weightInLbs;
    Animal(int age, String gender, double weightInLbs){
        this.age = age;
        this.gender = gender;
        this.weightInLbs = weightInLbs;
    }
    void eat(){
        System.out.println("Animal is eating.");
    }
    void sleep(){
        System.out.println("Animal is sleeping.");
    }
    abstract void move();
}
