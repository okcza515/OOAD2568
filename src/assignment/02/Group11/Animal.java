//abstract class
abstract public class Animal {

    public Animal(int age, String gender, int weightInLbs){}

    void eat() {
        System.out.println("The animal is eating...");
    }

    void sleep() {
        System.out.println("The animal is sleeping...");
    }
    abstract void move();
}
