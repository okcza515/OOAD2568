abstract class Animal { //65070501053
    private int age;
    private String gender;
    private int weightInLbs;

    public Animal(int age, String gender, int weightInLbs) {
        this.age = age;
        this.gender = gender;
        this.weightInLbs = weightInLbs;
    }

    public void eat() {
        System.out.println("This " + this.getClass().getName() + " is eating.");
    }

    public void sleep() {
        System.out.println("This " + this.getClass().getName() + " is sleeping.");
    }

    public abstract void move();
}

//by 65070501042
class Fish extends Animal {
    public Fish(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }

    public void swim() {
        System.out.println("The fish is swimming.");
    }

    @Override
    public void move() {
        swim();
    }
}

//by 65070501085
interface Flyable {
    void fly();
}

//by 65070501085
class Bird extends Animal {
    public Bird(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    public void move() {
        System.out.println("This bird is flying.");
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