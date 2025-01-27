class Fish extends Animal {
    Fish(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }

    private void swim() {
        System.out.println("Fish is swimming(moving)");
    }

    @Override
    void move() {
        swim();
    }
}
