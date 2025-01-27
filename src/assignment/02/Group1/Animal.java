abstract class Animal {
    int age;
    int weightInLbs;
    String gender;

    Animal(int age, String gender, int weightInLbs) {
        this.age = age;
        this.weightInLbs = weightInLbs;
        this.gender = gender;
    }

    void eat() {
        System.out.println("The animal is eating...");
    }

    void sleep() {
        System.out.println("The animal is sleeping...");
    }

    abstract void move();
}
