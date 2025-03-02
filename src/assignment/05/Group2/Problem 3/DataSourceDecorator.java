public abstract class DataSourceDecorator implements Datasource {
    private Datasource wrappee;

    public DataSourceDecorator(Datasource source) {
        this.wrappee = source;
    }

    @Override
    public void writeData(String data) {
        wrappee.writeData(data);
    }

    @Override
    public String readData() {
        return wrappee.readData();
    }
}
// Ratchanon Tarawan 65070503464