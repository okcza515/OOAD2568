class LaptopManufacturingProcess extends GeneralManufacturingProcess {
  
    @Override
    protected void assembly() {
        System.out.println("Laptop assembly");
    }

    @Override
    protected void testing() {
        System.out.println("Laptop testing");
    }

    @Override
    protected void packaging() {
        System.out.println("Laptop packaging");
    }

    @Override
    protected void storage() {
        System.out.println("Laptop storage");
    }
}

// 65070501067 Kanasorn Sudyodbunphot