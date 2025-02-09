public class LaptopManufacturingProcess extends GeneralManufacturingProcess {

    @Override
    public void assembly() {
        super.assembly();
        System.out.println("Assembling the laptop");
    }

    @Override
    public void testing() {
        super.testing();
        System.out.println("Testing the laptop");
    }

    @Override
    public void packaging() {
        super.packaging();
        System.out.println("Packaging the laptop");
    }

    @Override
    public void storage() {
        super.storage();
        System.out.println("Storaging the laptop");
    }
}

//Supanut Wongtanom 65070503437
