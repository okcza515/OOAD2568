public class BaseDecorator implements FileData {
    private FileData wrappee;

    public BaseDecorator(FileData source) {
        this.wrappee = source;
    }

    public void writeData(String data) {
        wrappee.writeData(data);
    }

    public String readData() {
        return wrappee.readData();
    }
}