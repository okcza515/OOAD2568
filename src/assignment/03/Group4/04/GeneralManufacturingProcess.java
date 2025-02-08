abstract public class GeneralManufacturingProcess implements ManufacturingProcess{
    protected abstract void assembly();
    protected abstract void testing();
    protected abstract void packaging();
    protected abstract void storage();

    @Override
    public void manufacturingProcess(){
        System.out.println("Gum lung rem");
        this.assembly();
        this.testing();
        this.packaging();
        this.storage();
        System.out.println("Sed leaw ja!");
    }
}
