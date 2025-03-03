//65070501019 Natlada Simasathien
public class Application {
    public static void main(String[] args) {
        String salaryRecords = "Name,Salary\nRudfaan Narak,100000\nBangfaan Rakjring,912000";

        DataSource plainFile = new FileDataSource("Output_Plain.txt");
        plainFile.writeData(salaryRecords);

        DataSource encryptedFile = new EncryptionDecorator(new FileDataSource("Output_Encrypted.txt"));
        encryptedFile.writeData(salaryRecords);

        DataSource compressedFile = new CompressionDecorator(new FileDataSource("Output_Compressed.txt"));
        compressedFile.writeData(salaryRecords);

        DataSource encryptedCompressedFile = new CompressionDecorator(new EncryptionDecorator(new FileDataSource("Output_Encrypted_Compressed.txt")));
        encryptedCompressedFile.writeData(salaryRecords);

        System.out.println("Plain Text Output:\n" + plainFile.readData());
        System.out.println("\nEncrypted Output:\n" + encryptedFile.readData());
        System.out.println("\nCompressed Output:\n" + compressedFile.readData());
        System.out.println("\nEncrypted & Compressed Output:\n" + encryptedCompressedFile.readData());
    }
}