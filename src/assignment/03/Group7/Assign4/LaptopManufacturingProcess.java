public class LaptopManufacturingProcess extends GeneralManufacturingProcess {

  @Override
  protected void assembly() {
    System.out.println("Assemble Laptop");
  }

  @Override
  protected void testing() {
    System.out.println("Testing Laptop");
  }

  @Override
  protected void packaging() {
    System.out.println("Packaging Laptop");
  }

  @Override
  protected void storage() {
    System.out.println("Storing Laptop");
  }
}

// 65070501051