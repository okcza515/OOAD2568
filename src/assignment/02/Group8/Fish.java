class Fish extends Animal{
    public Fish(int age, String gender, int weightInLbs)
    {
        super(age, gender, weightInLbs);
    }
    @Override
    public void eat()
    {
        System.out.println("The fish is eating");
    }
    @Override
    public void sleep()
    {
        System.out.println("The fish is sleeping Zzz");
    }
    public void swim()
    {
        System.out.println("The fish is swimming");
    }
    @Override
    public void move()
    {
        swim();
    }
}