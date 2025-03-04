public class Application {
    public static void main(String[] args) {
        String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000";
        DataSource plain = new FileDataSource("OutputDemo.txt");
        DataSource encrypted = new EncryptionDecorator(plain);
        DataSource compressed = new CompressionDecorator(encrypted);
        // Write data with encryption and compression
        compressed.writeData(salaryRecords);
        // Test Read Data from Compressed and Encrypted File
        System.out.println(compressed.readData());
    }
}
