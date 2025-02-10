class SmartphoneManufacturingProcess extends GeneralManufacturingProcess{

    // 65070501011 Chayapol Wongpuwarak
    @Override
    public void assembly(){
        System.out.println("Assembly -> Smartphone");
    }

    @Override
    public void testing(){
        System.out.println("Testing -> Smartphone");
    }

    // 65070501019 Natlada Simasathien
    @Override
    public void packaging(){
        System.out.println("Packaging -> Smartphone");
    }

    @Override
    public void storage(){
        System.out.println("Storage -> Smartphone");
    }

}