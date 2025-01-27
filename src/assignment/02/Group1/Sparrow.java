class Sparrow extends Bird implements Flyable {
    Sparrow(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    void fly() {
        System.out.println("Sparrow is flying.");
    }
}