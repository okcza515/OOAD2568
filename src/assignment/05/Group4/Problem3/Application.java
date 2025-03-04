public class Application {
    public static void main(String[] args) {
        String salaryRecords = "Name,Salary\nJohn Smith,100000\nSteven Jobs,912000";

        System.out.println("Normal Output file");
        FileDataSource plain = new FileDataSource("OutputDemo.txt");
        plain.writeData(salaryRecords);
        System.out.println(plain.readData() + "\n");

        System.out.println("Compression Output file");
        FileDataSource compressionFile = new CompressionDecorator(
                new FileDataSource("OutputCompression.txt"));
        compressionFile.writeData(salaryRecords);
        System.out.println(compressionFile.readData() + "\n");

        System.out.println("Encryption Output file");
        FileDataSource encodeFile = new EncryptionDecorator(
                new FileDataSource("OutputEncode.txt"));
        encodeFile.writeData(salaryRecords);
        System.out.println(encodeFile.readData() + "\n");

        System.out.println("Compress-Encryption Output file");
        FileDataSource compressEncryptFile = new CompressionDecorator(
                new EncryptionDecorator(
                        new FileDataSource("OutputCompressEncode.txt")));
        compressEncryptFile.writeData(salaryRecords);
        System.out.println(compressEncryptFile.readData() + "\n");

        System.out.println("Encryption-Compress Output file");
        FileDataSource encryptCompressFile = new EncryptionDecorator(
                new CompressionDecorator(
                        new FileDataSource("OutputEncodeCompress.txt")));
        encryptCompressFile.writeData(salaryRecords);
        System.out.println(encryptCompressFile.readData() + "\n");
    }
}