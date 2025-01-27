public abstract class Animal {

    public Animal(int age, String gender, int weightInLbs) {
    }

    public void eat() {
        System.out.println("Animal is eating.");
    }

    public void sleep() {
        System.out.println("Animal is sleeping.");
    }

    public abstract void move();
}
