public class Application {
    public static void main(String[] args) {
        String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000";

        // Plain File Storage
        DataSource plain = new FileDataSource("OutputDemo.txt");
        plain.writeData(salaryRecords);
        System.out.println("Plain Read: \n" + plain.readData());

        // Encryption Wrapper
        DataSource encrypted = new Encryption(plain);
        encrypted.writeData(salaryRecords);
        System.out.println("\nEncrypted Read: \n" + encrypted.readData());

        // Compression Wrapper
        DataSource compressed = new Compression(plain);
        compressed.writeData(salaryRecords);
        System.out.println("\nCompressed Read: \n" + compressed.readData());

        // Both Encryption and Compression
        DataSource encryptedCompressed = new Encryption(new Compression(plain));
        encryptedCompressed.writeData(salaryRecords);
        System.out.println("\nEncrypted + Compressed Read: \n" + encryptedCompressed.readData());
    }
}
//65070501042