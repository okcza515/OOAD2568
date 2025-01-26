// 65070501025 นายธเนศ จอมพูล
class Chicken extends Bird {

    public Chicken(int age, String gender, double weightInLbs) {
        super(age, gender, weightInLbs);
    }

    @Override
    public void move() {
        System.out.println("Chicken is moving.");
    }
}
