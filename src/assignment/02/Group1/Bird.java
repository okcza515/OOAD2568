class Bird extends Animal{
    Bird(int age, String gender, int weightInLbs) {
        super(age, gender, weightInLbs);
    }

    void move(){
        System.out.println("Bird is moving");
    }
}
