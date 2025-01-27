public class Zoo {
    public static void main(String[] args) {
        Animal fish = new Fish(1, "M", 2);
        fish.eat();
        fish.sleep();
        fish.swim();
        fish.move();
        Animal bird = new Bird(3, "F", 5);
        bird.eat();
        bird.sleep();
        bird.move();
        Animal chicken = new Chicken(9, "M", 6);
        chicken.eat();
        chicken.sleep();
        chicken.move();
        Animal sparrow = new Sparrow(2, "F", 1);
        sparrow.eat();
        sparrow.sleep();
        sparrow.move();
        sparrow.fly();
    }
}
