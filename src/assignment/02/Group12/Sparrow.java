class Sparrow extends Bird implements Flyable {
    Sparrow(int age, String gender, int weightInLbs) {}

    @Override
    public void fly() {
        System.out.println("Sparrow is flying jib jib jib.");
    }
}