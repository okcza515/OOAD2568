class Bird extends Animal{
    public Bird(int age, String gender, int weightInLbs){
        super(age, gender, weightInLbs);
    }
    @Override
    void eat(){
        System.out.println("Bird is eating");
    }
    @Override
    void sleep(){
        System.out.println("Bird is sleeping");
    }
    @Override
    void move(){
        System.out.println("Bird is moving");
    }
}