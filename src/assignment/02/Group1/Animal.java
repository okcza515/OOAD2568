abstract class Animal {
    int age;
    int weightInLbs;
    String gender;

    public Animal(int age, String gender, int weightInLbs) {
        this.age = age;
        this.weightInLbs = weightInLbs;
        this.gender = gender;
    }

    abstract void eat();

    abstract void sleep();

    abstract void move();
}
