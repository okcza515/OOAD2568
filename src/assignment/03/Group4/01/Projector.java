public class Projector implements IDevice {
  @Override
  public void turnOn() {
    System.out.println("Projector Power On");
  }

  @Override
  public void turnOff() {
    System.out.println("Projector Power Off");
  }

}
