public class SmartphoneManufacturingProcess extends GeneralManufacturingProcess {

    @Override
    public void assembly(){
        super.assembly();
        System.out.println("Assembling Smartphone components");
    }

    @Override
    public void testing(){
        super.testing();
        System.out.println("Testing Smartphone");
    }

    @Override
    public void packaging(){
        super.packaging();
        System.out.println("Packaging Smartphone");
    }

    @Override
    public void storage(){
        super.storage();
        System.out.println("Storing Smartphone");
    }
    
}
//Korawit Sritotum 65070503402