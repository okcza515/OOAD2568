public class TV implements Device {
  @Override
  public void turnOn() {
    System.out.println("TV Power On");
  }

  @Override
  public void turnOff() {
    System.out.println("TV Power Off");
  }

}
