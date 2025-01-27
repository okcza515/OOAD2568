class Fish extends Animal {
    public Fish(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }

    public void swim() {
        System.out.println("I'm fish, now I'm swimming.");
    }

    @Override
    public void move() {
        this.swim();
    }
}
