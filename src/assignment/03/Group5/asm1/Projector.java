class Projector implements Device {
    @Override
    public void turnOn() {
        System.out.println("Projector is turned on.");
    }

    @Override
    public void turnOff() {
        System.out.println("Projector is turned off.");
    }
}