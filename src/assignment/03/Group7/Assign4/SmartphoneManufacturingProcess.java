public class SmartphoneManufacturingProcess extends GeneralManufacturingProcess {

  @Override
  protected void assembly() {
      System.out.println("Assemble Smartphone");
  }

  @Override
  protected void testing() {
      System.out.println("Testing Smartphone");
  }

  @Override
  protected void packaging() {
      System.out.println("Packaging Smartphone");
  }

  @Override
  protected void storage() {
      System.out.println("Storing Smartphone");
  }
}

// 65070501051