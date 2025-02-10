class SmartphoneManufacturingProcess extends GeneralManufacturingProcess {
    public SmartphoneManufacturingProcess() {
        super("Smartphone");
    }

    @Override
    public void assembly() {
        System.out.println("Now, Smartphone is assemblying.");
    }

    @Override
    public void testing() {
        System.out.println("Now, Smartphone is testing.");
    }

    @Override
    public void packaging () {
        System.out.println("Now, Smartphone is packaging.");
    }

    @Override
    public void storage() {
        System.out.println("Now, Smartphone is storing.");
    }
}