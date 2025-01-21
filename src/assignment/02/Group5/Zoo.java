// Animal class (Abstract class) by Natchanon 65070501018
abstract class Animal {
    int age;
    String gender;
    double weightInLbs;

    public Animal(int age, String gender, double weightInLbs) {
        this.age = age;
        this.gender = gender;
        this.weightInLbs = weightInLbs;
    }

    public void eat() {
        System.out.println("nom nom nomm nomm, This animal is enjoy eating food.");
    }

    public void sleep() {
        System.out.println("kok feeee...ZzzzZZzz, This animal is sleeping peacefully.");
    }

    public abstract void move();
}

// Fish class by Napat 65070501074
class Fish extends Animal {
    public Fish(int age, String gender, double weightInLbs) {
        super(age, gender, weightInLbs);
    }

    public void swim() {
        System.out.println("The fish is swimming very fast.");
    }

    @Override
    public void move() {
        swim();
    }
}

// Bird class by Chatchanan 65070501014
class Bird extends Animal {
    public Bird(int age, String gender, double weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    public void move() {
        System.out.println("the brid is walking into a food");
    }

}

// Flyable (interface class)
interface Flyable {
    void fly();
}
// Chicken class by Nutchanon 65070501075
class Chicken extends Bird {
    public Chicken(int age, String gender, double weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    public void move() {
        System.out.println("jae jae jae! The chicken is walking around.");
    }
}

// Sparrow class by Rattipong 65070501048
class Sparrow extends Bird implements Flyable {
    public Sparrow(int age, String gender, double weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    public void fly() {
        System.out.println("pub pub pub... The sparrow is flying in to the moonnnn.");
    }
}

public class Zoo {
    public static void main(String[] args) {
        Animal fish1 = new Fish(1, "M", 2);
        Animal bird1 = new Bird(2, "F", 3);
        Animal chicken1 = new Chicken(1, "F", 4);
        Animal sparrow1 = new Sparrow(1, "M", 1);
        Flyable sparrow2 = new Sparrow(1, "M", 1);

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