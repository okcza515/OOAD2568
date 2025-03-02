public class CompressionDecorator extends DataSourceDecorator {
    public CompressionDecorator(Datasource source) {
        super(source);
    }

    @Override
    public void writeData(String data) {
        Compression compression = new Compression();
        super.writeData(compression.compress(data));
    }

    @Override
    public String readData() {
        Compression compression = new Compression();
        return compression.decompress(super.readData());
    }
}
// Ratchanon Tarawan 65070503464
