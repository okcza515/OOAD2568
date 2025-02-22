// group 2
abstract class Animal{
	private int age;
	private String gender;
	private double weightInLbs;

	public Animal(int age, String gender, double weightInLbs) {
		this.age = age;
		this.gender = gender;
		this.weightInLbs = weightInLbs;
	}

	abstract void eat();
	abstract void sleep();
	abstract void move();
}

class Fish extends Animal{
	public Fish(int age, String gender, double weightInLbs){
		super(age, gender, weightInLbs);
	}

	@Override
	public void eat(){
		System.out.println("Eating");
	}

	@Override
	public void sleep(){
		System.out.println("Sleeping");
	}

	@Override
	public void move(){
		System.out.println("Moving");
	}

	public void swim(){
		System.out.println("Swimming");
	}
}

class Bird extends Animal{
	public Bird(int age, String gender, double weightInLbs){
		super(age, gender, weightInLbs);
	}

	@Override
	public void eat(){
		System.out.println("Eating");
	}

	@Override
	public void sleep(){
		System.out.println("Sleeping");
	}

	@Override
	public void move(){
		System.out.println("Moving");
	}
}

interface Flyable {
	void fly();
}

class Chicken extends Animal{
	public Chicken(int age, String gender, double weightInLbs){
		super(age, gender, weightInLbs);
	}

	@Override
	public void eat(){
		System.out.println("Eating");
	}

	@Override
	public void sleep(){
		System.out.println("Sleeping");
	}

	@Override
	public void move(){
		System.out.println("Moving");
	}
}

class Sparrow extends Animal implements Flyable{
	public Sparrow(int age, String gender, double weightInLbs){
		super(age, gender, weightInLbs);
	}

	@Override
	public void eat(){
		System.out.println("Eating");
	}

	@Override
	public void sleep(){
		System.out.println("Sleeping");
	}

	@Override
	public void fly(){
		System.out.println("Flying");
	}

}

public class Zoo {
	public static void main(String[] args) {
		
		Animal fish1 = new Fish(1, "M", 2);
		Animal bird1 = new Bird(1,"F",1);
		Animal chicken1 = new Chicken(1,"F",2);
		Animal sparrow1 = new Sparrow(1, "M", 4);
		Flyable sparrow2 = new Sparrow(1,"M",4);
		
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