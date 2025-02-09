class SmartphoneManufacturingProcess extends GeneralManufacturingProcess {

    @Override
    void assembly() {
        System.out.println("Smartphone assembly");
    }

    @Override
    void testing() {
        System.out.println("Smartphone testing");
    }

    @Override
    void packaging() {
        System.out.println("Smartphone packaging");
    }

    @Override
    void storage() {
        System.out.println("Smartphone storage");
    }
}