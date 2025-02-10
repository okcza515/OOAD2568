// Pitchayuth Jampong 65070501079
class LaptopManufacturingProcess extends GeneralManufacturingProcess{

    @Override
    public void assembly(){
        System.out.println("Assembly -> Laptop");
    }

    @Override
    public void testing(){
        System.out.println("Testing -> Laptop");
    }

    @Override
    public void packaging(){
        System.out.println("Packaging -> Laptop");
    }

    @Override
    public void storage(){
        System.out.println("Storage -> Laptop");
    }
}