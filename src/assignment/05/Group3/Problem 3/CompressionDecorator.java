public class CompressionDecorator extends DataSourceDecorator {
    private final Compression compression;

    public CompressionDecorator(DataSource source) {
        super(source);
        this.compression = new Compression();
    }

    @Override
    public void writeData(String data) {
        super.writeData(this.compression.compress(data));
    }

    @Override
    public String readData() {
        return this.compression.decompress(super.readData());
    }
}

//65070501067
//Kanasorn Sudyodbunphot