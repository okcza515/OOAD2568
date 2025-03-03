//Nutchanon Boonyato
public class Test {

    public static void main(String[] args) {
        String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000\nNutchanon,200000";

        // Wrap FileDataSource with Compression and Encryption
        DataSource compressedAndEncrypted = new CompressionDecorator(
                new EncryptionDecorator(new FileDataSource("OutputDemo.txt"))
        );

        // Compress and Encrypt before writing
        compressedAndEncrypted.writeData(salaryRecords);

        // Read the compressed and encrypted data (Base64 encoded)
        String compressedEncryptedData = new FileDataSource("OutputDemo.txt").readData();
        System.out.println("\nCompressed & Encrypted Data Written: \n" + compressedEncryptedData);

        // Print original and compressed sizes
        System.out.println("\nOriginal Data Size: " + salaryRecords.getBytes().length + " bytes");
        System.out.println("Compressed Data Size: " + compressedEncryptedData.getBytes().length + " bytes");

        // Read, Decrypt, and Decompress the data
        String decryptedAndDecompressedData = compressedAndEncrypted.readData();
        System.out.println("\nDecrypted & Decompressed Data Read: \n" + decryptedAndDecompressedData);
    }
}
