public class Zoo_Fish {

    static abstract class Animal {
        int age;
        String gender;
        int weight;

        public Animal(int age, String gender, int weight) {
            this.age = age;
            this.gender = gender;
            this.weight = weight;
        }

        public void eat() {
            System.out.println("This animal is eating.");
        }

        public void sleep() {
            System.out.println("This animal is sleeping.");
        }

        public abstract void move();
    }

    static class Fish extends Animal {
        public Fish(int age, String gender, int weight) {
            super(age, gender, weight);
        }

        public void swim() {
            System.out.println("Fish is swimming.");
        }

        public void move() {
            swim();
        }
    }

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
