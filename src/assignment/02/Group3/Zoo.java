
// 65070501001 Kantapong
abstract class Animal {

    protected int age;
    protected String gender;
    protected double weightInLbs;

    public void Animal(int _age, String _gender, double _weight) {
        this.age = _age;
        this.gender = _gender;
        this.weightInLbs = _weight;
    }

    abstract void eat();

    abstract void move();

    abstract void sleep();
}

// Fish Cake
// Chicken Tonhom
// Wavie Flyable
// Kanasorn Sparrow
// Q Bird
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
