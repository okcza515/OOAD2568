class LaptopManufacturingProcess extends GeneralManufacturingProcess {
    public LaptopManufacturingProcess() {
        super("Laptop");
    }

    protected void assembly() {
        System.out.println("Assembled Laptop");
    }

    protected void testing() {
        System.out.println("Tested Laptop");
    }

    protected void packaging() {
        System.out.println("Packaged Laptop");
    }

    protected void storage() {
        System.out.println("Stored Laptop");
    }
}