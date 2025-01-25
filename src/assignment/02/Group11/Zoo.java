abstract class Animal {
    private int age;
    private String gender;
    private int weightInLbs;

    public Animal(int age, String gender, int weightInLbs){
        this.age = age;
        this.gender = gender;
        this.weightInLbs = weightInLbs;
    }

    void eat() {
        System.out.println("The animal is eating...");
    }

    void sleep() {
        System.out.println("The animal is sleeping...");
    }
    abstract void move();
}

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

public class Zoo {
    public static void main(String[] args) {
        Animal fish1 = new Fish(1, "M", 2);

        fish1.eat();
        fish1.sleep();
        fish1.move();

        moveAnimals(fish1);
    }
    public static void moveAnimals(Animal animal) {
        animal.move();
    }
}

