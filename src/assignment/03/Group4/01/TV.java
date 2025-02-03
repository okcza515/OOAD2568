public class TV implements IDevice {
  @Override
  public void turnOn() {
    System.out.println("TV Power On");
  }

  @Override
  public void turnOff() {
    System.out.println("TV Power Off");
  }

}
