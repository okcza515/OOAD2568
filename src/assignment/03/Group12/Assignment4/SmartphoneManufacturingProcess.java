class SmartphoneManufacturingProcess extends GeneralManufacturingProcess {
    public SmartphoneManufacturingProcess() {
        super("Smartphone");
    }

    @Override
    protected void assembly() {
        System.out.println("Now, Smartphone is assembling.");
    }

    @Override
    protected void testing() {
        System.out.println("Now, Smartphone is testing.");
    }

    @Override
    protected void packaging () {
        System.out.println("Now, Smartphone is packaging.");
    }

    @Override
    protected void storage() {
        System.out.println("Now, Smartphone is storing.");
    }
}