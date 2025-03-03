public class Application {
    public static void main(String[] args) {
        String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000";

        DataSource file = new FileDataSource("OutputDemo.txt");

        // use Decorators
        DataSource encrypted = new Encryption(file);
        DataSource compressed = new Compression(encrypted);

        // write data (encrypt + compress)
        compressed.writeData(salaryRecords);
        System.out.println("Data written successfully!");

        // read data (decrypt + decompress)
        String readData = compressed.readData();
        System.out.println("Decrypted Data: \n" + readData);
    }
}
