
// 65070501001 Kantapong Vongpanich
abstract class Animal {

    protected int age;
    protected String gender;
    protected double weightInLbs;

    public Animal(int _age, String _gender, double _weight) {
        this.age = _age;
        this.gender = _gender;
        this.weightInLbs = _weight;
    }

    abstract void eat();

    abstract void move();

    abstract void sleep();
}

interface Flyable {
    public void fly();
}
// 65070501065 Kamolpop Poonsawat

class Fish extends Animal {

    public Fish(int age, String gender, double weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    public void eat() {
        System.out.println("Fishy is eatting ngum ngum");
    }

    @Override
    public void move() {
        System.out.println("Fishy is moving");
    }

    @Override
    public void sleep() {
        System.out.println("Fishy is Sleepping");
    }

    public void swim() {
        System.out.println("Fishy is swimming");
    }

}
// Chicken Tonhom
// Wavie Flyable
// Kanasorn Sparrow

// 65070501023 Thanaphol Thangthaweesuk
class Bird extends Animal {

    public Bird(int _age, String _gender, double _weight) {
        super(_age, _gender, _weight);
    }

    @Override
    public void eat() {
        System.out.println("Bird is eatting");
    }

    @Override
    public void move() {
        System.out.println("Bird is moving");
    }

    @Override
    public void sleep() {
        System.out.println("Bird is Sleepping");
    }

}

public class Zoo {

    public static void main(String[] args) {
        Animal fish1 = new Fish(1, "M", 2);
        Animal bird1 = new Bird(1, "F", 1);
        Animal chicken1 = new Chicken(1, "F", 2);
        Animal sparrow1 = new Sparrow(1, "M", 4);
        Flyable sparrow2 = new Sparrow(1, "M", 4);

        fish1.eat();
        fish1.sleep();
        fish1.move();

        bird1.eat();
        bird1.sleep();
        bird1.move();

        chicken1.eat();
        chicken1.sleep();
        chicken1.move();

        sparrow1.eat();
        sparrow1.sleep();
        sparrow1.move();

        moveAnimals(fish1);
        moveAnimals(bird1);
        moveAnimals(chicken1);
        moveAnimals(sparrow1);

        sparrow2.fly();
    }

    public static void moveAnimals(Animal animal) {
        animal.move();
    }

}
