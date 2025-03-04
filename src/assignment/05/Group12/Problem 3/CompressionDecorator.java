public class CompressionDecorator implements DataSource {
    private DataSource wrappee;
    private Compression compression = new Compression();

    public CompressionDecorator(DataSource source) {
        this.wrappee = source;
    }

    @Override
    public void writeData(String data) {
        wrappee.writeData(compression.compress(data));
    }

    @Override
    public String readData() {
        return compression.decompress(wrappee.readData());
    }
}