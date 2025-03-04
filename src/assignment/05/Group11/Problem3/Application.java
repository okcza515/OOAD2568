public class Application {
    public static void main(String[] args) {
        String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000";
        
        DataSource plain = new FileDataSourceAdapter("plain.txt");
        plain.writeData(salaryRecords);
        System.out.println("Plain data:");
        System.out.println(plain.readData());

        DataSource encrypted = new EncryptionDecorator(new FileDataSourceAdapter("encrypted.txt"));
        encrypted.writeData(salaryRecords);
        System.out.println("Encrypted data:");
        System.out.println(encrypted.readData());
        
        DataSource compressed = new CompressionDecorator(new FileDataSourceAdapter("compressed.txt"));
        compressed.writeData(salaryRecords);
        System.out.println("Compressed data:");
        System.out.println(compressed.readData());
        
        DataSource compressedAndEncrypted = new CompressionDecorator(new EncryptionDecorator(new FileDataSourceAdapter("compressedAndEncrypted.txt")));
        compressedAndEncrypted.writeData(salaryRecords);
        System.out.println("Compressed and encrypted data:");
        System.out.println(compressedAndEncrypted.readData());
    }
}
// 65070503412 Chitsanupong Jateassavapirom