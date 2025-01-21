abstract class Animal { ///Warapol Pratumta 65070503466
    protected int age;
    protected String gender;
    protected double weightInLbs;

    public Animal(int age, String gender, double weightInLbs) {
        this.age = age;
        this.gender = gender;
        this.weightInLbs = weightInLbs;
    }

    public void eat() {
        System.out.println("This animal is eating...");
    }

    public void sleep() {
        System.out.println("This animal is sleeping...");
    }

    // Abstract method to be implemented by subclasses
    public abstract void move();
}

//Fish //65070503469 Sawitt Ngamvilaisiriwong
class Fish extends Animal{
	public Fish(int age, String gender, double weightInLbs){
		super(age ,gender, weightInLbs);
	}

	public void swim(){
		System.out.println("The creature is swimming.");
	}

	public void move(){
		System.out.println("The creature is moving.");
	}
}

//Bird //65070503445 Chanawat Limpanatewin
class Bird extends Animal {
	public Bird(int age, String gender, double weightInLbs){
		super(age ,gender, weightInLbs);
	}

	public void move(){
		System.out.println("move");
	}
}

//Flyable //65070503457 Paratthakon Suksukhon
interface Flyable{
	void fly();
}


//Chicken //65070503457 Paratthakon Suksukhon
class Chicken extends Bird {
	public Chicken(int age, String gender, int weightInLbs) {
		super(age, gender, weightInLbs);
	}
}

//Sparrow 65070503409 Chayaphon Chaisangkha
class Sparrow extends Bird implements Flyable{
	public Sparrow(int age, String gender, int weightInLbs) {
		super(age, gender, weightInLbs);
	}
	public void fly() {
		System.out.println("The creature is flying");
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
