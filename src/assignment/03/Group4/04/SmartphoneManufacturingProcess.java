public class SmartphoneManufacturingProcess extends GeneralManufacturingProcess {

  @Override
  protected void assembly() {
    System.out.println("Assembling smartphone...");
  }

  @Override
  protected void testing() {
    System.out.println("Testing smartphone...");
  }

  @Override
  protected void packaging() {
    System.out.println("Packaging smartphone...");
  }

  @Override
  protected void storage() {
    System.out.println("Storage smartphone...");
  }
}
